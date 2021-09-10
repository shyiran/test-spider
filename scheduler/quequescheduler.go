//多通道任务调度器
package scheduler

import "test-spider/engine"

type QueueScheduler struct {
	//请求任务通道
	requestChan chan engine.Request_q
	//处理通道 二维
	workerChan  chan chan engine.Request_q
}
//处理方法
func (s *QueueScheduler) WorkChan() chan engine.Request_q {
	return make(chan engine.Request_q)
}
//提交任务至请求任务通道
func (s *QueueScheduler) Submit(r engine.Request_q) {
	s.requestChan <- r
}
//提交任务至处理任务通道
func (s *QueueScheduler) WorkReady(w chan engine.Request_q) {
	s.workerChan <- w
}

func (s *QueueScheduler) Run() {
	//
	s.workerChan = make(chan chan engine.Request_q)
	s.requestChan = make(chan engine.Request_q)
	go func() {
		var requestQ []engine.Request_q
		var workQ []chan engine.Request_q
		for {
			var activeRequest engine.Request_q
			var activework chan engine.Request_q

			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activework = workQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workQ = append(workQ, w)
			case activework <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
