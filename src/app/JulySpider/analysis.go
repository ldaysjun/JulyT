package JulySpider

type Parser interface {
	//解析处理
	Parse(body string)
}

type Parse func(body string)
func (parse Parse)Parse(body string)  {
	parse(body)
}



