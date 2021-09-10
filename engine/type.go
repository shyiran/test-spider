package engine
//解析结果
type ParseResult struct {
	Requesrts []Request_q
	Items     []interface{}
}
//请求
//URL  地址
//处理方法函数
type Request_q struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
