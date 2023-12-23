package kmnbot

import "regexp"

type Response struct {
	IsHere     bool
	Rin        string
	Rainbow    string
	NextSeason string
	Search     []string
	Other      []string
}

var rinRegexp = []*regexp.Regexp{
	regexp.MustCompile(`→ 亂入率：(.+)%`),
	regexp.MustCompile(`亂入率變為 (.+)% 了`),
}
var rainbowRegexp = []*regexp.Regexp{
	regexp.MustCompile(`→ 彩虹齒輪：(.+)個`),
	regexp.MustCompile(`彩虹齒輪.*（→ (.+)個）`),
}
var seasonRegexp = []*regexp.Regexp{
	regexp.MustCompile(`季節：今.+／明(.+)`),
	regexp.MustCompile(`→ 季節：明(.+)`),
}
var searchRegexp = regexp.MustCompile(`\*\*探索結果\*\*：(.+)（`)
var otherRegexp = []*regexp.Regexp{
	regexp.MustCompile(`\*\*活動贈禮：\*\*(\S+)`),
	regexp.MustCompile(`\*\*工作所得\*\*：(.+)（`),
	regexp.MustCompile(`→ 獲得：(\S+)`),
	regexp.MustCompile(`獲得\*\*(\S+)\*\*`),
}

func (r *Response) Add(text string) {
	r.IsHere = true
	for _, rgx := range seasonRegexp {
		if strs := rgx.FindStringSubmatch(text); len(strs) >= 2 {
			r.NextSeason = strs[1]
		}
	}
	for _, rgx := range rainbowRegexp {
		if strs := rgx.FindStringSubmatch(text); len(strs) >= 2 {
			r.Rainbow = strs[1]
		}
	}
	for _, rgx := range rinRegexp {
		if strs := rgx.FindStringSubmatch(text); len(strs) >= 2 {
			r.Rin = strs[1]
		}
	}
	for _, rgx := range otherRegexp {
		if strs := rgx.FindStringSubmatch(text); len(strs) >= 2 {
			r.Other = append(r.Other, strs[1])
		}
	}

	search := searchRegexp.FindAllStringSubmatch(text, -1)
	for _, v := range search {
		r.Search = append(r.Search, v[1])
	}

}
