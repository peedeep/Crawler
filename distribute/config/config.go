package config

const (
	ItemSaverPort = 1234
	// Elastic search
	ElasticIndex = "dating_profile"

	// RPC endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Parser names
	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// Rate limiting
	Qps = 20
)
