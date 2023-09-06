package tongda

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"regexp"
	"ui/common"
)

func UI(w fyne.Window) *fyne.Container {
	common.Txmuentry.Wrapping = fyne.TextWrapWord //自动换行
	common.Txmuentry.Disable()
	common.Tcmuentry.Disable()
	common.Tqmuentry.Disable()
	common.Txmuentry.SetText(`
这里是介绍
`)
	common.Tcmuentry.SetText(`请选择利用的漏洞名称`)
	common.Tqmuentry.SetText(`请输入利用的目标地址`)
	common.Tcentry.SetText(`whoami`)
	button := widget.NewButton("验证", func() {
		exp, _ := regexp.Compile(`http[s]{0,1}://(([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6})|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\&%_\./-~-]*)?`)
		if !exp.MatchString(common.Tentry.Text) {
			dia := dialog.NewCustom("错误", "确认", widget.NewLabel("url不符合要求 | 示例:http/https//:www.\n请输入正确的url"), w)
			dia.Resize(fyne.Size{200, 100})
			dia.Show()
			dia.Refresh()
		}
	})
	buttonc := widget.NewButton("执行", func() {})
	buttonx := widget.NewButton("解密", func() {})
	selects := widget.NewSelect([]string{
		"一键检测",
		"1",
		"2",
		"3",
		"4",
		"5",
	}, func(s string) {
		switch s {
		case "一键检测":
			button.OnTapped = func() {
				text := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%v\n\n%s%s\n\n", Exp1(common.Tentry.Text), Exp2(common.Tentry.Text), Exp3(common.Tentry.Text), Exp4(common.Tentry.Text), Exp5(common.Tentry.Text), Exp6(common.Tentry.Text))
				common.Txmuentry.Enable()
				common.Txmuentry.SetText(text)
				button.Refresh()
			}
		}
	})
	selects.Selected = "下拉选择"
	tabs := container.NewAppTabs(
		container.NewTabItem("信息", common.Txmuentry),
		container.NewTabItem("命令执行", container.NewBorder(container.NewBorder(nil, nil, common.Code, buttonc, common.Tcentry), nil, nil, nil, common.Tcmuentry)),
		container.NewTabItem("数据解密", container.NewBorder(container.NewBorder(nil, nil, common.Url, buttonx, common.Tqcentry), nil, nil, nil, common.Tqmuentry)),
	)
	border := container.NewBorder(nil, nil, nil, selects, container.NewBorder(nil, nil, common.Url, button, common.Tentry))
	rs := container.NewBorder(border, nil, nil, nil, tabs)
	return rs
}
