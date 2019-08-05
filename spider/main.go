package main

import (
	"spider/engine"
	"spider/ptt/parser"
	"spider/saver"
	"spider/scheduler"
	"spider_distributed/config"
	"time"
)

func main() {

	itemChan, err := saver.ItemSaver("ptt_lol")

	if err != nil {
		panic(err)
	}

	eng := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	time.AfterFunc(time.Minute*5, eng.Shutdown)

	eng.Run(engine.Request{
		Prefix: "https://www.ptt.cc",
		URL:    "/bbs/LoL/index.html",
		Parser: engine.NewFuncParser(parser.ParseArticleList, config.ParseArticleList),
	})

}
