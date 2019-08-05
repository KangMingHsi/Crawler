package scheduler

import "spider/engine"

// QueueScheduler : queue request and submit to each worker
type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan (chan engine.Request)
}

// Submit ;
func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

// WorkerChannel ;
func (s *QueueScheduler) WorkerChannel() chan engine.Request {
	return make(chan engine.Request)
}

// WorkerReady ;
func (s *QueueScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

// Run ;
func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}
	}()
}
