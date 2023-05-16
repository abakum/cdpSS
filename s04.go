package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func s04(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		params = conf.P[strconv.Itoa(slide)]
		imageBackground,
		visualContainerHost page.Viewport
	)
	stdo.Println(params)
	var (
		ct1 context.Context
		ca1 context.CancelFunc
	)
	if mb {
		ctLoc, caLoc := chrome()
		defer caLoc()
		ct1, ca1 = chromedp.NewContext(ctLoc)
	} else {
		ct1, ca1 = chromedp.NewContext(ctTab)
	}
	defer ca1()
	chromedp.Run(ct1,
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Navigate(params[0]),
		chromedp.Sleep(time.Second),
	)
	ct1, ca1 = context.WithTimeout(ct1, to)
	defer ca1()
	bytes := []byte{}
	if false {
		ex(slide, chromedp.Run(ct1,
			chromedp.Navigate(params[0]+"?rs:Embed=true"),
		))
		var (
			src string
			ok  bool
		)
		ex(slide, chromedp.Run(ct1,
			chromedp.WaitReady("//iframe"),
			chromedp.Sleep(time.Second),
			chromedp.AttributeValue("//iframe", "src", &src, &ok, chromedp.NodeReady),
		))
		if !ok {
			ex(slide, fmt.Errorf("no src of iframe"))
		}
		src = strings.Split(src, "&")[0]
		stdo.Println(src)
		ex(slide, chromedp.Run(ct1,
			chromedp.Navigate(src),
			chromedp.Sleep(time.Second),
		))
	}
	scs(slide, ct1, fmt.Sprintf("%02d iframe.png", slide))
	cb(slide, ct1, "СЦ")
	ex(slide, chromedp.Run(ct1,
		getClientRect("div.imageBackground", &imageBackground, chromedp.NodeVisible),
	))
	ex(slide, chromedp.Run(ct1,
		getClientRect("div.visualContainerHost", &visualContainerHost, chromedp.NodeVisible),
	))
	ex(slide, chromedp.Run(ct1,
		FullScreenshot(&bytes, 99, clip(
			visualContainerHost.X,
			visualContainerHost.Y,
			imageBackground.X-visualContainerHost.X,
			visualContainerHost.Height,
		)),
	))
	ss(bytes).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}

func cb(slide int, ct1 context.Context, key string) {
	label := fmt.Sprintf("div[aria-label='%s'] > i", key)
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(label, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d СЦ.png", slide))
	inp := "div.searchHeader.show > input"
	ex(slide, chromedp.Run(ct1,
		chromedp.SetValue(inp, sc, chromedp.NodeEnabled, chromedp.NodeVisible),
		chromedp.SendKeys(inp, kb.Enter),
		// chromedp.SendKeys(inp, sc, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d Поиск.png", slide))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(fmt.Sprintf("//span[.='%s']", sc), chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(label, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d %s.png", slide, sc))
	ex(slide, chromedp.Run(ct1,
		chromedp.WaitNotPresent("div.circle"),
		chromedp.Sleep(time.Second),
	))
}
