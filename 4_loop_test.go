package gotutorial

import (
	"testing"
)

func TestLoop(t *testing.T) {
	i := 1
	for i <= 3 {
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
	}

	// endless loop
	for {
		if i > 10 {
			break
		}
		i++
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
	}

	for i := 0; ; i++ { // <=> for i := 0; true; i++ {
	}

}

func TestLoopArrayAndSlice(t *testing.T) {
	nums := []int{1, 2, 3}
	for index, value := range nums {
		println(index, value)
		pPtr(&nums[index])
		pPtr(&value)
	}
	for index := range nums {
		println(index)
	}
}

func TestLoopMap(t *testing.T) {
	m := map[string]string{
		"1": "v-1",
		"2": "v-2",
		"3": "v-3",
		"4": "v-4",
		"5": "v-5",
		"6": "v-6",
		"7": "v-7",
	}
	// 順序不一定
	for key, value := range m {
		println(key, value)
	}

	for key := range m {
		println(key)
	}
}

func TestLoopChannel(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "range channel"
	for x := range ch {
		println(x)
	}
}
