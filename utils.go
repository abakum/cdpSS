package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/xlab/closer"
)

func sErr(s string, err error) string {
	if err != nil {
		return err.Error()
	}
	return s
}
func i2p(v int) (fn string) {
	fn = fmt.Sprintf("%02d.jpg", v)
	if v == 97 {
		fn = mov
	}
	fn = filepath.Join(root, fn)
	return
}

type ss []byte

func (i ss) write(fileName string) {
	fullName := filepath.Join(root, doc, fileName)
	jpg := strings.HasSuffix(fileName, ".jpg")
	if jpg {
		fullName = filepath.Join(root, fileName)
	}
	err := os.WriteFile(fullName, []byte(i), 0644)
	if err != nil {
		return
	}
	// exec.Command("rundll32", "url.dll,FileProtocolHandler", fullName).Run()
	exec.Command("cmd", "/c", "start", "chrome", fullName).Run()
}

type ii struct {
	img image.Image
	err error
	png []byte
}

func ssII(i *ii) (o *ii) { //Beginning of the pipe
	o = i
	o.img, o.err = png.Decode(bytes.NewReader(i.png))
	return
}

func (i *ii) crop(crop image.Rectangle) (o *ii) { //Middle of the pipe
	o = &ii{err: i.err}
	if i.err != nil {
		return
	}
	type subImager interface {
		SubImage(ir image.Rectangle) image.Image
	}
	// i.img is an Image interface. This checks if the underlying value has a
	// method called SubImage. If it does, then we can use SubImage to crop the
	// image.
	sImg, ok := i.img.(subImager)
	if !ok {
		o.img = i.img
		o.err = fmt.Errorf("image does not support cropping")
		return
	}
	o.img = sImg.SubImage(crop)
	return
}

func (i *ii) write(fileName string) (err error) { //Pipe end
	err = i.err
	if err != nil {
		return
	}
	fullName := filepath.Join(root, doc, fileName)
	jpg := strings.HasSuffix(fileName, ".jpg")
	if jpg {
		fullName = filepath.Join(root, fileName)
	}
	file, err := os.Create(fullName)
	if err != nil {
		return
	}
	defer file.Close()
	if jpg {
		err = jpeg.Encode(file, i.img, &jpeg.Options{Quality: 100})
	} else {
		err = png.Encode(file, i.img)
	}
	if err != nil {
		return
	}
	// err = exec.Command("rundll32", "url.dll,FileProtocolHandler", fullName).Run()
	// err = exec.Command("powershell", "Start-Process", "chrome", "-argumentlist", fullName).Run()
	err = exec.Command("cmd", "/c", "start", "chrome", fullName).Run()
	// err = exec.Command(chromeBin, fullName).Run()
	return
}
func ex(slide int, err error) {
	if err != nil {
		exit = slide
		stdo.Println(src(8), err.Error())
		closer.Close()
	}
}
func src(deep int) (s string) {
	s = string(debug.Stack())
	// for k, v := range strings.Split(s, "\n") {
	// 	stdo.Println(k, v)
	// }
	s = strings.Split(s, "\n")[deep]
	s = strings.Split(s, " +0x")[0]
	_, s = path.Split(s)
	return
}

func Screenshot(sel interface{}, picbuf *[]byte, quality int, opts ...chromedp.QueryOption) chromedp.QueryAction {
	if picbuf == nil {
		panic("picbuf cannot be nil")
	}

	return chromedp.QueryAfter(sel, func(ctx context.Context, execCtx runtime.ExecutionContextID, nodes ...*cdp.Node) error {
		if len(nodes) < 1 {
			return fmt.Errorf("selector %q did not return any nodes", sel)
		}

		// get box model
		var clip page.Viewport
		if err := callFunctionOnNode(ctx, nodes[0], getClientRectJS, &clip); err != nil {
			return err
		}
		// The "Capture node screenshot" command does not handle fractional dimensions properly.
		// Let's align with puppeteer:
		// https://github.com/puppeteer/puppeteer/blob/bba3f41286908ced8f03faf98242d4c3359a5efc/src/common/Page.ts#L2002-L2011
		x, y := math.Round(clip.X), math.Round(clip.Y)
		clip.Width, clip.Height = math.Round(clip.Width+clip.X-x), math.Round(clip.Height+clip.Y-y)
		clip.X, clip.Y = x, y

		// The next comment is copied from the original code.
		// This seems to be necessary? Seems to do the right thing regardless of DPI.
		clip.Scale = 1

		format := page.CaptureScreenshotFormatPng
		if quality != 100 {
			format = page.CaptureScreenshotFormatJpeg
		}
		// take screenshot of the box
		buf, err := page.CaptureScreenshot().
			WithFormat(format).
			WithQuality(int64(quality)).
			WithCaptureBeyondViewport(true).
			WithFromSurface(true).
			WithClip(&clip).
			Do(ctx)
		if err != nil {
			return err
		}

		*picbuf = buf
		return nil
	}, append(opts, chromedp.NodeVisible)...)
}
func FullScreenshot(res *[]byte, quality int, clip *page.Viewport) chromedp.EmulateAction {
	if res == nil {
		panic("res cannot be nil")
	}
	if clip == nil {
		panic("clip cannot be nil")
	}
	return chromedp.ActionFunc(func(ctx context.Context) error {
		format := page.CaptureScreenshotFormatPng
		if quality != 100 {
			format = page.CaptureScreenshotFormatJpeg
		}

		// capture screenshot
		var err error
		*res, err = page.CaptureScreenshot().
			WithCaptureBeyondViewport(true).
			WithFromSurface(true).
			WithFormat(format).
			WithQuality(int64(quality)).
			WithClip(clip).
			Do(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}

func getClientRect(sel interface{}, clip *page.Viewport, opts ...chromedp.QueryOption) chromedp.QueryAction {
	if clip == nil {
		panic("clip cannot be nil")
	}

	return chromedp.QueryAfter(sel, func(ctx context.Context, execCtx runtime.ExecutionContextID, nodes ...*cdp.Node) error {
		if len(nodes) < 1 {
			return fmt.Errorf("selector %q did not return any nodes", sel)
		}

		// get box model
		var Viewport page.Viewport
		if err := callFunctionOnNode(ctx, nodes[0], getClientRectJS, &Viewport); err != nil {
			return err
		}
		// The "Capture node screenshot" command does not handle fractional dimensions properly.
		// Let's align with puppeteer:
		// https://github.com/puppeteer/puppeteer/blob/bba3f41286908ced8f03faf98242d4c3359a5efc/src/common/Page.ts#L2002-L2011
		x, y := math.Round(Viewport.X), math.Round(Viewport.Y)
		Viewport.Width, Viewport.Height = math.Round(Viewport.Width+Viewport.X-x), math.Round(Viewport.Height+Viewport.Y-y)
		Viewport.X, Viewport.Y = x, y

		// The next comment is copied from the original code.
		// This seems to be necessary? Seems to do the right thing regardless of DPI.
		Viewport.Scale = 1

		*clip = Viewport
		return nil
	}, append(opts, chromedp.NodeVisible)...)
}

func callFunctionOnNode(ctx context.Context, node *cdp.Node, function string, res interface{}, args ...interface{}) error {
	r, err := dom.ResolveNode().WithNodeID(node.NodeID).Do(ctx)
	if err != nil {
		return err
	}
	err = chromedp.CallFunctionOn(function, res,
		func(p *runtime.CallFunctionOnParams) *runtime.CallFunctionOnParams {
			return p.WithObjectID(r.ObjectID)
		},
		args...,
	).Do(ctx)

	if err != nil {
		return err
	}

	// Try to release the remote object.
	// It will fail if the page is navigated or closed,
	// and it's okay to ignore the error in this case.
	_ = runtime.ReleaseObject(r.ObjectID).Do(ctx)

	return nil
}

var getClientRectJS = `function getClientRect() {
	const e = this.getBoundingClientRect(),
	  t = this.ownerDocument.documentElement.getBoundingClientRect();
	return {
	  x: e.left - t.left,
	  y: e.top - t.top,
	  width: e.width,
	  height: e.height,
	};
}`

func clip(X, Y, Width, Height float64) *page.Viewport {
	clip := page.Viewport{
		X:      X,
		Y:      Y,
		Width:  Width,
		Height: Height,
		Scale:  1,
	}
	return &clip
}

func scs(slide int, ct1 context.Context, fn string) {
	bytes := []byte{}
	if deb == slide {
		if chromedp.Run(ct1, chromedp.FullScreenshot(&bytes, 100)) == nil {
			ss(bytes).write(fn)
		}
	}
}

func chrome() (ctTab context.Context, caTab context.CancelFunc) {
	ctExe, _ := chromedp.NewExecAllocator(ctRoot, options...)
	ctTab, caTab = chromedp.NewContext(ctExe)
	// first tab create browser instance
	if !mb {
		ex(deb, chromedp.Run(ctTab,
			chromedp.EmulateViewport(1920, 1080),
			chromedp.Navigate("about:blank"),
		))
		time.AfterFunc(time.Second*3, func() {
			// close empty tab
			chromedp.Run(ctTab, chromedp.Evaluate("window.close();", nil))
		})
	}
	return
}
func exeFN() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	_, exeN := filepath.Split(exe)
	return strings.TrimSuffix(exeN, filepath.Ext(exeN)), err
}
