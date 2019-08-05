package main

import (
	"fmt"
	"spider/ptt/parser"
	"spider_distributed/config"
	"spider_distributed/rpcsupport"
	"spider_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"

	go rpcsupport.ServeRPC(host, worker.CrawlService{})

	time.Sleep(time.Second * 3)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Prefix: "https://www.ptt.cc",
		// URL:    "/bbs/LoL/index.html",
		// Parser: worker.SerializedParser{
		// 	Name: config.ParseArticleList,
		// 	Args: nil,
		// },
		URL: "/bbs/LoL/M.1562223193.A.CA4.html",
		Parser: worker.SerializedParser{
			Name: config.ParseMessage,
			Args: parser.MessageArgs{
				URL:         "/bbs/LoL/M.1562223193.A.CA4.html",
				ArticleName: "[戰棋] 虛空生物增加三隻變成3/6BUFF有搞頭嗎",
			},
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRPC, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
