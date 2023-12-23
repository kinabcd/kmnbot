package kmnbot

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Kmn struct {
	Name      string
	Star      int64
	Level     int64
	FullLevel bool
	Rank      int64
	FullRank  bool
	Type      string
	Evaluate  int64
}

func (k Kmn) Data() HandbookItem {
	if data, ok := Handbook[k.Name]; ok {
		return data
	}
	return HandbookItem{}
}

func FetchBox(id string) (kmns Box, err error) {
	res, err := http.Get("https://www.kmnbot.com/pets/" + id + ".txt")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("statusCode %d", res.StatusCode)
	}
	return ParseBox(res.Body)
}

func ParseBox(reader io.Reader) (kmns Box, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "★") {
			continue
		}
		spart := filterNotEmpty(strings.Split(line, " "))
		starStr := strings.TrimPrefix(spart[0], "★")
		star, err := strconv.ParseInt(starStr, 10, 64)
		if err != nil {
			return nil, err
		}
		typeStr := strings.TrimSuffix(strings.TrimPrefix(spart[1], "["), "]")
		suffixPart := strings.Split(spart[3], "／")
		var level int64 = 0
		var rank int64 = 0
		var fullLevel = false
		var fullRank = false
		var evaluate = -1
		for _, s := range suffixPart {
			switch {
			case strings.HasPrefix(s, "等級"):
				if strings.Contains(s, "（滿）") {
					fullLevel = true
				}
				levelStr := strings.TrimSuffix(strings.TrimPrefix(s, "等級"), "（滿）")
				level, err = strconv.ParseInt(levelStr, 10, 64)
				if err != nil {
					return nil, err
				}
			case strings.HasPrefix(s, "階級"):
				if strings.Contains(s, "（滿）") {
					fullRank = true
				}
				rankStr := strings.TrimSuffix(strings.TrimPrefix(s, "階級"), "（滿）")
				rank, err = strconv.ParseInt(rankStr, 10, 64)
				if err != nil {
					return nil, err
				}
			case s == "普通":
				evaluate = 0
			case s == "很不錯":
				evaluate = 1
			case s == "太完美了！":
				evaluate = 2
			}
		}

		kmns = append(kmns, Kmn{
			Name:      spart[2],
			Star:      star,
			Level:     level,
			FullLevel: fullLevel,
			Rank:      rank,
			FullRank:  fullRank,
			Type:      typeStr,
			Evaluate:  int64(evaluate),
		})

	}

	return
}
