package parser

import (
	"io/ioutil"
	"spider/model"
	"testing"
)

func TestParseMessage(t *testing.T) {
	contents, err := ioutil.ReadFile("message_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parseMessage(contents, "https://www.ptt.cc", "", "[電競] 2019 LEC Summer W4D2")

	const resultSize = 1

	expectedMessages := []model.Message{
		model.Message{
			IsRecommended: "推",
			AccountName:   "turningright",
			Msg:           "G2 0200 ==",
			Time:          "07/13 22:29",
		},
		model.Message{
			IsRecommended: "噓",
			AccountName:   "bear15328",
			Msg:           "沒興趣  有LMS再叫我",
			Time:          "07/13 22:29",
		},
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d resquests; but had %d", len(result.Items), resultSize)
	}

	messages := result.Items[0].Payload.([]model.Message)
	for i, item := range expectedMessages {

		if item.AccountName != messages[i].AccountName &&
			item.Msg != messages[i].Msg &&
			item.Time != messages[i].Time &&
			item.IsRecommended != messages[i].IsRecommended {
			t.Errorf("expected message #%d %v; but it was %v", i, messages[i], item)
		}
	}

}
