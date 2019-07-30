package worker

import "crawler/engine"

type CrawlService struct {
}

func (CrawlService) Process(req engine.SerializedRequest, result *engine.SerializedParseResult) error {
	request, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(request)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
