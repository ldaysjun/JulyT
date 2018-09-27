package main

import (
	"app/JulySpider"
	"app/julyTaskPool"
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet() {

	resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
	if err != nil {
		// handle error
		fmt.Println("怎么那么多事情")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}
//var complete chan int = make(chan int)

func parse(body string)  {
	fmt.Println(body)
}


//测试任务池
func main()  {


	//
	//fmt.Println("开始测试")
	//engine:=julyEngine.NewEngine()
	//engine.Run()
	//fmt.Println("初始化完毕")
	////
	//for i:=0;i<100;i++ {
	//	req := new(julyNet.CrawlRequest)
	//	fmt.Println()
	//	req.Url ="wwww:"+strconv.Itoa(i)
	//	req.Proxy = "web-proxy.tencent.com:8080"
	//
	//	spider:=&JulySpider.Spider{
	//		Request:req,
	//	}
	//	spider.Registered()
	//}

	spider := new(JulySpider.Spider)
	spider.Parse = JulySpider.Parse(parse)


	spider.Parse.Parse("测试看看")

	//
	//
	//
	//
	//req := new(julyNet.CrawlRequest)
	//fmt.Println()
	//req.Url ="wwww:"+strconv.Itoa(11)
	//engine.PushRequestToQueue(req)
	//
	//
	//julyNet.DownLoad(julyNet.CrawlRequest{
	//	Url:"http://lastdays.cn/",
	//})
	//
	//fmt.Println("//*[@id=\"main\"]/article[1]/header/h1/a")
	complete <- 0
}

//func main() {
//	proxy := func(_ *http.Request) (*url.URL, error) {
//		return url.Parse("")
//	}
//
//	transport := &http.Transport{Proxy: proxy}
//
//	client := &http.Client{Transport: transport}
//	resp, err := client.Get("http://lastdays.cn/")
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println(resp)
//}


type f func() error
var complete chan int = make(chan int)
//
func OK() {

	queue := make(chan string, 4)

	//range函数遍历每个从通道接收到的数据，因为queue再发送完两个
	//数据之后就关闭了通道，所以这里我们range函数在接收到两个数据
	//之后就结束了。如果上面的queue通道不关闭，那么range函数就不
	//会结束，从而在接收第三个数据的时候就阻塞了。

	fmt.Println("可以")
	queue<-"sss"
	queue<-"hello"
	//fmt.Println("可以")
	//queue<-f
	//queue<-f

	fmt.Println("测试")

	go func() {
		fmt.Println("运行到这里1")
		//complete <- 0
		for f := range queue {
			if f == "" {
				return
			}
			println("队列数量",len(queue))
			//f()
			fmt.Println(f)
		}

	}()
	queue<-"hello w"

	//time.Sleep(3*time.Second)

	//<-complete
	fmt.Println("运行到这里5")

	//for i:=0; i<10; i++ {
	//	if i == 4{
	//		queue<-nil
	//	}else {
	//		queue<-f
	//	}
	//}

	complete <- 0


}

//func test(request julyNet.CrawlRequest)  {
//	fmt.Println("这里")
//	fmt.Println(request.Url)
//}

func dome() string {
	return "hello world"
}

func test(pool *julyTaskPool.TaskPool) (test string) {
	pool.SubmitTask(func() error {
		fmt.Println("执行")
		test = dome()
		return nil
	})
	fmt.Println("吃屎")
	return test
}
