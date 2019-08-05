package client

import (
	"log"
	"spider/engine"
	"spider_distributed/config"
	"spider_distributed/rpcsupport"
)

// ItemSaver ;
func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		for {
			item := <-out
			log.Printf("Got item: %v\n", item)

			result := ""
			err = client.Call(config.ItemSaverRPC, item, &result)

			if err != nil {
				log.Printf("Item Server: error saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}
