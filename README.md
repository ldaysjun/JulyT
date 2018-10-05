# JulyT
一个高效的爬虫框架

## 背景

​	一直准备做一个iOS优质文章的聚合，每天定向爬取大V们的博客。所以就萌生了尝试打造一个通用爬虫框架的想法。加上近期开始golang的学习，所以选择使用go来写。

### 初识

​	我给它起名叫JulyT，目前0.1版本支持Xpath数据解析，批量任务并发。简单的规则编写。就可以完成定向数据的提取.

例如：实现抓取列表，再抓取列表页详情数据，接着翻页继续

```
func rule(node *Xpath.Node,spider *JulySpider.Spider)  {
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
		spider.RunNextStep("http://lastdays.cn"+url,rule)
	}
}
```



## 详细设计



### 1.1 任务池

​	为每一个爬虫实例提供独立的运行空间。自动调度，自动回收空闲任务节点，任务节点复用。提供最底层的任务环境

[传送门：任务池设计]()

### 1.2 调度器

​	管理所有请求，实现请求优先级调度。过滤重复请求。

[传送门：调度器设计]()

### 1.3 下载器

​	提供高并发的HTML下载。

[传送门：下载器设计]()

### 1.4 引擎

​	处理数据流，控制各个模块之间的调度。监控所有请求流程

[传送门：引擎设计]()

### 1.5 spider

​	爬虫实例，支持规则自定义。

[传送门：spider设计]()



## 结构图



![](1.png)



​	
