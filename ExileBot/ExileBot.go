package ExileBot

import (
	"github.com/go-vgo/robotgo"
	"gocv.io/x/gocv"
)

func Init() {
}

func Run() {
	Init()
	window := gocv.NewWindow("Hello")
	for {
		img := robotgo.ToImage(robotgo.CaptureScreen())
		mat, err := gocv.ImageToMatRGB(img)
		if err != nil {
			return
		}
		window.IMShow(mat)
		if window.WaitKey(1) == 17 {
			break
		}
	}
}
