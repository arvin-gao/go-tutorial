package gotutorial

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestStructTag(t *testing.T) {
	type _User struct {
		ID     int64  `json:"iD"`
		Name   string `json:"name1"`
		Age    int
		Gender bool `gorm:"gender"`
	}

	var u = _User{
		ID:   1,
		Name: "user1",
		Age:  10,
	}

	b, _ := json.MarshalIndent(&u, "", "\t")
	fmt.Println(string(b))

	// 取得 tag value
	field, ok := reflect.TypeOf(u).FieldByName("Name")
	if ok {
		fmt.Printf("field.tag: %s; Json value: %s\n", field.Tag, field.Tag.Get("json"))
	}

	fmt.Println("===")
	uType := reflect.TypeOf(u)
	for i := 0; i < uType.NumField(); i++ {
		field := uType.Field(i)
		if field.Tag == "" {
			fmt.Printf("non tag, value: %v\n", field.Name)
			continue
		}

		if v := string(field.Tag.Get("json")); v != "" {
			println("has json tag, value:")
			pfTree("tag: %s", field.Tag)
			pfTree("value: %s", v)
		}

		if v := string(field.Tag.Get("gorm")); v != "" {
			println("has gorm tag, value:")
			pfTree("tag: %s", field.Tag)
			pfTree("value: %s", v)
		}

	}
}

func TestStructSize(t *testing.T) {
	type Empty struct{}
	type V1 struct {
		v1 string
	}

	type A struct {
		v1 string
		v2 int32
	}

	type B struct {
		v1 string
		v2 int64
	}

	type C struct {
		v1 string
		v2 string
		A  A
	}

	fmt.Println(unsafe.Sizeof(Empty{}))
	fmt.Println(unsafe.Sizeof(V1{}))
	fmt.Println(unsafe.Sizeof(A{}))
	fmt.Println(unsafe.Sizeof(B{}))
	fmt.Println(unsafe.Sizeof(C{}))
}

func TestDuplicatedField(t *testing.T) {
	type A struct{ a int }
	type B struct{ a, b int }

	type C struct {
		A
		B
		b float32
	}
	var c C
	c.A.a = 1
	c.B.a = 1
	fmt.Println(c)
	// fmt.Println(c.a) // error!

	c.b = 2
	fmt.Println(c.b)
	c.b = 1.1
	fmt.Println(c.b)
}
