package seeyon

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
	common.SExmuentry.Wrapping = fyne.TextWrapWord //自动换行
	common.SExmuentry.Disable()
	common.SEcmuentry.Disable()
	common.SEqmuentry.Disable()
	common.SExmuentry.SetText(`
这里是介绍
`)
	common.SEcmuentry.SetText(`请选择利用的漏洞名称`)
	common.SEqmuentry.SetText(`请输入利用的目标地址`)
	common.SEcentry.SetText(`whoami`)
	button := widget.NewButton("验证", func() {
		exp, _ := regexp.Compile(`http[s]{0,1}://(([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6})|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\&%_\./-~-]*)?`)
		if !exp.MatchString(common.SEentry.Text) {
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
				text := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%v\n\n%s%s\n\n", Exp1(common.SEentry.Text), Exp2(common.SEentry.Text), Exp3(common.SEentry.Text), Exp4(common.SEentry.Text), Exp5(common.SEentry.Text), Exp6(common.SEentry.Text))
				common.SExmuentry.Enable()
				common.SExmuentry.SetText(text)
				button.Refresh()
			}
		}
	})
	selects.Selected = "下拉选择"
	tabs := container.NewAppTabs(
		container.NewTabItem("信息", common.SExmuentry),
		container.NewTabItem("命令执行", container.NewBorder(container.NewBorder(nil, nil, common.Code, buttonc, common.SEcentry), nil, nil, nil, common.SEcmuentry)),
		container.NewTabItem("数据解密", container.NewBorder(container.NewBorder(nil, nil, common.Url, buttonx, common.SEqcentry), nil, nil, nil, common.SEqmuentry)),
	)
	border := container.NewBorder(nil, nil, nil, selects, container.NewBorder(nil, nil, common.Url, button, common.SEentry))
	rs := container.NewBorder(border, nil, nil, nil, tabs)
	return rs
}
