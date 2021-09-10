//单个
package engine

import (
	"log"
	"test-spider/fetcher"
)

type simleEngine struct{}

func (s simleEngine) Run(seeds ...Request_q) {
	var requests []Request_q
	for _, e := range seeds {
		requests = append(requests, e)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching url:%s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch Error: %s", r.Url)
		}
		parseresult := r.ParseFunc(body)
		requests = append(requests, parseresult.Requesrts...)
		for _, item := range parseresult.Items {
			log.Printf("Got item:%s", item)
		}
	}
}
