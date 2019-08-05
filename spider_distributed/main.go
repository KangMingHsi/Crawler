package main

import (
	"fmt"
	"log"
	"net/rpc"
	"spider/engine"
	"spider/ptt/parser"
	"spider/scheduler"
	"spider_distributed/config"
	"spider_distributed/rpcsupport"
	itemsaver "spider_distributed/saver/client"
	worker "spider_distributed/worker/client"
	"time"
)

func main() {

	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}
	pool := createClientPool(
		[]string{
			fmt.Sprintf(":%d", config.WorkerPort0),
			fmt.Sprintf(":%d", config.WorkerPort1),
		})

	processor, err := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	eng := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	time.AfterFunc(time.Minute*5, eng.Shutdown)

	eng.Run(engine.Request{
		Prefix: "https://www.ptt.cc",
		URL:    "/bbs/LoL/index.html",
		Parser: engine.NewFuncParser(parser.ParseArticleList, config.ParseArticleList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}
