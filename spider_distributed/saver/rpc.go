package saver

import (
	"log"
	"spider/engine"
	"spider/saver"

	"gopkg.in/olivere/elastic.v6"
)

// ItemSaverService ;
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// Save ;
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := saver.Save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}

	return err
}
