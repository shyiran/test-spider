//单通道任务调度器
package scheduler

import "test-spider/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request_q
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request_q)
}

func (s *SimpleScheduler) WorkReady(chan engine.Request_q) {
	return
}

func (s *SimpleScheduler) WorkChan() chan engine.Request_q {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(r engine.Request_q) {
	go func() { s.workerChan <- r }()
}
