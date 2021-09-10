package main

import (
	"test-spider/engine"
	"test-spider/parse"
	"test-spider/scheduler"
)

func main(){
	//一下是多通道模式

	e:= engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		2,
	}
	e.Run(engine.Request_q{
		Url:       "https://book.douban.com",
		//Url:"http://www.zhenai.com/zhenghun",
		//ParseFunc:zhengai.ParseCity,
		ParseFunc:parse.ParseTag,
	})


	

	/*e:= engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		2,
	}
	e.Run(engine.Request_q{
		Url:       "https://book.douban.com",
		//Url:"http://www.zhenai.com/zhenghun",
		//ParseFunc:zhengai.ParseCity,
		ParseFunc:parse.ParseTag,
	})*/

	//一下是单通道模式
	/*engine.Run(engine.Request_q{
		Url: "https://book.douban.com",
		//Url:"http://www.zhenai.com/zhenghun",
		ParseFunc: parse.ParseTag,
		//ParseFunc: zhengai.ParseCityList,
	})*/


	/*engine.Run(engine.Request_q{
		Url:       "https://book.douban.com",
		ParseFunc: parse.ParseTag,
	})*/
}
