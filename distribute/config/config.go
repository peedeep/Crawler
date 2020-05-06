package config

const (
	// Service ports
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	// Elastic search
	ElasticIndex = "dating_movie"

	// RPC endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Parser names
	ParseMovie = "ParseMovie"
	ParseMovieList = "ParseMovieList"
	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// Rate limiting
	Qps = 20
)
