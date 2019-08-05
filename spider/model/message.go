package model

import "encoding/json"

// Message : for ptt message record
type Message struct {
	IsRecommended string
	AccountName   string
	Msg           string
	Time          string
}

// FromJSONObj ;
func FromJSONObj(o interface{}) ([]Message, error) {
	message := []Message{}
	s, err := json.Marshal(o)
	if err != nil {
		return message, err
	}

	err = json.Unmarshal(s, &message)
	return message, err
}
