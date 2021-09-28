package main

import (
	"github.com/go-vgo/robotgo"
	"gocv.io/x/gocv"
)

func main() {
	window := gocv.NewWindow("Hello")

	img := robotgo.ToImage(robotgo.CaptureScreen())
	mat, err := gocv.ImageToMatRGB(img)
	if err != nil {
		return
	}
	for {
		window.IMShow(mat)
		window.WaitKey(1)
	}
}