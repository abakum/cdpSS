package main

import (
	"context"
	"fmt"
	"strconv"

	dp "github.com/chromedp/chromedp"
)

func s13(slide int) {
	var (
		params = conf.P[strconv.Itoa(abs(slide))]
	)
	stdo.Println(params, sc, rf)
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

	tit = "mrf"
	sel := fmt.Sprintf("div[aria-label=%s]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))

	tit = params[1]
	sel = fmt.Sprintf("span[title='%s']", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "РФ"
	sel = fmt.Sprintf("div[aria-label='%s']", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))

	tit = rf
	sel = fmt.Sprintf("span[title='%s']", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "СЦ"
	se := fmt.Sprintf("div[aria-label='%s']", tit)
	ex(slide, dp.Run(ct,
		dp.Click(se, dp.NodeVisible),
		dp.Sleep(ms),
	))

	tit = sc
	sel = fmt.Sprintf("span[title='%s']", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	ex(slide, dp.Run(ct,
		dp.Click(se, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "Кавр.Динамика по:"
	sel = fmt.Sprintf("//*[text()='%s']", tit)
	ex(slide, dp.Run(ct,
		dp.WaitVisible(sel),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "Пред.Неделя"
	sel = fmt.Sprintf("//*[text()='%s']", tit)
	ex(slide, dp.Run(ct,
		dp.WaitVisible(sel),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	bytes := []byte{}
	ex(slide, dp.Run(ct,
		Screenshot("div.visualContainerHost", &bytes, 99),
	))
	ss(bytes).write(fmt.Sprintf("%02d.jpg", slide))
	done(slide)
}
