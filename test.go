package main

import (
	"app/JulySpider/Xpath"
	"app/julyEngine"
	"app/julyNet"
	"fmt"
	"io/ioutil"
	"time"

	//"net/http"
	"strings"
)

func httpGet() {

	req := new(julyNet.CrawlRequest)
	req.Url = "http://lastdays.cn/"
	req.NotFilter =true
	req.UUID = "1"
	down := julyNet.NewDownLoad()
	resp, _ := down.DownLoad(req)

	//resp, err := http.Get("http://lastdays.cn/")
	//if err != nil {
	//	// handle error
	//	fmt.Println("怎么那么多事情")
	//}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	go parseddd(string(body))
	go parsedd(string(body))
	//fmt.Println(string(body))
}
//var complete chan int = make(chan int)

func parseddd(body string)  {
	fmt.Println("解析1")

	inputReader := strings.NewReader(string(body))
	node,err:=Xpath.ParseHTML(inputReader)

	if err != nil {
		fmt.Println("xmlpath parse file failed!!!")
		return
	}

	path := Xpath.MustCompile("//*[@id=\"main\"]/article[1]/header/h1/a")
	result,_ := path.String(node)
	fmt.Println("输出结果：",result)


	//inputReader2 := strings.NewReader(string(body))
	//node2,err2:=Xpath.ParseHTML(inputReader2)
	//
	//if err2 != nil {
	//	fmt.Println("xmlpath parse file failed!!!")
	//	return
	//}
	//
	//path2 := Xpath.MustCompile("//*[@id=\"main\"]/article[1]/header/h1/a")
	//fmt.Println(path2.String(node2))

}


func parsedd(body string)  {
	fmt.Println("解析2")

	inputReader := strings.NewReader(string(body))
	node,err:=Xpath.ParseHTML(inputReader)


	if err != nil {
		fmt.Println("xmlpath parse file failed!!!")
		return
	}

	path := Xpath.MustCompile("//*[@id=\"main\"]/article[1]/header/h1/a")
	fmt.Println(path.String(node))
	//fmt.Println("结束llll")

}

var complete chan int = make(chan int)

//测试任务池
func main()  {
	fmt.Println("开始测试")
	engine := julyEngine.Run()
	t:=3
	go func() {
		time.Sleep(2*time.Second)
		fmt.Println("关闭")
		engine.CloseEngine()

		//v:=<-complete
		//fmt.Println(v)
	}(t)




	//for i:=0;i<20; i++ {
	//	req2 := new(julyNet.CrawlRequest)
	//	req2.UUID = strconv.Itoa(i)
	//	req2.Url = "http://lastdays.cn/"
	//	req2.NotFilter = true
	//
	//	spider2 := new(JulySpider.Spider)
	//	spider2.SpiderName = "测试1"
	//	spider2.Parse = JulySpider.Parse(parseddd)
	//	spider2.Request = req2
	//	spider2.Registered()
	//}

	complete <- 0
}




