package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime/debug"
	"strings"

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
