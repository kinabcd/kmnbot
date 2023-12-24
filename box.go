package kmnbot

import (
	"slices"
	"sort"
)

type Box []Kmn

func (box Box) Get(name string) *Kmn {
	return box.Filter(func(kmn Kmn) bool { return kmn.Name == name }).First()
}

func (box Box) First() *Kmn {
	if box.Size() > 0 {
		return &box[0]
	} else {
		return nil
	}
}

func (box Box) Filter(condi func(kmn Kmn) bool) Box {
	r := []Kmn{}
	for _, kmn := range box {
		if condi(kmn) {
			r = append(r, kmn)
		}
	}
	return r

}
func (box Box) FilterName(names []string) Box {
	w := make(map[string]bool)
	for _, wanted := range names {
		w[wanted] = true
	}
	return box.Filter(func(kmn Kmn) bool { return w[kmn.Name] })
}
func (box Box) FilterNotName(names []string) Box {
	w := make(map[string]bool)
	for _, droped := range names {
		w[droped] = true
	}
	return box.Filter(func(kmn Kmn) bool { return !w[kmn.Name] })
}
func (box Box) FilterNotFullRank() Box {
	return box.Filter(func(kmn Kmn) bool { return !kmn.FullRank })
}
func (box Box) FilterNotFullLevel() Box {
	return box.Filter(func(kmn Kmn) bool { return !kmn.FullLevel })
}
func (box Box) Kmns() Box {
	return box.FilterNotName(ItemNames)
}
func (box Box) Items() Box {
	return box.FilterName(ItemNames)
}
func (box Box) ItemRank() int64 {
	if items := box.Items(); len(items) > 0 {
		return items[0].Rank
	} else {
		return 0
	}
}
func (box Box) Size() int {
	return len(box)
}
func (box Box) Names() []string {
	res := make([]string, 0, box.Size())
	for _, v := range box {
		res = append(res, v.Name)
	}
	return res
}

// 邪樹使‧絕惡魔狼機
// 技能②：混沌儀式
// → 搭檔為「魔狼」時才能發動。
// 移除BOX中的「電子方塊」「資料結晶」「靈樹秘果」、獲得1隻Ni型態的寵物。
func (box Box) CanNi() bool {
	k1, k2, k3 := box.Get("電子方塊"), box.Get("資料結晶"), box.Get("靈樹秘果")
	return k1 != nil && k2 != nil && k3 != nil && k1.Level < 10 && k2.Level == 0 && k3.Level == 0
}
func (box Box) CanNiAEbox() bool {
	k1, k2, k3 := box.Get("電子方塊"), box.Get("資料結晶"), box.Get("靈樹秘果")
	return k2 != nil && k3 != nil && k1 == nil && k2.Level == 0 && k3.Level == 0
}
func (box Box) CanEboxChange() bool {
	k1 := box.Get("電子方塊")
	return k1 != nil && k1.FullLevel
}
func (box Box) Level() int64 {
	sum := int64(0)
	for _, kmn := range box {
		sum += kmn.Level
	}
	return sum
}
func (box Box) Sort(less func(kmn1, kmn2 Kmn) bool) Box {
	newBox := slices.Clone(box)
	sort.Slice(newBox, func(i, j int) bool {
		return less(newBox[i], newBox[j])
	})
	return newBox
}
