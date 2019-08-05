package engine

// ConcurrentEngine : concurrent type engine
type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
	Close            bool
}

// Processor ;
type Processor func(Request) (ParseResult, error)

// Scheduler : schedule requests
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChannel() chan Request
	Run()
}

// ReadyNotifier ;
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run : run the crawler
func (eng *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParseResult)
	//eng.Scheduler.ConfiguerMasterWorkerChan(in)
	eng.Scheduler.Run()

	for i := 0; i < eng.WorkerCount; i++ {
		eng.createWorker(eng.Scheduler.WorkerChannel(), out, eng.Scheduler)
	}

	for _, r := range seeds {
		eng.Scheduler.Submit(r)
	}

	eng.Close = false

	for {
		result := <-out
		for _, item := range result.Items {
			go func(item Item) { eng.ItemChan <- item }(item)
			//fmt.Printf("Got item: %v\n", item)
		}

		for _, request := range result.Requests {
			eng.Scheduler.Submit(request)
		}

		if eng.Close {
			break
		}
	}
}

// Shutdown ;
func (eng *ConcurrentEngine) Shutdown() {
	eng.Close = true
}

func (eng *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {

	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := eng.RequestProcessor(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
