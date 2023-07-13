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
		ptr(index, value)
		pPtr(&nums[index])
		pPtr(&value)
	}
	for index := range nums {
		ptr(index)
	}
}

func TestLoopMap(t *testing.T) {
	m := map[string]string{
		"1": "v1",
		"2": "v2",
		"3": "v3",
		"4": "v4",
		"5": "v5",
		"6": "v6",
		"7": "v7",
	}
	// 順序不一定
	for key, value := range m {
		ptr(key, value)
	}

	for key := range m {
		ptr(key)
	}
}

func TestLoopChannel(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "range channel"
	for x := range ch {
		ptr(x)
	}
}
