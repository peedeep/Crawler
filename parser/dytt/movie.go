package dytt

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

//<td style="WORD-WRAP: break-word" bgcolor="#fdfddf"><a href="([^"]+)">[^<]+</a></td>

var (
	movieIdUrlRe   = regexp.MustCompile(`https://www.dytt8.net/html/gndy/dyzz/[\d]+/([0-9]+).html`)
	movieNameRe    = regexp.MustCompile(`译[^名<]*名　([^<]+) <br />[^年]*年[^代<]*代　([\d]+) <br />[^产]*产[^地<]*地　([^<]+) <br />[^上]*上映日期　([^<]+) <br />`)
	movieScoreRe   = regexp.MustCompile(`豆瓣评分　([^<]+) <br />`)
	movieThunderRe = regexp.MustCompile(`<td style="WORD-WRAP: break-word" bgcolor="#fdfddf"><a href="([^"]+)">[^<]+</a></td>`)
)

func ParseMovie(contents []byte, url string) engine.ParseResult {
	result := engine.ParseResult{}
	nameMatch := movieNameRe.FindSubmatch(contents)
	scoreMatch := movieScoreRe.FindSubmatch(contents)
	thunderMatch := movieThunderRe.FindSubmatch(contents)

	if nameMatch != nil && len(nameMatch) > 4 {
		var scoreStr string
		var thunderStr string
		if scoreMatch != nil && len(scoreMatch) > 1 {
			scoreStr = string(scoreMatch[1])
		}
		if thunderMatch != nil && len(thunderMatch) > 1 {
			thunderStr = string(thunderMatch[1])
		}
		movie := model.Movie{
			Name:       string(nameMatch[1]),
			Years:      string(nameMatch[2]),
			Origin:     string(nameMatch[3]),
			Date:       string(nameMatch[4]),
			Score:      scoreStr,
			ThunderUrl: thunderStr,
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
