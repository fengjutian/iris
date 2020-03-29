package main

import (
	"fmt"
	qrcode "github.com/yeqown/go-qrcode"
)

func main() {
	qrc, err := qrcode.New("http://www.baidu.com")

	if err != nil {
		fmt.Printf("coulid not generate QRCode: %v", err)
	}

	if err := qrc.Save("./qr-code.jpeg");
	err != nil {
		fmt.Printf(err)
	}
}