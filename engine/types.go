package engine

import "encoding/json"

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() SerializedParser
}

type ParseResult struct {
	Requests []Request
	Items	 []Item
}

type Request struct {
	Url 	string
	Parser 	Parser
}

type Item struct {
	Url	 		string
	Type 		string
	Id 			string
	Payload 	interface{}
}

func (i *Item)FromJsonObj(o interface{}) error {
	bytes, err := json.Marshal(o)
	if err == nil {
		err = json.Unmarshal(bytes, i)
	}
	return err
}

type ParserFunc func(contents []byte, url string) ParseResult

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() SerializedParser {
	return SerializedParser{
		f.name, nil,
	}
}

func NewFuncParser(parser ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: parser,
		name:   name,
	}
}

type NilParser struct {}

func (NilParser) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() SerializedParser {
	return SerializedParser{
		"NilParser",
		nil,
	}
}

/// for serialized
type SerializedParser struct {
	Name string
	Args interface{}
}

type SerializedRequest struct {
	Url string
	Parser SerializedParser
}

type SerializedParseResult struct {
	Requests 	[]SerializedRequest
	Items 		[]Item
}