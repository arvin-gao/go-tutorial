package mypkg

type myStrMap struct {
	m map[string]int
}

var myMap = myStrMap{m: make(map[string]int)}

func SetMyMap(k string, v int) {
	myMap.m[k] = v
}

func GetMyMapValue(k string) int {
	return myMap.m[k]
}
