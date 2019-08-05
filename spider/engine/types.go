package engine

// ParserFunc (custom)
type ParserFunc func([]byte, string) ParseResult

// Parser (Custom)
type Parser interface {
	Parse([]byte, string) ParseResult
	Serialize() (name string, args interface{})
}

// Request (Custom)
type Request struct {
	Prefix string
	URL    string
	Parser Parser
}

// ParseResult (Custom)
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// Item (Custom)
type Item struct {
	URL     string
	Type    string
	ID      string
	Payload interface{}
}

// NilParser : do nothing
type NilParser struct{}

// Parse : nothing
func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

// Serialize : nothing
func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// FuncParser (Custom)
type FuncParser struct {
	parser ParserFunc
	name   string
}

// Parse ;
func (f *FuncParser) Parse(contents []byte, prefix string) ParseResult {
	return f.parser(contents, prefix)
}

// Serialize ;
func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// NewFuncParser ;
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
