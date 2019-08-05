package parser

import (
	"regexp"
	"spider/engine"
	"spider/model"
	"spider_distributed/config"
)

const messageRe = `<div class="push"><span class=[^>]*>([^<^\s]+)[^<]*</span><span class=[^>]*>([^<^\s]+)[^<]*</span><span class=[^>]*>: ([^<]+)</span><span class=[^>]*>([^<^\n]+)[^<]*</span></div>`

// MessageArgs ;
type MessageArgs struct {
	URL         string
	ArticleName string
}

// MessageParser ;
type MessageParser struct {
	Args MessageArgs
}

// Parse ;
func (m *MessageParser) Parse(contents []byte, prefix string) engine.ParseResult {
	return parseMessage(contents, prefix, m.Args.URL, m.Args.ArticleName)
}

// Serialize ;
func (m *MessageParser) Serialize() (name string, args interface{}) {
	return config.ParseMessage, m.Args
}

// NewMessageParser ;
func NewMessageParser(url string, articleName string) *MessageParser {
	return &MessageParser{
		Args: MessageArgs{
			URL:         url,
			ArticleName: articleName,
		},
	}
}

func parseMessage(contents []byte, prefix string, url string, articleName string) engine.ParseResult {

	re := regexp.MustCompile(messageRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	if matches != nil {
		item := engine.Item{
			URL:     prefix + url,
			Type:    "message",
			ID:      articleName,
			Payload: []model.Message{},
		}

		payload := []model.Message{}

		for _, match := range matches {
			payload = append(payload, model.Message{
				IsRecommended: string(match[1]),
				AccountName:   string(match[2]),
				Msg:           string(match[3]),
				Time:          string(match[4]),
			})

		}
		item.Payload = payload
		result.Items = append(result.Items, item)
	}

	return result
}
