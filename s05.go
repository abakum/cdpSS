package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func s05(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		params = conf.P[strconv.Itoa(slide)]
		title,
		innerContainer,
		visualContainerHost page.Viewport
	)
	stdo.Println(params, sc)
	var (
		ct1 context.Context
		ca1 context.CancelFunc
	)
	if mb {
		var (
			ctx,
			ct context.Context
			cancel,
			ca context.CancelFunc
		)
		ctx, cancel = chromedp.NewExecAllocator(context.Background(), options...)
		defer cancel()
		ct, ca = chromedp.NewContext(ctx)
		defer ca()
		ex(deb, chromedp.Run(ct,
			chromedp.EmulateViewport(1920, 1080),
			chromedp.Navigate("about:blank"),
		))
		ct1, ca1 = chromedp.NewContext(ct)
	} else {
		ct1, ca1 = chromedp.NewContext(ct0)
	}
	defer ca1()
	ct1, ca1 = context.WithTimeout(ct1, to)
	defer ca1()
	bytes := []byte{}
	chromedp.Run(ct1,
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Navigate(params[0]),
		chromedp.Sleep(time.Second),
	)
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
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("div[title='Статистика по сотрудникам']", chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d Статистика по сотрудникам.png", slide))
	cb(slide, ct1, "СЦ/ЦЭ")
	ex(slide, chromedp.Run(ct1,
		getClientRect("div[title='Ср. длительность работ сотрудника за день, часы']", &title, chromedp.NodeVisible),
	))
	ex(slide, chromedp.Run(ct1,
		getClientRect("div.innerContainer", &innerContainer, chromedp.NodeVisible),
	))
	ex(slide, chromedp.Run(ct1,
		getClientRect("div.visualContainerHost", &visualContainerHost, chromedp.NodeVisible),
	))
	scs(slide, ct1, fmt.Sprintf("%02d visualContainerHost.png", slide))
	ex(slide, chromedp.Run(ct1,
		FullScreenshot(&bytes, 99, clip(
			visualContainerHost.X,
			visualContainerHost.Y,
			title.X-visualContainerHost.X,
			innerContainer.Y+innerContainer.Height-2*visualContainerHost.Y,
		)),
	))
	ss(bytes).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
