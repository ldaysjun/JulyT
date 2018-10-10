package JulySpider

import (
	"app/JulySpider/Xpath"
)

type AnalysisInterface interface {
	Parse(node *Xpath.Node,spider *Spider) //解析处理,解析入口
	Next(node *Xpath.Node,spider *Spider)
}

type Analysis func(node *Xpath.Node,spider *Spider)

func (analysis Analysis)Parse(node *Xpath.Node,spider *Spider)  {
	analysis(node,spider)
}

func (analysis Analysis)Next(node *Xpath.Node,spider *Spider)  {
	analysis(node,spider)
}









