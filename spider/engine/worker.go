package engine

import (
	"log"
	"spider/fetcher"
)

// Worker ;
func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching: %s", r.Prefix+r.URL)
	body, err := fetcher.Fetch(r.Prefix + r.URL)
	if err != nil {
		log.Printf("Fetch error with url %s: %v", r.URL, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Prefix), nil
}
