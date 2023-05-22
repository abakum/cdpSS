package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chromedp/cdproto/page"
	dp "github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func s04(slide int) {
	var (
		params = conf.P[strconv.Itoa(abs(slide))]
		imageBackground,
		visualContainerHost page.Viewport
		tit string
	)
	stdo.Println(params)
	ct, ca := chrome()
	defer ca()
	dp.Run(ct,
		EmulateViewport(1920, 1080),
		dp.Navigate(params[0]),
		dp.Sleep(sec),
		dp.Title(&tit),
	)
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))
	ct, ca = context.WithTimeout(ct, to)
	defer ca()
	bytes := []byte{}
	// iframe(slide, ct, params[0])
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	cb(slide, ct, "СЦ")

	ex(slide, dp.Run(ct,
		getClientRect("div.imageBackground", &imageBackground, dp.NodeVisible),
	))
	ex(slide, dp.Run(ct,
		getClientRect("div.visualContainerHost", &visualContainerHost, dp.NodeVisible),
	))
	ex(slide, dp.Run(ct,
		FullScreenshot(&bytes, 99, clip(
			visualContainerHost.X,
			visualContainerHost.Y,
			imageBackground.X-visualContainerHost.X,
			visualContainerHost.Height,
		)),
	))
	ss(bytes).write(fmt.Sprintf("%02d.jpg", slide))
	done(slide)
}

func cb(slide int, ctx context.Context, key string) {
	tit := "СЦ"
	se := fmt.Sprintf("div[aria-label=%q] > i", key)
	ex(slide, dp.Run(ctx,
		dp.Click(se, dp.NodeVisible),
		dp.Sleep(ms),
	))

	scs(slide, ctx, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "Поиск"
	sel := "div.searchHeader.show > input"
	ex(slide, Run(ctx, to*2,
		dp.SetValue(sel, sc, dp.NodeEnabled, dp.NodeVisible),
		dp.SendKeys(sel, kb.Enter),
		// chromedp.SendKeys(inp, sc, chromedp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ctx, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = sc
	sel = fmt.Sprintf("//span[.=%q]", tit)
	ex(slide, Run(ctx, to*2,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	ex(slide, dp.Run(ctx,
		dp.Click(se, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ctx, fmt.Sprintf("%02d %s.png", slide, tit))

	ex(slide, Run(ctx, to*3,
		dp.WaitNotPresent("div.circle"),
		dp.Sleep(sec),
	))
}
