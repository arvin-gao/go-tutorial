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
}

func TestLoopArrayAndSlice(t *testing.T) {
	nums := []int{1, 2, 3}
	for index, value := range nums {
		println(index, value)
		printPtr(&nums[index])
		printPtr(&value)
	}
	for index := range nums {
		println(index)
	}
}

func TestLoopMap(t *testing.T) {
	m := map[string]string{
		"1": "value-1",
		"2": "value-2",
		"3": "value-3",
		"4": "value-4",
		"5": "value-5",
		"6": "value-6",
		"7": "value-7",
	}
	// 順序不一定
	for key, value := range m {
		println(key, value)
	}
}

func TestLoopChannel(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "range channel"
	for x := range ch {
		println(x)
	}
}
