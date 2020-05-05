package worker

import (
	"crawler/distribute/config"
	parser2 "crawler/dytt/parser"
	"crawler/engine"
	"crawler/zhenai/parser"
	"errors"
	"fmt"
	"log"
)

func SerializeRequest(r engine.Request) engine.SerializedRequest {
	serialize := r.Parser.Serialize()
	return engine.SerializedRequest{
		Url: r.Url,
		Parser: engine.SerializedParser{
			Name: serialize.Name,
			Args: serialize.Args,
		},
	}
}

func SerializeResult(r engine.ParseResult) engine.SerializedParseResult {
	result := engine.SerializedParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r engine.SerializedRequest) (engine.Request, error) {
	p, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, nil
}

func DeserializeResult(r engine.SerializedParseResult) (engine.ParseResult, error) {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}
	return result, nil
}

func deserializeParser(p engine.SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseMovieList:
		return parser.NewFuncParser(parser2.ParseMovieList, config.ParseMovieList), nil
	case config.ParseMovie:
		return parser.NewFuncParser(parser2.ParseMovie, config.ParseMovie), nil
	case config.ParseCityList:
		return parser.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return parser.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		}
		return nil, fmt.Errorf("invalid arg: %v", p.Args)
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
