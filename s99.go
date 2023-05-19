package main

import (
	"fmt"
	"path/filepath"
	"time"

	dp "github.com/chromedp/chromedp"
)

func s99(slide int) {
	var (
		params = conf.P["99"]
		sel    string
		tit    string
		err    error
	)
	stdo.Println(params)
	ct, ca := chrome()
	defer ca()
	dp.Run(ct,
		EmulateViewport(1920, 1080),
		dp.Navigate(params[0]),
		dp.Sleep(sec),
	)
	scs(slide, ct, fmt.Sprintf("%02d get.png", slide))
	sel = "input[name='ar-user-name']"
	ct2, ca2, err := RunTO(ct, sec*13,
		dp.Click(sel, dp.ByQuery, dp.NodeVisible, dp.NodeEnabled),
		dp.Sleep(ms),
		dp.SendKeys(sel, params[1], dp.ByQuery, dp.NodeVisible),
		dp.Sleep(ms),
	)
	defer ca2()
	stdo.Println(sel, err)
	if err == nil {
		ex(slide, dp.Run(ct2,
			dp.Click("button[type=submit]", dp.ByQuery, dp.NodeEnabled),
			dp.Sleep(ms),
		))
		scs(slide, ct2, fmt.Sprintf("%02d %s.png", slide, "ar-user-name"))
		sel = "input[name='ar-user-password']"
		ex(slide, dp.Run(ct2,
			dp.Click(sel, dp.ByQuery, dp.NodeVisible, dp.NodeEnabled),
			dp.Sleep(ms),
			dp.SendKeys(sel, params[2], dp.ByQuery, dp.NodeVisible),
			dp.Sleep(ms),
		))
		ex(slide, dp.Run(ct2,
			dp.Click("button[type='submit']", dp.ByQuery, dp.NodeEnabled),
			dp.Sleep(ms),
		))
		scs(slide, ct2, fmt.Sprintf("%02d %s.png", slide, "ar-user-password"))
	}
	tit = "Редактировать"
	sel = "div.multiBtnInner_xbp:nth-child(1)"
	ex(slide, dp.Run(ct,
		dp.Click(sel, dp.ByQuery, dp.NodeEnabled),
		dp.Sleep(ms),
	))
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))

	tit = "Удалить"
	se := "button.menu-button_J9B"
	sel = "button.align-left_-232488494:nth-child(3)"
	ct3, ca3, err := RunTO(ct, sec*7,
		dp.Click(se, dp.ByQuery, dp.NodeVisible),
		dp.Sleep(ms),
		dp.Click(sel, dp.ByQuery, dp.NodeVisible),
		dp.Sleep(ms),
	)
	defer ca3()
	stdo.Println(tit, err)
	scs(slide, ct3, fmt.Sprintf("%02d %s.png", slide, tit))
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
	if false {
		Run(ct, sec*7,
			dp.WaitVisible(fmt.Sprintf("//*[contains(text(),'%s')]", tit)),
			// dp.WaitVisible(tit, dp.BySearch), //no node
		)
		scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))
		stdo.Println(tit, err)
	} else {
		ct4, ca4, err := RunTO(ct, sec*7,
			dp.WaitVisible(fmt.Sprintf("//*[contains(text(),'%s')]", tit)),
		)
		defer ca4()
		scs(slide, ct4, fmt.Sprintf("%02d %s.png", slide, tit))
		stdo.Println(tit, err)
	}

	tit = "отменена"
	ct5, ca5, err := RunTO(ct, sec*3,
		dp.WaitVisible(fmt.Sprintf("//*[contains(text(),'%s')]", tit)),
	)
	defer ca5()
	stdo.Println(tit, err)
	if err == nil {
		scs(slide, ct5, fmt.Sprintf("%02d %s.png", slide, tit))
		Scanln()
		return
	}
	tit = "завершена"
	ct6, ca6, err := RunTO(ct, sec*7,
		dp.WaitVisible(fmt.Sprintf("//*[contains(text(),'%s')]", tit)),
	)
	defer ca6()
	stdo.Println(tit, err)
	if err != nil {
		scs(slide, ct6, fmt.Sprintf("%02d %s.png", slide, "Загрузка НЕ завершена"))
		Scanln()
		return
	}
	scs(slide, ct, fmt.Sprintf("%02d %s.png", slide, tit))
	time.Sleep(sec * 7)
	ex(slide, dp.Run(ct,
		dp.Click("div.multiBtnInner_xbp:nth-child(4)", dp.ByQuery, dp.NodeEnabled),
		dp.Sleep(sec),
	))
	scs(deb, ct, fmt.Sprintf("%02d.png", slide))
	done(slide)
}
