package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chromedp/cdproto/page"
	dp "github.com/chromedp/chromedp"
)

func s05(slide int) {
	var (
		params = conf.P[strconv.Itoa(abs(slide))]
		title,
		innerContainer,
		visualContainerHost page.Viewport
	)
	stdo.Println(params, sc)
	ct, ca := chrome()
	defer ca()
	dp.Run(ct,
		EmulateViewport(1920, 1080),
		dp.Navigate(params[0]),
		dp.Sleep(sec),
	)
	ct, ca = context.WithTimeout(ct, to)
	defer ca()
	// iframe(slide, ct, params[0])

	tit := "Navigate"
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "Статистика по сотрудникам"
	sel := fmt.Sprintf("div[title='%s']", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	cb(slide, ct, "СЦ/ЦЭ")

	tit = "Ср. длительность работ сотрудника за день, часы"
	sel = fmt.Sprintf("div[title='%s']", tit)
	ex(slide, dp.Run(ct,
		getClientRect(sel, &title, dp.NodeVisible),
	))
	sel = "div.innerContainer"
	ex(slide, dp.Run(ct,
		getClientRect(sel, &innerContainer, dp.NodeVisible),
	))
	tit = "visualContainerHost"
	sel = "div.visualContainerHost"
	ex(slide, dp.Run(ct,
		getClientRect(sel, &visualContainerHost, dp.NodeVisible),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	bytes := []byte{}
	ex(slide, dp.Run(ct,
		FullScreenshot(&bytes, 99, clip(
			visualContainerHost.X,
			visualContainerHost.Y,
			title.X-visualContainerHost.X,
			innerContainer.Y+innerContainer.Height-2*visualContainerHost.Y,
		)),
	))
	ss(bytes).write(fmt.Sprintf("%02d.jpg", slide))
	done(slide)
}
