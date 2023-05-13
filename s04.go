package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
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
		src    string
		ok     bool
	)
	stdo.Println(params)
	ct1, ca1 := chromedp.NewContext(ct0)
	defer ca1()
	ct1, ca1 = context.WithTimeout(ct1, time.Minute)
	defer ca1()
	bytes := []byte{}
	chromedp.Run(ct1,
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Navigate(params[0]),
		chromedp.Sleep(time.Second),
	)
	ex(slide, chromedp.Run(ct1,
		chromedp.Navigate(params[0]+"?rs:Embed=true"),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d get.png", slide))
		}
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.AttributeValue("//iframe", "src", &src, &ok, chromedp.NodeReady),
	))
	stdo.Println(src, ok)
	if !ok {
		ex(slide, fmt.Errorf("no src of iframe"))
	}
	src = strings.Split(src, "&")[0]
	stdo.Println(src)
	ex(slide, chromedp.Run(ct1,
		chromedp.Navigate(src),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d iframe.png", slide))
		}
	}
	cb(slide, ct1, "СЦ")
	img := ii{}
	ex(slide, chromedp.Run(ct1,
		chromedp.Screenshot("div.visualContainerHost", &img.png, chromedp.ByQuery),
	))
	ssII(&img).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}

func cb(slide int, ct1 context.Context, key string) {
	bytes := []byte{}
	label := fmt.Sprintf("div[aria-label=%s] > i", key)
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(label, chromedp.ByQuery, chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d СЦ.png", slide))
		}
	}
	inp := "div.searchHeader.show > input"
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(inp, chromedp.ByQuery, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	ex(slide, chromedp.Run(ct1,
		chromedp.SendKeys(inp, sc, chromedp.ByQuery, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d Поиск.png", slide))
		}
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(fmt.Sprintf("//span[contains(.,'%s')]", sc), chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(label, chromedp.ByQuery, chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d %s.png", slide, sc))
		}
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.WaitNotPresent("div.circle", chromedp.ByQuery),
	))
}
