package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Gui() {
	a := app.New()
	w := a.NewWindow("OA tool")
	Menu(w)
	w.SetContent(Apptable(w))
	w.Resize(fyne.Size{1000, 800})
	w.CenterOnScreen() //居中
	w.ShowAndRun()
}
