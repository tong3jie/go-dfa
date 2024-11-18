package dfa

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

}

func BenchmarkDirty(b *testing.B) {
	Dirty := New()
	words := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	Dirty.AddWords(words)
	src := make([]string, 0, 30)
	for i := 0; i < b.N; i++ {
		Dirty.Check("明天是你的生日，a，明天是你的生日", nil)
		src = src[:0]
	}

}

func TestCheckAndReplace(t *testing.T) {
	Dirty := New()
	words := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "生日"}
	Dirty.AddWords(words)
	src := make([]string, 0, 30)
	abc, ok := Dirty.CheckAndReplace("明天是你的生日，a，明天是你的生日，a，明天是你的生日，a，明天是你的生日，a，明天是你的生日，a，明天是你的生日，a，明天是你的生日，a，明天是你的生日，a，明天是你的生日，a", &src)
	fmt.Println(abc, ok)
	fmt.Println(src)
}
