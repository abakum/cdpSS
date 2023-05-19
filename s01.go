package main

import (
	"fmt"
	"strconv"
	"time"

	dp "github.com/chromedp/chromedp"
)

func s01(slide int) {
	var (
		params = conf.P[strconv.Itoa(abs(slide))]
	)
	stdo.Println(params)
	ct, ca := chrome()
	defer ca()
	go func() {
		dp.Run(ct,
			EmulateViewport(1920, 1080),
			dp.Navigate(params[0]),
		)
	}()
	time.Sleep(sec * 3) //Navigate
	bytes := []byte{}
	ex(slide, Run(ct, to,
		Screenshot("div > table.weather", &bytes, 99),
	))
	ss(bytes).write(fmt.Sprintf("%02d.jpg", slide))
	done(slide)
}
