kmnbot
==
kmnbot 是Plurk遊戲機器狼的golang Parser

可以
* 整理機器狼Box的內容
* 整理機器狼在Plurk中的回應，建議搭配 [golurk](https://github.com/kinabcd/goplurk)

機器狼BOX
--
從機器狼Server下載最新的BOX資料

其中plurk account是Plruk網址中最後一段，例如 https://www.plurk.com/kinabcd 中的 kinabcd
```golang
var box Box = kmnbot.FetchBox("plurk account")
```
可以取出以下內容
```golang
type Box []Kmn
type Kmn struct {
	Name string
	// 稀有度
	Star int64
	// 等級
	Level int64
	// 若 Level 滿時為 true，滿級等級可從 HandbookItem.Level 取得
	FullLevel bool
	// 階級
	Rank int64
	// 若 Rank 為 100 則為 true
	FullRank bool
	// 必為 紅, 藍, 黃, 綠, 黑, 白 其中之一
	Type string
	// -1:未知, 0:普通, 1:不錯, 2:太完美了!
	// 體質未知的機器狼需要使用「靈能者機器狼」系列的技能「靈腦檢驗」使其顯示
	Evaluate int64
}
```

機器狼回應
--
```golang
client, _ := goplurk.NewClient(consumerToken, consumerSecret, accessToken, accessSecret)
kmnResponce := &kmnbot.Response{}
if ress, err := client.Responses.GetById(plurkId, 0, 100); err == nil && len(ress.Responses) > 0 {
    for _, res := range ress.Responses {
        if res.UserId == 7863974 {
            kmnResponce.Add(res.ContentRaw)
        }
    }
}
```
可以取出以下內容
```golang
type Response struct {
	// 是否有Add過回應
	IsHere bool
	// 亂入率
	Rin string
	// 彩虹齒輪
	Rainbow string
	// 明季節
	NextSeason string
	// 探索結果
	Search []string
	// 其他所得，例如活動贈禮、工作所得等
	Other []string
}
```