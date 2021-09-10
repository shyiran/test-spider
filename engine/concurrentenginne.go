package engine

import (
	"fmt"
	"log"
	"test-spider/fetcher"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(q Request_q)
	Run()
	WorkReady(chan Request_q)
	WorkChan() chan Request_q
}

func (e *ConcurrentEngine) Run(seeds ...Request_q) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkCount; i++ {
		CreateWork(e.Scheduler.WorkChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemcount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item:%d,%v", itemcount, item)
			itemcount++
		}
		for _, request := range result.Requesrts {
			e.Scheduler.Submit(request)
		}
	}
}
func CreateWork(in chan Request_q, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			s.WorkReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request_q) (ParseResult, error) {
	fmt.Printf("Fetch url:%s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetch Error: %s", r.Url)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
