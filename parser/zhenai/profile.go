package zhenai

import (
	"crawler/engine"
	"crawler/model"
	"log"
	"regexp"
	"strings"
)

//<div class="des f-cl" data-v-3c42fade>南昌 | 30岁 | 大专 | 未婚 | 162cm | 5001-8000元</div>
const infoRegex = `<div class="des f-cl" [^>]*>([^<]+)</div>`

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)

const basicInfoRegex = `"basicInfo":["离异","36岁","魔羯座(12.22-01.19)","160cm","54kg","工作地:阿坝金川","月收入:3千以下","自由职业","中专"]`
const detailInfoRegex = `"detailInfo":["籍贯:四川成都","不吸烟","不喝酒","和家人同住","未买车","有孩子但不在身边","是否想要孩子:以后再告诉你"]`

func ParseProfile(url string, contents []byte, name string) engine.ParseResult {
	infoRe := regexp.MustCompile(infoRegex)
	matches := infoRe.FindSubmatch(contents)
	result := engine.ParseResult{}
	if matches != nil {
		//log.Printf("parse profile info: %s, %s", name, matches[1])
		s := strings.Split(strings.Trim(string(matches[1]), ""), "|")
		profile := model.Profile{
			Name:      name,
			Hokou:     s[0],
			Age:       s[1],
			Education: s[2],
			Marriage:  s[3],
			Height:    s[4],
			Income:    s[5],
		}

		result.Items = append(result.Items, engine.Item{
			Url:     url,
			Id:      extractString([]byte(url), idUrlRe),
			Type:    "zhenai",
			Payload: profile,
		})
	}
	log.Printf("parse profile result size: %d", len(result.Items))
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
