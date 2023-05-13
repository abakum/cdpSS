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
	ct1, ca1 := chromedp.NewContext(ct0)
	defer ca1()
	ct1, ca1 = context.WithTimeout(ct1, time.Minute)
	defer ca1()
	bytes := []byte{}
	ex(slide, chromedp.Run(ct1,
		chromedp.EmulateViewport(1920, 1080),
		browser.SetDownloadBehavior("allow").WithDownloadPath(filepath.Join(root, doc)),
		chromedp.Navigate(params[0]),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d get.png", slide))
		}
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.Sleep(time.Second),
		chromedp.Click("//input[@id='login_form-username']", chromedp.NodeReady),
		chromedp.SendKeys("//input[@id='login_form-username']", params[1], chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//input[@id='login_form-password']", chromedp.NodeReady),
		chromedp.SendKeys("//input[@id='login_form-password']", params[2], chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[contains(.,'Войти')]", chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d Войти.png", slide))
		}
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[contains(.,'По работникам и типу задачи')]", chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d По работникам и типу задачи.png", slide))
		}
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[contains(.,'месяцы')]", chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d месяцы.png", slide))
		}
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//ul[contains(@class,'ui-selectcheckboxmenu-multiple-container')]", chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d ui-selectcheckboxmenu-multiple-container.png", slide))
		}
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//li[5]/label", chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fmt.Sprintf("%02d Обработка наряда.png", slide))
		}
	}
	for i := 4; i < 9; i++ {
		ex(slide, chromedp.Run(ct1,
			chromedp.Nodes("//*[contains(@class,'ui-tree-toggler')]", &nodes, chromedp.NodeReady),
		))
		ex(slide, chromedp.Run(ct1,
			chromedp.MouseClickNode(nodes[i]),
			chromedp.Sleep(time.Second),
		))
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[contains(.,'Группа инсталляций')]", chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[contains(.,'Группа клиентского сервиса')]", chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
		ss(bytes).write(fmt.Sprintf("%02d Группа клиентского сервиса.png", slide))
	}
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[contains(.,'ОК')]", chromedp.NodeReady),
		chromedp.Sleep(time.Second*7),
	))
	ex(slide, chromedp.Run(ct1,
		// chromedp.WaitVisible("//button[@id='report_actions_form-export_report_data']/span", chromedp.NodeReady),
		chromedp.Click("//button[@id='report_actions_form-export_report_data']/span", chromedp.NodeReady),
		chromedp.Sleep(time.Second),
	))
	os.Remove(filepath.Join(root, doc, TaskClosed))
	ex(slide, chromedp.Run(ct1,
		chromedp.Click("//span[contains(.,'EXCEL')]", chromedp.NodeReady),
		chromedp.Sleep(time.Second*3), //for download
	))
	if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
		ss(bytes).write(fmt.Sprintf("%02d.png", slide))
	}
	stdo.Printf("%02d Done", slide)
}
