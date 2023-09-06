package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"ui/poc/fanwei"
	"ui/poc/lanling"
	"ui/poc/nc"
	"ui/poc/seeyon"
	"ui/poc/tongda"
)

func Menu(w fyne.Window) {
	s := fyne.NewMainMenu(
		fyne.NewMenu("设置", fyne.NewMenuItem("设置", func() {})),
		fyne.NewMenu("代理", fyne.NewMenuItem("设置代理", func() {
		})),
		fyne.NewMenu("关于", fyne.NewMenuItem("关于", func() {
			dia := dialog.NewCustom("关于", "确定", widget.NewLabel("本程序是由https://github.com/YusinoMy编写\n此程序只用于安全测试\n请在授权的情况下使用"), w)
			dia.Resize(fyne.Size{200, 100})
			dia.Show()
			dia.Refresh()
		})),
	)
	w.SetMainMenu(s)
}

func Apptable(w fyne.Window) *container.Scroll {
	tap := container.NewAppTabs(
		container.NewTabItem("泛微OA", fanwei.Fui(w)),
		container.NewTabItem("致远OA", seeyon.UI(w)),
		container.NewTabItem("用友OA", nc.UI(w)),
		container.NewTabItem("通达OA", tongda.UI(w)),
		container.NewTabItem("蓝凌OA", lanling.Ui(w)),
	)
	tap.SetTabLocation(container.TabLocationLeading)
	c := container.NewScroll(tap)
	return c
}
