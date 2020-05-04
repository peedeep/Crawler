package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

// ◎译　　名　助理/女助理 <br />
// ◎片　　名　The Assistant <br />
// ◎年　　代　2019 <br />
// ◎产　　地　美国 <br />
// ◎类　　别　剧情 <br />
// ◎语　　言　英语 <br />
// ◎字　　幕　中英双字幕 <br />
// ◎上映日期　2019-08-30(特柳赖德电影节) / 2020-01-31(美国) <br />
// ◎IMDb评分 6.1/10 from 1678 users <br />
// ◎豆瓣评分　7.1/10 from 111 users

var (
	movieIdUrlRe  = regexp.MustCompile(`https://www.dytt8.net/html/gndy/dyzz/[\d]+/([0-9]+).html`)
	movieNameRe   = regexp.MustCompile(`译[^名<]*名　([^<]+) <br />[^年]*年[^代<]*代　([\d]+) <br />[^产]*产[^地<]*地　([^<]+) <br />[^上]*上映日期　([^<]+) <br />`)
	movieOriginRe = regexp.MustCompile(`豆瓣评分　([^<]+) <br />`)
)

func ParseMovie(contents []byte, url string) engine.ParseResult {
	result := engine.ParseResult{}
	name := movieNameRe.FindSubmatch(contents)
	origin := movieOriginRe.FindSubmatch(contents)

	if name != nil && len(name) > 4 {
		var s string
		if origin != nil && len(origin) > 1 {
			s = string(origin[1])
		}
		movie := model.Movie{
			Name:   string(name[1]),
			Years:  string(name[2]),
			Origin: string(name[3]),
			Date:   string(name[4]),
			Score:  s,
		}
		result.Items = append(result.Items, engine.Item{
			Url:     url,
			Id:      extractString([]byte(url), movieIdUrlRe),
			Type:    "dytt",
			Payload: movie,
		})
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
