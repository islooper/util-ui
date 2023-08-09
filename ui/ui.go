package ui

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/ncruces/zenity"
	"math/rand"
	"util-ui/domian"
)

func InitUi() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(domian.MainLogo)
	strLengthItem := systray.AddMenuItem("字符长度计算", "")
	randStrItem := systray.AddMenuItem("随机字符串生成", "")
	for {
		select {
		case <-strLengthItem.ClickedCh:
			strLength()

		case <-randStrItem.ClickedCh:
			randStr()
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

// 生成随机字符串
func randStr() {
	inputStr, _ := zenity.Entry("需求的字符长度:",
		zenity.Title("随机字符串生成"))

	l := stringToInt(inputStr)
	if l <= 0 {
		return
	}

	randStrData := randomString(l)

	zenity.Info(randStrData, zenity.Title("随机字符串生成"),
		zenity.Width(200), zenity.Height(100), zenity.NoIcon)
}

// 随机生成字符串
func randomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// string 转 int
func stringToInt(str string) int {
	var result int
	for i := 0; i < len(str); i++ {
		result = result*10 + int(str[i]-'0')
	}
	return result
}
