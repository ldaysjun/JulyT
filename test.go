package main

import (
	"app/JulySpider"
	"app/JulySpider/Xpath"
	"app/julyEngine"
	"app/julyNet"
	"fmt"
	"strings"
)

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

}

var complete chan int = make(chan int)

//测试任务池
func main()  {
	fmt.Println("开始测试")
	julyEngine.Run()


	for i:=0;i<20; i++ {
		req2 := new(julyNet.CrawlRequest)
		//req2.UUID = strconv.Itoa(i)
		req2.Url = "http://lastdays.cn/"
		req2.NotFilter = true

		spider2 := new(JulySpider.Spider)
		spider2.SpiderName = "测试1"
		spider2.Parse = JulySpider.Parse(parseddd)
		spider2.Request = req2
		spider2.Registered()
	}

	complete <- 0
}




