package main

import (
	"app/JulySpider"
	"app/JulySpider/Xpath"
	"app/julyEngine"
	"app/julyNet"
	"fmt"
	"strconv"
)

var ch = make(chan int)
func parse(node *Xpath.Node,spider *JulySpider.Spider)  {

	path := Xpath.MustCompile("//*[@id=\"archive-page\"]/section")
	it := path.Iter(node)

	for it.Next() {
			urlPath := Xpath.MustCompile("a/@href")
			url,_:= urlPath.String(it.Node())
			spider.RunNextStep("http://lastdays.cn"+url,analysisData)
			}

	fmt.Println("================一页数据==================")
	nextPath := Xpath.MustCompile("//*[@id=\"page-nav\"]/a[@class=\"extend next\"]/@href")
	if nextPath.Exists(node) {
		url,_ := nextPath.String(node)
		spider.RunNextStep("http://lastdays.cn"+url,parse)
	}
}


func analysisData(node *Xpath.Node,spider *JulySpider.Spider)  {
	if node == nil {
		return
	}
	titlePath := Xpath.MustCompile("//*[@id=\"main\"]/article/header/h1/a")
	title,_:= titlePath.String(node)
	fmt.Println(spider.Request.Url,":",title)
}

var complete chan int = make(chan int)

//测试任务池
func main()  {
	julyEngine.Run()
	for i:=0;i<10; i++ {
		req2 := new(julyNet.CrawlRequest)
		req2.Url = "http://lastdays.cn/archives"
		req2.NotFilter = true
		spider2 := new(JulySpider.Spider)
		spider2.SpiderName = "测试"+strconv.Itoa(i)
		//spider2.Parse = JulySpider.Analysis(parsedd)
		spider2.ParseHandle = parse
		spider2.Request = req2
		spider2.Registered()
	}
	complete <- 0
}


