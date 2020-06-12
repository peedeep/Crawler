package zhenai

import (
	"crawler/engine"
)

type ProfileParser struct {
	username string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(url, contents, p.username)
}

func (p *ProfileParser) Serialize() engine.SerializedParser {
	return engine.SerializedParser{
		Name: "ParseProfile", Args: p.username,
	}
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		name,
	}
}