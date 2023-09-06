package fanwei

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"regexp"
	"ui/common"
)

func Fui(w fyne.Window) *fyne.Container {
	common.Fxmuentry.Wrapping = fyne.TextWrapWord //自动换行
	common.Fxmuentry.Disable()
	common.Fcmuentry.Disable()
	common.Fqmuentry.Disable()
	common.Fxmuentry.SetText(`
1.泛微OA weaver.common.Ctrl 任意文件上传漏洞
2.泛微云桥 e-Bridge 任意文件读取
3.泛微OA Bsh 远程代码执行漏洞 CNVD-2019-32204
4.泛微OA e-cology 数据库配置信息泄漏漏洞
5.泛微OA WorkflowCenterTreeData接口SQL注入(仅限oracle数据库) CNVD-2019-34241
6.泛微OA V9 任意文件上传
7.泛微OA V8 SQL注入漏洞
`)
	common.Fcmuentry.SetText(`请选择利用的漏洞名称`)
	common.Fqmuentry.SetText(`请输入利用的目标地址`)
	common.Fcentry.SetText(`whoami`)
	button := widget.NewButton("验证", func() {
		exp, _ := regexp.Compile(`http[s]{0,1}://(([a-zA-Z0-9\._-]+\.[a-zA-Z]{2,6})|([0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\&%_\./-~-]*)?`)
		if !exp.MatchString(common.Fentry.Text) {
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
		"泛微 weaver任意文件上传漏洞",
		"泛微 e-Bridge 任意文件读取",
		"泛微 Bsh远程代码执行漏洞",
		"泛微 数据库配置信息泄漏漏洞",
		"泛微 Workflow接口SQL注入",
		"泛微 V9 任意文件上传",
		"泛微 V8 SQL注入漏洞",
	}, func(s string) {
		switch s {
		case "一键检测":
			button.OnTapped = func() {
				text := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%v\n\n%s%s\n\n", Exp1(common.Fentry.Text), Exp2(common.Fentry.Text), Exp3(common.Fentry.Text), Exp4(common.Fentry.Text), Exp5(common.Fentry.Text), Exp6(common.Fentry.Text))
				common.Fxmuentry.Enable()
				common.Fxmuentry.SetText(text)
				button.Refresh()
			}
		}
	})
	selects.Selected = "下拉选择"
	tabs := container.NewAppTabs(
		container.NewTabItem("信息", common.Fxmuentry),
		container.NewTabItem("命令执行", container.NewBorder(container.NewBorder(nil, nil, common.Code, buttonc, common.Fcentry), nil, nil, nil, common.Fcmuentry)),
		container.NewTabItem("数据解密", container.NewBorder(container.NewBorder(nil, nil, common.Url, buttonx, common.Fqcentry), nil, nil, nil, common.Fqmuentry)),
	)
	border := container.NewBorder(nil, nil, nil, selects, container.NewBorder(nil, nil, common.Url, button, common.Fentry))
	rs := container.NewBorder(border, nil, nil, nil, tabs)
	return rs
}
