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
	to          = time.Minute * 4
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
	multiBrowser = false
	headLess     = true
	upload       = false
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
		slides = append(slides, i)
		if i > 0 {
			headLess = false
		}
	}
	// ""  mb 1 hl 1
	// "0" mb 0 hl 1
	// "2" mb 1 hl 0
	//"-2" mb 0 hl 0
	// "x" mb 0 hl 0 only sx
	//"-x" mb 0 hl 1 only sx
	//"100" mb 0 hl 1  no publicate
	//"-100" mb 0 hl 0 no publicate
	if len(slides) == 0 {
		multiBrowser = true
		slides = append(slides, 0)
	} else {
		switch slides[0] {
		case 2:
			multiBrowser = true
			slides[0] = 0
		case -2:
			multiBrowser = false
			headLess = false
			slides[0] = 0
		case 100:
			slides = []int{1, 4, 5, 8, 9, 12, 13}
		case -100:
			slides = []int{-1, -4, -5, -8, -9, -12, -13}
		}
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
		time.AfterFunc(sec*3, func() {
			dp.Run(ctTab, dp.Evaluate("window.close();", nil)) // close empty tab
		})
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
	for _, de := range slides {
		stdo.Println(de, "multiBrowser:", multiBrowser, "headLess:", headLess)
		deb = de
		go start(s01, 1, &wg)
		go start(s04, 4, &wg)
		go start(s05, 5, &wg)
		go start(s08, 8, &wg)
		go start(s12, 12, &wg)
		go start(s13, 13, &wg)
		if deb < 97 {
			time.Sleep(sec * 2) //for wg.Add
		}
	}
	wg.Wait()
	start(s97, 97, nil)    //bat jpgs to mov
	go start(s98, 98, &wg) //telegram
	if deb == 98 {
		time.Sleep(sec) //for wg.Add
	}
	start(s99, 99, nil)
	wg.Wait()
}
