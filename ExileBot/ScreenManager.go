package ExileBot

import "github.com/go-vgo/robotgo"

var WINDOW_NAME = ""

func getWindow() {
	fpid, err := robotgo.FindIds(WINDOW_NAME)
	if err == nil {
		if len(fpid) > 0 {
		}
	}
}
