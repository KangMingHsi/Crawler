package main

import (
	"spider/engine"
	"spider/model"
	"spider_distributed/config"
	"spider_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	go serveRPC(host, "test1")
	time.Sleep(time.Second * 3)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		URL:  "",
		Type: "message",
		ID:   "8787",
		Payload: []model.Message{
			model.Message{
				IsRecommended: "推",
				AccountName:   "turningright",
				Msg:           "G2 0200 ==",
				Time:          "07/13 22:29",
			},
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRPC,
		item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
