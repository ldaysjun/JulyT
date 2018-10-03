package main

import (
	"app/JulySpider"
	"app/JulySpider/Xpath"
	"app/julyEngine"
	"app/julyNet"
	"fmt"
	"strconv"
)

func parseddd(node *Xpath.Node,spider *JulySpider.Spider)  {
	//fmt.Println("解析1")
	//path := Xpath.MustCompile("//*[@id=\"main\"]/article/header/h1/a")
	//result,_ := path.String(node)
	//fmt.Println("输出结果：",result)
}

var ch = make(chan int)
func parsedd(node *Xpath.Node,spider *JulySpider.Spider)  {

	path := Xpath.MustCompile("//*[@id=\"archive-page\"]/section")
	it := path.Iter(node)

	fmt.Println(spider.SpiderName)

	for it.Next() {
			titlePath := Xpath.MustCompile("a/@title")
			urlPath := Xpath.MustCompile("a/@href")
			//fmt.Println(titlePath.String(it.Node()))
			//fmt.Println(urlPath.String(it.Node()))
			url,_:= urlPath.String(it.Node())
			title,_:= titlePath.String(it.Node())
			fmt.Println(spider.Request.UUID+":"+title)
		spider.RunNextStep("http://lastdays.cn"+url,analysisData)
		}
	fmt.Println("==============================================")

	nextPath := Xpath.MustCompile("//*[@id=\"page-nav\"]/a[@class=\"extend next\"]/@href")
	if nextPath.Exists(node) {
		fmt.Println("执行")
		url,_ := nextPath.String(node)
		spider.RunNextStep("http://lastdays.cn"+url,parsedd)
	}
}

//func analysisData(node *Xpath.Node,spider *JulySpider.Spider)  {
//	fmt.Println("lalalal analysisData")
//}

func analysisData(node *Xpath.Node,spider *JulySpider.Spider)  {
	fmt.Println("analysisData")
	if node == nil {
		return
	}
	titlePath := Xpath.MustCompile("//*[@id=\"main\"]/article/header/h1/a")
	title,_:= titlePath.String(node)
	fmt.Println("详情页数据：",title)
}

var complete chan int = make(chan int)

//测试任务池
func main()  {
	fmt.Println("开始测试")
	julyEngine.Run()
	for i:=0;i<10; i++ {
		req2 := new(julyNet.CrawlRequest)
		req2.Url = "http://lastdays.cn/archives"
		req2.NotFilter = true
		spider2 := new(JulySpider.Spider)
		spider2.SpiderName = "测试"+strconv.Itoa(i)
		//spider2.Parse = JulySpider.Analysis(parsedd)
		spider2.ParseHandle = parsedd
		spider2.Request = req2
		spider2.Registered()
		fmt.Println("第一个处理结束========================")
	}
	complete <- 0
}


