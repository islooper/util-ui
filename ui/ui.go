package ui

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/ncruces/zenity"
	"util-ui/domian"
)

func InitUi() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(domian.MainLogo)
	strLengthItem := systray.AddMenuItem("字符长度计算", "")

	for {
		select {
		case <-strLengthItem.ClickedCh:
			strLength()
		}
	}
}

func onExit() {
	systray.Quit()
}

// strLength 计算字符长度
func strLength() {
	inputStr, _ := zenity.Entry("请输入字符:",
		zenity.Title("字符长度统计"))

	length := len(inputStr)
	if length == 0 {
		return
	}

	zenity.Info(fmt.Sprintf("长度为：%d", length), zenity.Title("字符长度统计"),
		zenity.Width(200), zenity.Height(100), zenity.NoIcon)
}
