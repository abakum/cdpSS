package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/cdproto/cdp"
	dp "github.com/chromedp/chromedp"
)

func s08(slide int) {
	var (
		TaskClosed = "TaskClosed.xlsx"
		params     = conf.P[strconv.Itoa(abs(slide))]
		nodes      []*cdp.Node
		aNodes,
		oNodes []*cdp.Node
	)
	stdo.Println(params)
	ct, ca := chrome()
	defer ca()
	dp.Run(ct,
		browser.SetDownloadBehavior("allow").WithDownloadPath(filepath.Join(root, doc)),
		EmulateViewport(1920, 1080),
		dp.Navigate(params[0]),
		dp.WaitReady("body"),
	)
	ct, ca = context.WithTimeout(ct, to)
	defer ca()
	tit := "По работникам и типу задачи"
	sel := fmt.Sprintf("//span[.=%q]", tit)
	for i := 0; i < 7; i++ {
		Run(ct, sec,
			dp.Title(&tit),
			dp.Nodes("#login_form-username", &aNodes),
			dp.Nodes(sel, &oNodes),
		)
		stdo.Println(i, tit, len(strings.Trim(tit, " ")), len(aNodes), len(oNodes))
		if len(strings.Trim(tit, " ")) > 0 || len(aNodes) > 0 || len(oNodes) > 0 {
			break
		}
	}
	if len(oNodes) > 0 {
		tit = "Отчеты по задачам"
	}
	if len(aNodes) > 0 {
		tit = "Аргус"
	}
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	if tit == "Аргус" {
		sel = "#login_form-username"
		ex(slide, dp.Run(ct,
			// chromedp.MouseClickNode(nodes[0]),
			dp.SetValue(sel, params[1], dp.NodeEnabled),
			// dp.SendKeys(sel, params[1], dp.NodeEnabled),
			dp.Sleep(ms),
		))
		sel = "#login_form-password"
		ex(slide, dp.Run(ct,
			dp.SetValue(sel, params[2], dp.NodeEnabled),
			// dp.SendKeys(sel, params[2], dp.NodeEnabled),
			dp.Sleep(ms),
		))

		tit = "Войти"
		sel = fmt.Sprintf("//span[.=%q]", tit)
		ex(slide, dp.Run(ct,
			// chromedp.SendKeys(input, kb.Enter, chromedp.NodeEnabled),
			// chromedp.Click("#login_form-submit", chromedp.NodeEnabled),
			dp.Click(sel, dp.NodeEnabled),
			dp.Sleep(sec),
		))
		scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))
	}

	tit = "По работникам и типу задачи"
	sel = fmt.Sprintf("//span[.=%q]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeEnabled),
		dp.Sleep(sec),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "месяцы"
	sel = fmt.Sprintf("//span[.=%q]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeEnabled),
		dp.Sleep(sec),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "Обработка наряда"
	sel = "ul.ui-widget"
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.ByQuery, dp.NodeEnabled),
		dp.Sleep(sec),
	))
	ex(slide, dp.Run(ct,
		dp.Click("//li[5]/label", dp.NodeEnabled),
		dp.Sleep(ms*2),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	for i := 4; i < 9; i++ {
		ex(slide, dp.Run(ct,
			dp.Nodes(".ui-tree-toggler", &nodes, dp.NodeEnabled),
		))
		ex(slide, dp.Run(ct,
			dp.MouseClickNode(nodes[i]),
			dp.Sleep(ms*4),
		))
	}

	tit = "Группа инсталляций"
	sel = fmt.Sprintf("//span[.=%q]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeEnabled),
		dp.Sleep(ms*3),
	))
	tit = "Группа клиентского сервиса"
	sel = fmt.Sprintf("//span[.=%q]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeEnabled),
		dp.Sleep(ms*3),
	))

	tit = "ОК"
	sel = fmt.Sprintf("//span[.=%q]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible, dp.NodeEnabled),
		dp.Sleep(ms),
	))
	os.Remove(filepath.Join(root, doc, TaskClosed))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	// time.Sleep(sec * 7)
	ex(slide, dp.Run(ct,
		dp.Click("span.ui-chkbox-label", dp.NodeVisible, dp.NodeVisible, dp.NodeEnabled),
		dp.Sleep(ms*3),
	))
	ex(slide, dp.Run(ct,
		dp.Click("button#report_actions_form-export_report_data > span", dp.NodeVisible, dp.NodeEnabled),
		dp.Sleep(sec),
	))

	tit = "EXCEL"
	sel = fmt.Sprintf("//span[.=%q]", tit)
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.NodeVisible, dp.NodeEnabled),
		dp.Sleep(sec), //for download
	))
	scs(deb, ct, fmt.Sprintf("%02d.png", slide))

	done(slide)
	s09(9)
}
