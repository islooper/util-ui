package ui

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
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

	// AES 加密
	aesItem := systray.AddMenuItem("AES", "")
	_ = aesItem.AddSubMenuItem("Key", "")
	_ = aesItem.AddSubMenuItem("Iv", "")
	_ = aesItem.AddSubMenuItem("Data", "")
	_ = aesItem.AddSubMenuItem("Encrypt", "")
	_ = aesItem.AddSubMenuItem("Decrypt", "")

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

	err := zenity.Info(fmt.Sprintf("长度为：%d", length), zenity.Title("字符长度统计"),
		zenity.Width(200), zenity.Height(100), zenity.NoIcon)
	if err != nil {
		return
	}
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

	err := zenity.Info(randStrData, zenity.Title("随机字符串生成"),
		zenity.Width(200), zenity.Height(100), zenity.NoIcon)
	if err != nil {
		return
	}
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

// AESEncrypt AES 加密
func AESEncrypt(data string, key string, iv string) (res string, err error) {

	//切换 AES mode
	block, cbcErr := aes.NewCipher([]byte(key))
	if cbcErr != nil {
		fmt.Printf("AESEncrypt CBC NewCipher err %v\n", err)
		return "", cbcErr
	}

	//填充数据
	originData := paddingData(block, []byte(data))

	//切换 AES mode
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)

	res = hex.EncodeToString(crypted)

	return res, err
}

func paddingData(block cipher.Block, data []byte) []byte {
	var origData []byte
	blockSize := block.BlockSize()
	origData = PKCS7Padding(data, blockSize)
	return origData
}

// PKCS7Padding 使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// AESDecrypt AES 解密
func AESDecrypt(data string, key string, iv string) (res []byte, err error) {

	//解析数据
	encrypted, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}

	//切换 AES mode
	block, cbcErr := aes.NewCipher([]byte(key))
	if cbcErr != nil {
		fmt.Println("AESDecrypt CBC NewCipher err %v", err)
		return nil, cbcErr
	}

	bs := block.BlockSize()
	if len(encrypted)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	res = make([]byte, len(encrypted))
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	blockMode.CryptBlocks(res, encrypted)

	res = unpaddingData(res)

	return res, err
}

func unpaddingData(data []byte) []byte {
	return PKCS7UnPadding(data)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
