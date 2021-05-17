package toolx

import (
	"fmt"
	"strings"
)

type Spider interface {
	BuildIconLink(url string) string
}

type Article struct {
	url string
}

const icoFormat = "[<img src=\"%s\" weight=\"20\" height=\"20\">](%s)"

const (
	wechatIcoUrl = "https://mp.weixin.qq.com/favicon.ico"
	wechatHost   = "mp.weixin.qq.com"

	juejinIcoUrl = "https://juejin.cn/favicon.ico"
	juejinHost   = "juejin.cn"

	zhihuIcoUrl = "https://static.zhihu.com/heifetz/favicon.ico"
	zhihuHost   = "zhuanlan.zhihu.com"

	tencentIcoUrl = "https://cloud.tencent.com/favicon.ico"
	tencentHost   = "cloud.tencent.com"

	csdnIcoUrl = "https://blog.csdn.net/favicon.ico"
	csdnHost   = "blog.csdn.net"

	jianshuIcoUrl = "https://www.jianshu.com/favicon.ico"
	jianshuHost   = "jianshu.com"

	infoqIcoUrl = "https://static001.infoq.cn/static/infoq/favicon/favicon-32x32.png"
	infoqHost   = "infoq.cn"
)

const (
	IcoUnknown = 0
	IcoWechat  = 1
	IcoJuejin  = 2
	IcoZhihu   = 3
	IcoTencent = 4
	IcoCsdn    = 5
	IcoJianshu = 6
	IcoInfoq   = 7

	IcoMin = 1
	IcoMax = 7
)

func GetIconLink(articleUrl string) (string, int) {
	if strings.Contains(articleUrl, wechatHost) {
		return fmt.Sprintf(icoFormat, wechatIcoUrl, articleUrl), IcoWechat
	} else if strings.Contains(articleUrl, juejinHost) {
		return fmt.Sprintf(icoFormat, juejinIcoUrl, articleUrl), IcoJuejin
	} else if strings.Contains(articleUrl, zhihuHost) {
		return fmt.Sprintf(icoFormat, zhihuIcoUrl, articleUrl), IcoZhihu
	} else if strings.Contains(articleUrl, tencentHost) {
		return fmt.Sprintf(icoFormat, tencentIcoUrl, articleUrl), IcoTencent
	} else if strings.Contains(articleUrl, csdnHost) {
		return fmt.Sprintf(icoFormat, csdnIcoUrl, articleUrl), IcoCsdn
	} else if strings.Contains(articleUrl, jianshuHost) {
		return fmt.Sprintf(icoFormat, jianshuIcoUrl, articleUrl), IcoJianshu
	} else if strings.Contains(articleUrl, infoqHost) {
		return fmt.Sprintf(icoFormat, infoqIcoUrl, articleUrl), IcoInfoq
	}
	return "", IcoUnknown
}