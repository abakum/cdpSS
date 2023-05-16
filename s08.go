package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func s08(slide int) {
	switch deb {
	case 0, slide, -slide:
	default:
		return
	}
	wg.Add(1)
	defer wg.Done()
	var (
		TaskClosed = "TaskClosed.xlsx"
		params     = conf.P[strconv.Itoa(slide)]
		nodes      []*cdp.Node
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
	ct1, ca1 = context.WithTimeout(ct1, to)
	defer ca1()
	ex(slide, chromedp.Run(ct1,
		chromedp.EmulateViewport(1920, 1080),
		browser.SetDownloadBehavior("allow").WithDownloadPath(filepath.Join(root, doc)),
		chromedp.Navigate(params[0]),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d get.png", slide))
	input := "#login_form-username"
	ct2, ca2 := context.WithTimeout(ct1, time.Second)
	defer ca2()
	chromedp.Run(ct2,
		chromedp.Nodes(input, &nodes),
	)
	if len(nodes) > 0 {
		stdo.Println(input)
		ex(slide, chromedp.Run(ct1,
			// chromedp.MouseClickNode(nodes[0]),
			chromedp.SetValue(input, params[1], chromedp.NodeEnabled),
			// chromedp.SendKeys(input, params[1], chromedp.NodeEnabled),
			chromedp.Sleep(time.Second),
		))
		input = "#login_form-password"
		ex(slide, chromedp.Run(ct1,
			chromedp.SetValue(input, params[2], chromedp.NodeEnabled),
			// chromedp.SendKeys(input, params[2], chromedp.NodeEnabled),
			chromedp.Sleep(time.Second),
		))
		ex(slide, chromedp.Run(ct1,
			// chromedp.SendKeys(input, kb.Enter, chromedp.NodeEnabled),
			// chromedp.Click("#login_form-submit", chromedp.NodeEnabled),
			chromedp.Click("//span[.='Войти']", chromedp.NodeEnabled),
			chromedp.Sleep(time.Second),
		))
		scs(slide, ct1, fmt.Sprintf("%02d Войти.png", slide))
	}
	title := "По работникам и типу задачи"
	ex(slide, chromedp.Run(ct1,
		// chromedp.Click(fmt.Sprintf("span[title='%s']", title), chromedp.NodeEnabled),
		chromedp.Click(fmt.Sprintf("//span[.='%s']", title), chromedp.NodeEnabled),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d %s.png", slide, title))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[.='месяцы']", chromedp.NodeEnabled),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d месяцы.png", slide))
	ul := "ul.ui-widget"
	ex(slide, chromedp.Run(ct1,
		chromedp.Click(ul, chromedp.ByQuery, chromedp.NodeEnabled),
		// chromedp.Click("//ul[contains(@class,'ui-selectcheckboxmenu-multiple-container')]", chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d %s.png", slide, ul))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//li[5]/label", chromedp.NodeEnabled),
		chromedp.Sleep(time.Second),
	))
	scs(slide, ct1, fmt.Sprintf("%02d Обработка наряда.png", slide))
	for i := 4; i < 9; i++ {
		ex(slide, chromedp.Run(ct1,
			chromedp.Nodes(".ui-tree-toggler", &nodes, chromedp.NodeEnabled),
		))
		ex(slide, chromedp.Run(ct1,
			chromedp.MouseClickNode(nodes[i]),
			chromedp.Sleep(time.Second),
		))
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[.='Группа инсталляций']", chromedp.NodeEnabled),
		chromedp.Sleep(time.Second),
	))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[.='Группа клиентского сервиса']", chromedp.NodeEnabled),
		chromedp.Sleep(time.Second),
	))
	scs(deb, ct1, fmt.Sprintf("%02d Группа клиентского сервиса.png", slide))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[.='ОК']", chromedp.NodeVisible, chromedp.NodeEnabled),
		chromedp.Sleep(time.Second),
	))
	os.Remove(filepath.Join(root, doc, TaskClosed))
	time.Sleep(time.Second * 7)
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("button#report_actions_form-export_report_data > span", chromedp.NodeVisible, chromedp.NodeEnabled),
		// chromedp.Click("//button[@id='report_actions_form-export_report_data']/span", chromedp.NodeEnabled),
		chromedp.Sleep(time.Second),
	))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[.='EXCEL']", chromedp.NodeVisible, chromedp.NodeEnabled),
		chromedp.Sleep(time.Second), //for download
	))
	scs(deb, ct1, fmt.Sprintf("%02d.png", slide))
	stdo.Printf("%02d Done", slide)
}
