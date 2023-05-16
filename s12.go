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

func s12(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		params = conf.P[strconv.Itoa(slide)]
		vc27,
		vc22,
		vcHost page.Viewport
	)
	stdo.Println(params, sc)
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
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(fmt.Sprintf("div[aria-label='%s']", params[1]), chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d %s.png", slide, params[1]))
	div := "div[aria-label=RF]"
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(div, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d RF.png", slide))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(fmt.Sprintf("span[title='%s']", params[2]), chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d %s.png", slide, params[2]))
	div = "div[aria-label=SC_NAME]"
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(div, chromedp.NodeVisible),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d SC_NAME.png", slide))
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
		chromedp.WaitNotPresent("div.circle"),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d no circle.png", slide))

	ex(slide, chromedp.Run(ct1,
		getClientRect("//*[@id='pvExplorationHost']/div/div/exploration/div/explore-canvas/div/div[2]/div/div[2]/div[2]/visual-container-repeat/visual-container[27]/transform/div/div[3]/div/visual-modern/div/div/div[2]/div[1]/div[3]/div/div[2]",
			&vc27),
	))
	ex(slide, chromedp.Run(ct1,
		getClientRect("//*[@id='pvExplorationHost']/div/div/exploration/div/explore-canvas/div/div[2]/div/div[2]/div[2]/visual-container-repeat/visual-container[22]/transform/div/div[3]/div/visual-modern/div/div/div[2]/div[1]/div[4]",
			&vc22),
	))
	ex(slide, chromedp.Run(ct1,
		getClientRect("div.visualContainerHost", &vcHost),
	))
	ex(slide, chromedp.Run(ct1,
		FullScreenshot(&bytes, 99, clip(
			vcHost.X,
			vcHost.Y,
			vc22.X+vc22.Width-vcHost.X,
			vc27.Y-vcHost.Y,
		)),
	))
	ss(bytes).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
