package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/xlab/closer"
)

const (
	doc = "doc"
	bat = "abaku.bat"
	mov = "abaku.mp4"
	// chromeBin        = `C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`
	userDataDir = `Google\Chrome\User Data\Default`
	to          = time.Minute * 4
)

var (
	deb  int
	stdo *log.Logger
	wg   sync.WaitGroup
	cd   string // s:\bin
	root string // s:
	exit int
	sc   string
	rf   string
	ctRoot,
	ctTab context.Context
	caRoot,
	caTab context.CancelFunc
	options []func(*chromedp.ExecAllocator)
	mb      = false
)

func main() {
	var (
		err error
	)
	defer closer.Close()
	stdo = log.New(os.Stdout, "", log.Lshortfile|log.Ltime)
	cd, err = os.Getwd()
	ex(2, err)
	stdo.Println(cd)
	root = filepath.Dir(cd)
	slides := []int{}
	headless := true
	for _, s := range os.Args[1:] {
		i, err := strconv.Atoi(s)
		if err == nil {
			slides = append(slides, i)
		}
		if i > 0 {
			headless = false
		}
	}
	if len(slides) == 0 {
		mb = true
		slides = append(slides, 0)
	}
	options = append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("enable-automation", false),
		// chromedp.ExecPath(chromeBin),
		chromedp.Flag("window-position", "0,0"),
		// chromedp.WindowSize(1920, 1080),
	)
	if headless {
		options = append(options,
			// chromedp.Flag("headless", false),
			chromedp.Flag("headless", "new"),
			// chromedp.DisableGPU,
		)
	} else {
		options = append(options,
			chromedp.Flag("start-maximized", true),
			// chromedp.WindowSize(1366, 768),
			chromedp.Flag("headless", false),
		)
	}
	conf, err = loader(filepath.Join(cd, goSSjson))
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
	if !mb {
		// in multitab mode with one browser instance some tab has hang
		// regardless of chrome://flags/#high-efficiency-mode-available
		options = append(options,
			chromedp.UserDataDir(filepath.Join(os.Getenv("LOCALAPPDATA"), userDataDir)),
		)
		ctTab, caTab = chrome()
		defer caTab()
	}
	closer.Bind(func() {
		deb = 2 //exit
		caRoot()
		stdo.Println("main Done", exit)
		switch {
		case exit == 0:
		case exit < 0:
			exit = -exit
			fallthrough
		default:
			cmd := exec.Command("taskKill.exe", "/F", "/IM", "cdpSS.exe", "/T")
			stdo.Println(cmd.Path, strings.Join(cmd.Args[1:], " "))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
		os.Exit(exit)
	})
	for _, de := range slides {
		stdo.Println(de)
		deb = de
		go s01(1)
		// go s01(3)
		go s04(4)
		go s05(5)
		go func() {
			s08(8)
			s09(9)
		}()
		go s12(12)
		go s13(13)
		if deb < 97 {
			time.Sleep(time.Second * 2) //for wg.Add
		}
		wg.Wait()
		// closer.Close() //do not publicate
		// s97(97) //bat
		go func() {
			// s98(98) //telegram
		}()
		if deb == 98 {
			time.Sleep(time.Second) //for wg.Add
		}
		wg.Wait()
		// s99(99) //ss
	}
}
