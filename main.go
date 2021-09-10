package main

import (
	"test-spider/engine"
	"test-spider/parse"
)

func main() {
	/*engine.Run(engine.Request{
		Url: "https://book.douban.com/tag/%E4%BA%92%E8%81%94%E7%BD%91",
		ParseFunc: parse.ParseBookList,
	})
	*/
	engine.Run(engine.Request_q{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseTag,
	})
}
