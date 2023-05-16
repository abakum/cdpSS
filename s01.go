package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

func s01(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		params = conf.P[strconv.Itoa(1)]
	)
	stdo.Println(params)
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
	go func() {
		chromedp.Run(ct1,
			chromedp.EmulateViewport(1920, 1080),
			chromedp.Navigate(params[0]),
		)
	}()
	time.Sleep(time.Second * 3) //Navigate
	ct1, ca1 = context.WithTimeout(ct1, to)
	defer ca1()
	bytes := []byte{}
	ex(slide, chromedp.Run(ct1,
		Screenshot("div > table.weather", &bytes, 99),
	))
	ss(bytes).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
