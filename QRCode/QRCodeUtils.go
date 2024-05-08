package myqrcode

import (
    "fmt"
    "image"
    _ "image/jpeg" // 导入JPEG解码器
    _ "image/png"  // 导入PNG解码器
    "os"

    "github.com/makiuchi-d/gozxing"
    "github.com/makiuchi-d/gozxing/qrcode"
)

func GetUrlFromQRCode(imagePath string) (string, error) {
    // 打开包含二维码的图片文件
    file, err := os.Open(imagePath)
    if err != nil {
        fmt.Println("os.Open: ", err)
        return "", err
    }
    defer file.Close()

    // 解码图片文件
    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println("image.Decode: ", err)
        return "", err
    }

    // 将image.Image转换为*gozxing.LuminanceSource
    source := gozxing.NewLuminanceSourceFromImage(img)
    // 将*gozxing.LuminanceSource转换为*gozxing.BinaryBitmap
    bitmap, err := gozxing.NewBinaryBitmap(gozxing.NewHybridBinarizer(source))
	if err != nil {
		fmt.Println("NewBinaryBitmap: ", err)
        return "", err
	}
    // 创建一个二维码读取器
    qrReader := qrcode.NewQRCodeReader()

    // 使用读取器从BinaryBitmap中解析二维码
    result, err := qrReader.Decode(bitmap, nil)
    if err != nil {
        fmt.Println("qrReader.Decode: ", err)
        return "", err
    }

    // 输出解析结果
    fmt.Println("Decoded QR Code content:", result.String())
	return result.String(), err
}
