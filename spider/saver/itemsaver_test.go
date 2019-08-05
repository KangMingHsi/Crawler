package saver

import (
	"context"
	"encoding/json"
	"spider/engine"
	"spider/model"
	"testing"

	"gopkg.in/olivere/elastic.v6"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
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

	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	err = Save(client, "ptt_lol", expected)

	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search docker go client

	resp, err := client.Get().
		Index("ptt_lol").
		Type(expected.Type).
		Id(expected.ID).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualMessage, _ := model.FromJSONObj(actual.Payload)
	actual.Payload = actualMessage

	if expected.Payload.(([]model.Message))[0] != actual.Payload.([]model.Message)[0] {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
