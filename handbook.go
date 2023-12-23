package kmnbot

import (
	"bufio"
	"embed"
	"strconv"
	"strings"
)

type HandbookItem struct {
	Level  int64
	Exp    int64
	Type   string
	IsItem bool
}

//go:embed handbook.txt
var handbookFs embed.FS
var Handbook map[string]HandbookItem = map[string]HandbookItem{}
var ItemNames = []string{"電子方塊", "絕冬雪網", "節氣琥珀", "龍脈逆鱗", "脈衝磁盤", "國戰契約", "資料結晶", "星圖炸彈", "重生命盤", "靈樹秘果"}

func init() {
	file, _ := handbookFs.Open("handbook.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		splited := strings.Split(line, "\t")
		if len(splited) != 6 {
			panic(line)
		}
		expStr := splited[4]
		exp, err := strconv.ParseInt(expStr, 10, 64)
		if err != nil {
			exp = -1
		}
		level, _ := strconv.ParseInt(splited[3], 10, 64)
		data := HandbookItem{
			Exp:    exp,
			Level:  level,
			Type:   splited[1],
			IsItem: splited[5] == "道具" || splited[5] == "召喚物",
		}
		Handbook[splited[0]] = data
	}
}
