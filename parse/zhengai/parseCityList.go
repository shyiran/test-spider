package zhengai

import (
	"regexp"
	"test-spider/engine"
)

const cityListRe = `(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult{
	re:=regexp.MustCompile(cityListRe)

	matches:= re.FindAllSubmatch(contents,-1)
	result:=engine.ParseResult{}

	for _,m :=range matches{
		//result.Items = append(result.Items,string(m[2]))
		result.Requesrts  = append(
			result.Requesrts,engine.Request_q{
				Url:string(m[1]),
				ParseFunc:ParseCity,
			})
	}
	return result
}
