package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chromedp/cdproto/page"
	dp "github.com/chromedp/chromedp"
)

func s12(slide int) {
	var (
		params = conf.P[strconv.Itoa(abs(slide))]
		vc27,
		vc22,
		vcHost page.Viewport
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

	tit = params[1]
	sel := fmt.Sprintf("div[aria-label=%s]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "RF"
	sel = fmt.Sprintf("div[aria-label=%s]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))

	tit = params[2]
	sel = fmt.Sprintf("span[title='%s']", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "SC_NAME"
	sel = fmt.Sprintf("div[aria-label=%s]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))

	tit = sc
	se := fmt.Sprintf("span[title='%s']", tit)
	ex(slide, dp.Run(ct,
		dp.Click(se, dp.NodeVisible),
		dp.Sleep(ms),
	))
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "circle"
	sel = "div.circle"
	ex(slide, dp.Run(ct,
		dp.WaitNotPresent(sel),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	sel = "//*[@id='pvExplorationHost']/div/div/exploration/div/explore-canvas/div/div[2]/div/div[2]/div[2]/visual-container-repeat/visual-container[27]/transform/div/div[3]/div/visual-modern/div/div/div[2]/div[1]/div[3]/div/div[2]"
	ex(slide, dp.Run(ct,
		getClientRect(sel, &vc27),
	))
	sel = "//*[@id='pvExplorationHost']/div/div/exploration/div/explore-canvas/div/div[2]/div/div[2]/div[2]/visual-container-repeat/visual-container[22]/transform/div/div[3]/div/visual-modern/div/div/div[2]/div[1]/div[4]"
	ex(slide, dp.Run(ct,
		getClientRect(sel, &vc22),
	))
	ex(slide, dp.Run(ct,
		getClientRect("div.visualContainerHost", &vcHost),
	))

	bytes := []byte{}
	ex(slide, dp.Run(ct,
		FullScreenshot(&bytes, 99, clip(
			vcHost.X,
			vcHost.Y,
			vc22.X+vc22.Width-vcHost.X,
			vc27.Y-vcHost.Y,
		)),
	))
	ss(bytes).write(fmt.Sprintf("%02d.jpg", slide))
	done(slide)
}
