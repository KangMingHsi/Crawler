package parser

import (
	"log"
	"regexp"
	"spider/engine"
	"spider_distributed/config"
)

var articleListRe = regexp.MustCompile(`<a href="(/bbs/LoL/[0-9a-zA-Z.%?=]+html+)"[^>]*>([^<]+)</a>`)
var previousPageRe = regexp.MustCompile(`<a class="btn wide" href="(/bbs/LoL/[0-9a-zA-Z.]+html+)"[^>]*>[^\s]* 上頁</a>`)

// ParseArticleList : return ArticleList
func ParseArticleList(contents []byte, prefix string) engine.ParseResult {
	log.Println("------ParseArticleList---------")
	matches := articleListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		articleName := string(m[2])
		// result.Items = append(result.Items, engine.Item{
		// 	Payload: articleName,
		// })
		result.Requests = append(result.Requests,
			engine.Request{
				Prefix: prefix,
				URL:    string(m[1]),
				Parser: NewMessageParser(string(m[1]), articleName),
			})
	}

	matchURL := previousPageRe.FindSubmatch(contents)
	result.Requests = append(result.Requests, engine.Request{
		Prefix: prefix,
		URL:    string(matchURL[1]),
		Parser: engine.NewFuncParser(ParseArticleList, config.ParseArticleList),
	})

	return result
}
