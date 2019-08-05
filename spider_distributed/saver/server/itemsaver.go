package main

import (
	"fmt"
	"log"
	"spider_distributed/config"
	"spider_distributed/rpcsupport"
	"spider_distributed/saver"

	"gopkg.in/olivere/elastic.v6"
)

func main() {
	log.Fatal(serveRPC(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRPC(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	rpcsupport.ServeRPC(host,
		&saver.ItemSaverService{
			Client: client,
			Index:  index,
		})

	return nil
}
