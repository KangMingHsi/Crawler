package worker

import (
	"errors"
	"fmt"
	"log"
	"spider/engine"
	"spider/ptt/parser"
	"spider_distributed/config"
)

// SerializedParser (Custom)
type SerializedParser struct {
	Name string
	Args interface{}
}

// Request ;
type Request struct {
	Prefix string
	URL    string
	Parser SerializedParser
}

// ParseResult ;
type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

// SerializeRequest ;
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Prefix: r.Prefix,
		URL:    r.URL,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

// SerializeResult ;
func SerializeResult(r engine.ParseResult) ParseResult {

	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}

	return result
}

// DeserializeRequest ;
func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Prefix: r.Prefix,
		URL:    r.URL,
		Parser: parser,
	}, nil
}

// DeserializeResult ;
func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request %v", err)
			continue
		}

		result.Requests = append(result.Requests, engineReq)
	}

	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseArticleList:
		return engine.NewFuncParser(parser.ParseArticleList, config.ParseArticleList), nil
	case config.ParseMessage:
		args, ok := p.Args.(map[string]interface{})
		if ok {
			return parser.NewMessageParser(args["URL"].(string), args["ArticleName"].(string)), nil
		}
		return nil, fmt.Errorf("invalid args: %v", p.Args)
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
