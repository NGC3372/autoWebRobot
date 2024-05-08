package main

import (
	"fmt"
	myqrcode "myTestProject/QRCode"
)

//autochrome "myTestProject/AutoChrome"

func main() {
	//autochrome.AutoScreen()
	url, err := myqrcode.GetUrlFromQRCode("./res/qrcode.png")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("url: " + url)
}