package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	dp "github.com/chromedp/chromedp"
	"github.com/xlab/closer"
)

const (
	doc         = "doc"
	bat         = "abaku.bat"
	mov         = "abaku.mp4"
	userDataDir = `Google\Chrome\User Data\Default`
	to          = time.Minute * 3
	ms          = time.Millisecond * 200
	sec         = time.Second
)

var (
	deb  int
	stdo *log.Logger
	cd   string // s:\bin
	root string // s:
	exit int    = 2
	sc   string
	rf   string
	ctRoot,
	ctTab context.Context
	caRoot,
	caTab context.CancelFunc
	options      []func(*dp.ExecAllocator)
	multiBrowser = true
	headLess     = true
	upload       = false
	sequentially = false
)

func main() {
	var (
		wg  sync.WaitGroup
		err error
	)
	defer func() {
		exit = 0
		closer.Close()
	}()
	stdo = log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	cd, err = os.Getwd()
	ex(2, err)
	stdo.Println("Getwd:", cd)
	root = filepath.Dir(cd)
	slides := []int{}

	for _, s := range os.Args[1:] {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		switch i {
		case 0:
			multiBrowser = false
			continue
		case 2:
			headLess = false
			continue
		case -2:
			multiBrowser = true
			continue
		case 3, 6:
			slides = []int{1, 4, 5, 8, 12, 13}
		case -3, -6:
			slides = []int{-1, -4, -5, -8, -12, -13}
		}
		if abs(i) == 3 {
			break
		}
		if abs(i) == 6 {
			sequentially = true
			break
		}
		slides = append(slides, i)
		if i > 0 {
			headLess = false
		}
		if i < 0 {
			multiBrowser = true
		}
	}
	// ""  mb 1 hl 1 debug 0
	// "0" mb 1 hl 1 debug 0
	// "2" mb x hl 0 debug 1
	//"-2" mb 1 hl x debug 0
	// "x" mb 0 hl 0 debug 1 only sx
	//"-x" mb 0 hl 1 debug 0 only sx
	// "3" mb 0 hl 1 debug 1 no publicate
	//"-3" mb 0 hl 0 debug 0 no publicate
	if len(slides) == 0 {
		slides = append(slides, 0)
	}
	options = append(
		dp.DefaultExecAllocatorOptions[:],
		dp.Flag("enable-automation", false),
		dp.Flag("start-maximized", true),
	)
	if headLess {
		options = append(options,
			dp.Flag("headless", "new"),
			// dp.DisableGPU,
		)
	} else {
		options = append(options,
			dp.Flag("headless", false),
		)
	}
	exeN, exeF, err := exeFN()
	ex(2, err)
	conf, err = loader(filepath.Join(cd, exeF+".json"))
	if err != nil {
		conf.P = map[string][]string{}
		conf.Ids = []int{}
		conf.saver()
		ex(2, err)
		return
	}
	sc = conf.P["4"][1]
	rf = conf.P["12"][2]
	ctRoot, caRoot = context.WithCancel(context.Background())
	defer caRoot()
	if !multiBrowser {
		// in multitab mode with one browser instance some tab has hang
		// regardless of chrome://flags/#high-efficiency-mode-available
		options = append(options,
			dp.UserDataDir(filepath.Join(os.Getenv("LOCALAPPDATA"), userDataDir)),
		)
		ctTab, caTab = dp.NewExecAllocator(ctRoot, options...)
		defer caTab()
		ctTab, caTab = dp.NewContext(ctTab)
		defer caTab()
		ex(deb, dp.Run(ctTab, // first Run create browser instance
			dp.Navigate("about:blank"),
		))
		// time.AfterFunc(sec*3, func() {
		// 	dp.Run(ctTab, dp.Evaluate("window.close();", nil)) // close empty tab
		// })
	}
	closer.Bind(func() {
		deb = 2 //exit
		caRoot()
		if upload {
			taskKill("/fi", "windowtitle eq Открытие")
		}
		stdo.Println("main Done", exit)
		switch {
		case exit == 0:
		case exit < 0:
			exit = -exit
			fallthrough
		default:
			time.Sleep(sec) // for caRoot
			taskKill("/F", "/IM", exeN, "/T")
		}
		os.Exit(exit)
	})
	started := make(chan bool, 10)
	autoStart(started, sec) //for wg.Add
	for _, de := range slides {
		stdo.Println(de, "multiBrowser:", multiBrowser, "headLess:", headLess, "sequentially:", sequentially)
		deb = de
		go start(s01, 1, &wg, started)
		go start(s04, 4, &wg, started)
		go start(s05, 5, &wg, started)
		go start(s08, 8, &wg, started)
		go start(s12, 12, &wg, started)
		go start(s13, 13, &wg, started)
		if sequentially {
			<-started //for wg.Add
			wg.Wait()
			autoStart(started, sec) //for wg.Add
		}
	}
	if !sequentially {
		<-started //for wg.Add
		wg.Wait()
	}
	start(s97, 97, nil, nil)        //bat jpgs to mov
	autoStart(started, sec)         //for wg.Add
	go start(s98, 98, &wg, started) //telegram
	start(s99, 99, nil, nil)
	<-started //for wg.Add of s98
	wg.Wait()
}
