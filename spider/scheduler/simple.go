package scheduler

import (
	"spider/engine"
)

// SimpleScheduler : first version
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// Submit ;
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

// WorkerChannel ;
func (s *SimpleScheduler) WorkerChannel() chan engine.Request {
	return s.workerChan
}

// WorkerReady ;
func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {

}

// Run ;
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
