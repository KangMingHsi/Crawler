package client

import (
	"net/rpc"
	"spider/engine"
	"spider_distributed/config"
	"spider_distributed/worker"
)

// CreateProcessor ;
func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor, error) {

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		client := <-clientChan

		var sResult worker.ParseResult
		err := client.Call(config.ParseArticleList, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil
	}, nil
}
