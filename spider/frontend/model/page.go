package model

// SearchResult ;
type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PreFrom  int
	NextFrom int
	Items    []interface{}
}
