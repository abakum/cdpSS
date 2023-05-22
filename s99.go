package main

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	dp "github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func s99(slide int) {
	var (
		params = conf.P[strconv.Itoa(abs(slide))]
		err    error
		aNodes,
		kNodes []*cdp.Node
	)
	stdo.Println(params)
	ct, ca := chrome()
	defer ca()
	dp.Run(ct,
		EmulateViewport(1920, 1080),
		dp.Navigate(params[0]),
		dp.WaitReady("body"),
	)
	ct, ca = context.WithTimeout(ct, to)
	defer ca()
	tit := "ar-user-name"
	sel := fmt.Sprintf("input[name=%q]", tit)
	for i := 0; i < 7; i++ {
		Run(ct, sec,
			dp.Title(&tit),
			dp.Nodes(sel, &aNodes),
			dp.Nodes("div.multiBtnInner_xbp:nth-child(1)", &kNodes),
		)
		stdo.Println(i, tit, len(strings.Trim(tit, " ")), len(aNodes), len(kNodes))
		if len(strings.Trim(tit, " ")) > 0 || len(aNodes) > 0 || len(kNodes) > 0 {
			break
		}
	}
	if len(kNodes) > 0 {
		tit = "Кампания"
	}
	if len(aNodes) > 0 {
		tit = "Авторизация"
	}
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	if tit == "Авторизация" {
		tit = "ar-user-name"
		// sel = fmt.Sprintf("input[name=%q]", tit)
		ex(slide, dp.Run(ct,
			dp.Click(sel, dp.ByQuery, dp.NodeVisible, dp.NodeEnabled),
			dp.Sleep(ms),
			dp.SendKeys(sel, params[1], dp.ByQuery, dp.NodeVisible),
			dp.Sleep(ms),
			dp.SendKeys(sel, kb.Enter, dp.ByQuery, dp.NodeVisible),
			dp.Sleep(ms),
		))
		scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

		tit = "ar-user-password"
		sel = fmt.Sprintf("input[name=%q]", tit)
		ex(slide, dp.Run(ct,
			dp.Click(sel, dp.ByQuery, dp.NodeVisible, dp.NodeEnabled),
			dp.Sleep(ms),
			dp.SendKeys(sel, params[2], dp.ByQuery, dp.NodeVisible),
			dp.Sleep(ms),
			dp.SendKeys(sel, kb.Enter, dp.ByQuery, dp.NodeVisible),
			dp.Sleep(ms),
		))
		scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

		// tit = "submit"
		// sel = fmt.Sprintf("button[type=%q]", tit)
		// ex(slide, dp.Run(ct,
		// 	dp.Click(sel, dp.ByQuery, dp.NodeEnabled),
		// 	dp.Sleep(ms),
		// ))
	}

	tit = "Кампания"
	sel = "div.multiBtnInner_xbp:nth-child(1)"
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.ByQuery, dp.NodeEnabled),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "Удалить"
	sel = "button.menu-button_J9B"
	err = Run(ct, sec,
		dp.Click(sel, dp.ByQuery, dp.NodeVisible),
		dp.Sleep(ms),
	)
	stdo.Println(tit, err)
	if err == nil {
		sel = "button.align-left_-232488494:nth-child(3)"
		ex(slide, dp.Run(ct,
			dp.Click(sel, dp.ByQuery, dp.NodeVisible),
			dp.Sleep(ms),
		))
	}

	tit = "Файл"
	sel = "button.addFilesBtn_RvX"
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.ByQuery, dp.NodeVisible),
		dp.Sleep(ms),
	))
	sel = "button.align-left_-232488494:nth-child(3)"
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.ByQuery, dp.NodeVisible),
		dp.Sleep(ms),
	))
	upload = true
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	sel = "input[type=file]"
	files := []string{filepath.Join(root, mov)}
	stdo.Println(files)
	ex(slide, dp.Run(ct,
		dp.SetUploadFiles(sel, files, dp.ByQuery),
		dp.Sleep(ms),
	))

	tit = "Загрузка"
	sel = fmt.Sprintf("//*[contains(text(),%q)]", tit)
	err = Run(ct, sec*3,
		dp.WaitVisible(sel),
		// dp.WaitVisible(tit, dp.BySearch), //no node
	)
	stdo.Println(tit, err)
	if err != nil {
		scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, "Загрузка НЕ началась"))
		Scanln()
		return
	}
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "отменена"
	sel = fmt.Sprintf("//*[contains(text(),%q)]", tit)
	err = Run(ct, sec*3,
		dp.WaitVisible(sel),
	)
	stdo.Println(tit, err)
	if err == nil {
		scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))
		Scanln()
		return
	}

	tit = "завершена"
	sel = fmt.Sprintf("//*[contains(text(),%q)]", tit)
	err = Run(ct, sec*7,
		dp.WaitVisible(sel),
	)
	stdo.Println(tit, err)
	if err != nil {
		scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, "Загрузка НЕ завершена"))
		Scanln()
		return
	}
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))
	time.Sleep(sec * 7)

	tit = "Сохранить и закрыть"
	sel = "div.multiBtnInner_xbp:nth-child(4)"
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.ByQuery, dp.NodeEnabled),
		dp.Sleep(ms),
	))

	scs(deb, ct, fmt.Sprintf("%02d.png", slide))
	done(slide)
}
