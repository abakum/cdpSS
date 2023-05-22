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
		tit    string
	)
	stdo.Println(params, sc, rf)
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
	// iframe(slide, ct, params[0])

	tit = "mrf"
	sel := fmt.Sprintf("div[aria-label=%s]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))

	tit = params[1]
	sel = fmt.Sprintf("span[title=%q]", tit)
	ex(slide, Run(ct, to*2,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "РФ"
	sel = fmt.Sprintf("div[aria-label=%q]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))

	tit = rf
	sel = fmt.Sprintf("span[title=%q]", tit)
	ex(slide, Run(ct, to*2,
		dp.Click(sel, dp.NodeVisible),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "СЦ"
	se := fmt.Sprintf("div[aria-label=%q]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(se, dp.NodeVisible),
		dp.Sleep(ms),
	))

	tit = sc
	sel = fmt.Sprintf("span[title=%q]", tit)
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
	sel = fmt.Sprintf("//*[text()=%q]", tit)
	ex(slide, dp.Run(ct,
		dp.WaitVisible(sel),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "Пред.Неделя"
	sel = fmt.Sprintf("//*[text()=%q]", tit)
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
