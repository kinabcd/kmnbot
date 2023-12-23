package kmnbot_test

import (
	"embed"
	"strings"
	"testing"

	"github.com/kinabcd/kmnbot"
)

//go:embed fakebox
var fakeboxFs embed.FS

func fakeBox() kmnbot.Box {
	file, _ := fakeboxFs.Open("fakebox/1.txt")
	defer file.Close()
	box, _ := kmnbot.ParseBox(file)
	return box
}
func TestBoxSize(t *testing.T) {
	if kmnbot.Box(nil).Size() != 0 {
		t.Errorf("nil box should be 0")
	}
	f1 := fakeBox()
	assert(t, 275, f1.Size())
	assert(t, 8, f1.Items().Size())
	assert(t, 267, f1.Kmns().Size())
}

func TestSort(t *testing.T) {
	f1 := fakeBox()
	assert(t, "星圖炸彈", f1.Items().Sort(func(kmn1, kmn2 kmnbot.Kmn) bool { return kmn1.Star > kmn2.Star }).First().Name)
	f1_1 := f1.Kmns().FilterNotFullLevel()
	f1_1 = f1_1.Filter(func(kmn kmnbot.Kmn) bool { return kmn.Data().Exp > 1 })
	f1_1 = f1_1.Sort(func(kmn1, kmn2 kmnbot.Kmn) bool { return kmn1.Data().Exp < kmn2.Data().Exp })
	assert(t, "機器冥狼神,逐星者‧占卜師機器狼,藥物論‧藥師機器狼,黑暗紳士‧伯爵機器狼", strings.Join(f1_1.Names(), ","))

}
func TestFilter(t *testing.T) {
	f1 := fakeBox()
	f1_1 := f1.FilterName([]string{"機器冥狼神", "逐星者‧占卜師機器狼", "藥物論‧藥師機器狼", "黑暗紳士‧伯爵機器狼"})
	assert(t, 4, f1_1.Size())
	f1_1_1 := f1_1.FilterName([]string{"NotExisted", "逐星者‧占卜師機器狼", "機器冥狼神", "星圖炸彈"})
	assert(t, 2, f1_1_1.Size())
	f1_2 := f1.FilterName([]string{""})
	assert(t, 0, f1_2.Size())
	f1_3 := f1.FilterName([]string{"NotExisted", "逐星者‧占卜師機器狼", "機器冥狼神", "星圖炸彈"})
	assert(t, 3, f1_3.Size())
	f1_4 := f1.Filter(func(kmn kmnbot.Kmn) bool { return !kmn.FullLevel && kmn.Data().Exp == 1 })
	assert(t, 3, f1_4.Size())

}

func assert[T comparable](t *testing.T, expect, actual T) {
	if expect != actual {
		t.Errorf("expect is %v, actual is %v", expect, actual)
	}
}
