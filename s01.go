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
	// time.Sleep(time.Second) //start chrome
	ct1, ca1 := chromedp.NewContext(ct0)
	defer ca1()
	go func() {
		ct1, ca1 = context.WithTimeout(ct1, time.Minute)
		defer ca1()
		chromedp.Run(ct1,
			chromedp.EmulateViewport(1920, 1080),
			chromedp.Navigate(params[0]),
		)
	}()
	time.Sleep(time.Second) //Navigate
	img := ii{}
	ct1, ca1 = context.WithTimeout(ct1, time.Minute)
	defer ca1()
	ex(slide, chromedp.Run(ct1,
		chromedp.Screenshot("div > table.weather", &img.png, chromedp.NodeVisible),
	))
	ssII(&img).write(fmt.Sprintf("%02d.jpg", slide))
	stdo.Printf("%02d Done", slide)
}
