package autochrome

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

var url = "https://pipay.pingan.com/pc/#/cpcpayPc?cpcNoticeNo=80240517081834294268&paySessionId=f1dde7cc0d2f11efbd30b4055d175461"
func AutoScreen() {
    println("start")
    // 创建一个上下文和取消函数
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // 创建一个通道来接收截图的数据
    var buf []byte

    // 访问指定的页面
    if err := chromedp.Run(ctx,
        chromedp.Navigate(url), // 替换为你要访问的网址
        chromedp.Sleep(3*time.Second),           // 等待页面加载完全
        // 等待弹窗出现并点击按钮
        chromedp.WaitVisible(`div.sure_btn`, chromedp.ByQuery),
        chromedp.Click(`div.sure_btn`, chromedp.ByQuery),
        // 等待页面加载完全
        chromedp.Sleep(3*time.Second),
        // 截图并保存到缓冲区
        chromedp.CaptureScreenshot(&buf),
    ); err != nil {
        log.Fatal(err)
        println(err)
    }
    println("访问页面")
    // 将截图数据保存到文件
    if err := os.WriteFile("screenshot.png", buf, 0644); err != nil {
        log.Fatal(err)
        println(err)
    }
    println("over")
}
