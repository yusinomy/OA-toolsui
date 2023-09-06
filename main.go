package main

import (
	"os"
	"ui/gui"
)

func main() {
	os.Setenv("FYNE_FONT", "./方正粗黑宋简体.ttf")
	gui.Gui()
	os.Unsetenv("FYNE_FONT")
}
