package saver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"spider/engine"

	"gopkg.in/olivere/elastic.v6"
)

// ItemSaver ;
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		for {
			item := <-out
			fmt.Printf("Got item: %v\n", item)

			err := Save(client, index, item)

			if err != nil {
				log.Printf("Item Server: error saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

// Save ;
func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.ID != "" {
		indexService.Id(item.ID)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
