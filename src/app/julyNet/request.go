package julyNet

import (
	"net/http"
	"sync"
	"time"
)

type CrawlRequest struct {

	UUID          string         //请求标识，作为映射
	Url           string         //请求URL
	Method        string         //请求方法
	PostData      string         //post请求数据
	Header        http.Header    //请求头
	UseCookie     bool           //是否使用cookeie
	DialTimeout   time.Duration  //创建超时时间
	ConnTimeout   time.Duration  //连接超时时间
	RetryTimes    int            //重试次数
	RetryPause    time.Duration  //重试延时
	RedirectTimes string         //重定向次数
	Proxy         string         //请求代理
	NotFilter     bool           //是否入队校验
	Priority      int            //优先级
	Once          sync.Once      //Once控制，避免重复
	DownloaderEngine int         //下载引擎
}
