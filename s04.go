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
		params = conf.P[strconv.Itoa(slide)]
		imageBackground,
		visualContainerHost page.Viewport
	)
	stdo.Println(params)
	ct, ca := chrome()
	defer ca()
	dp.Run(ct,
		EmulateViewport(1920, 1080),
		dp.Navigate(params[0]),
		dp.Sleep(sec),
	)
	ct, ca = context.WithTimeout(ct, to)
	defer ca()
	bytes := []byte{}
	// iframe(slide, ct, params[0])
	tit := "Navigate"
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
	se := fmt.Sprintf("div[aria-label='%s'] > i", key)
	ex(slide, dp.Run(ctx,
		dp.Click(se, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ctx, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "Поиск"
	sel := "div.searchHeader.show > input"
	ex(slide, dp.Run(ctx,
		dp.SetValue(sel, sc, dp.NodeEnabled, dp.NodeVisible),
		dp.SendKeys(sel, kb.Enter),
		// chromedp.SendKeys(inp, sc, chromedp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ctx, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = sc
	sel = fmt.Sprintf("//span[.='%s']", tit)
	ex(slide, dp.Run(ctx,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	ex(slide, dp.Run(ctx,
		dp.Click(se, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ctx, fmt.Sprintf("%02d %s.png", slide, tit))

	ex(slide, dp.Run(ctx,
		dp.WaitNotPresent("div.circle"),
		dp.Sleep(ms),
	))
}
