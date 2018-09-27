package julyNet

import (
	"net/http"
	"sync"
	"time"
)

type CrawlRequest struct {
	//请求标识，作为映射
	UUID string
	//请求URL
	Url string
	//请求方法
	Method string
	//post请求数据
	PostData string
	//请求头
	Header http.Header
	//是否使用cookeie
	UseCookie bool
	//创建超时时间
	DialTimeout time.Duration
	//连接超时时间
	ConnTimeout time.Duration
	//重试次数
	RetryTimes int
	//重试延时
	RetryPause time.Duration
	//重定向次数
	RedirectTimes string
	//请求代理
	Proxy string
	//下载引擎
	DownloaderEngine int
	//是否入队校验
	NotFilter bool
	//优先级
	Priority int
	//Once控制，避免重复
	Once sync.Once
}

//type Request interface {
//
//	GetUrl() string
//	//请求方法
//	getMethod() string
//	//post请求数据
//	getPostData() string
//	//请求头
//	getHeader() http.Header
//	//是否使用cookeie
//	getUseCookie() bool
//	//创建超时时间
//	getDialTimeout() string
//	//连接超时时间
//	getConnTimeout() string
//	//重试词素
//	getRetryTimes() int
//	//重定向次数
//	getRedirectTimes() string
//	//请求代理
//	getProxy() string
//	//下载引擎
//	getDownloaderEngine() int
//}
//
//func (self *RequestPackage) GetUrl() string  {
//	return self.Url
//}
