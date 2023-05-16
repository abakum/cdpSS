package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func s13(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		params = conf.P[strconv.Itoa(slide)]
	)
	stdo.Println(params, sc, rf)
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
	div := "div[aria-label=mrf]"
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(div, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(fmt.Sprintf("span[title='%s']", params[1]), chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d %s.png", slide, params[1]))
	div = "div[aria-label='РФ']"
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(div, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d РФ.png", slide))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(fmt.Sprintf("span[title='%s']", rf), chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d %s.png", slide, rf))
	div = "div[aria-label='СЦ']"
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(div, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d СЦ.png", slide))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(fmt.Sprintf("span[title='%s']", sc), chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(div, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d %s.png", slide, sc))

	ex(slide, chromedp.Run(ct1,
		chromedp.WaitVisible("//*[text()='Кавр.Динамика по:']"),
	))
	scs(slide, ct1, fmt.Sprintf("%02d Кавр.Динамика по.png", slide))

	ex(slide, chromedp.Run(ct1,
		chromedp.WaitVisible("//*[text()='Пред.Неделя']"),
	))
	scs(slide, ct1, fmt.Sprintf("%02d Пред.Неделя.png", slide))
	// ex(slide, chromedp.Run(ct1,
	// 	chromedp.WaitNotPresent("div.circle"),
	// 	chromedp.Sleep(time.Second),
	// ))
	// scs(slide, ct1, fmt.Sprintf("%02d no circle.png", slide))
	ex(slide, chromedp.Run(ct1,
		Screenshot("div.visualContainerHost", &bytes, 99),
	))
	ss(bytes).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
