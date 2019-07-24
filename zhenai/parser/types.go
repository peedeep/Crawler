package parser

import (
	"crawler/engine"
)

type ParserFunc func(contents []byte, url string) engine.ParseResult

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) engine.ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() engine.SerializedParser {
	return engine.SerializedParser{
		f.name, nil,
	}
}

func NewFuncParser(parser ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: parser,
		name:   name,
	}
}

type ProfileParser struct {
	username string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(url, contents, p.username)
}

func (p *ProfileParser) Serialize() engine.SerializedParser {
	return engine.SerializedParser{
		"ParseProfile", p.username,
	}
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		name,
	}
}