package gotutorial

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type MyStructure struct {
	MySecondStructure
	Field1 string
	Field2 string
	Field3 string
}

type MySecondStructure struct {
	count int
}

func (m *MySecondStructure) Count() {
	m.count++
}

func (m MySecondStructure) Count2() {
	m.count++ // ineffective assignment to field MySecondStructure.count
}

func TestStruct(t *testing.T) {
	var s MyStructure
	s2 := MyStructure{
		MySecondStructure: MySecondStructure{
			count: 1,
		},
		Field1: "v1",
		Field2: "v2",
		Field3: "v3",
	}
	println(s2.Field1, s2.Field2, s2.Field3, s2.count)

	pass(s, s2)
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
	println(c)
	// println(c.a) // error!

	c.b = 2
	println(c.b)
	c.b = 1.1
	println(c.b)
}

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
	println(string(b))

	// 取得 tag value
	field, ok := reflect.TypeOf(u).FieldByName("Name")
	if ok {
		fmt.Printf("field.tag: %s; Json value: %s\n", field.Tag, field.Tag.Get("json"))
	}

	println("===")
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

/*
Refer to https://stackoverflow.com/questions/2113751/sizeof-struct-in-go

bool, int8/uint8 take 1 byte
int16, uint16 - 2 bytes
int32, uint32, float32 - 4 bytes
int64, uint64, float64, pointer - 8 bytes
string - 16 bytes (2 alignments of 8 bytes)
any slice takes 24 bytes (3 alignments of 8 bytes). So []bool, [][][]string are the same (do not forget to reread the citation I added in the beginning)
array of length n takes n * type it takes of bytes.
*/
func TestStructSize(t *testing.T) {
	type Empty struct{}

	type A struct {
		v1 string
	}

	type B struct {
		v1 int32
	}

	type C struct {
		v1 string
		v2 int32
	}

	type D struct {
		v1 string
		v2 int64
	}

	type E struct {
		v1 string
		v2 string
		A  A
	}

	println("Empty size:", unsafe.Sizeof(Empty{}))
	println("A size:", unsafe.Sizeof(A{}))
	println("B size:", unsafe.Sizeof(B{}))
	println("C size:", unsafe.Sizeof(C{}))
	println("D size:", unsafe.Sizeof(D{}))
	println("E size:", unsafe.Sizeof(E{}))
}

type Car struct {
	DoorCount int
}

func (c *Car) Move() {
	fmt.Println("car moving")
}

func (c *Car) CarMethod() {
	fmt.Println("car method")
}

type A struct {
	Value int
}

func (a *A) Amethod() {
	fmt.Println("a method")
}

type SportCar struct {
	Car
	DoorCount int
	MyA       A
}

func (c *SportCar) Move() {
	fmt.Println("sportCar moving")
}

func TestObjectFeatures(t *testing.T) {
	mySportCar := &SportCar{
		DoorCount: 2,
	}

	println(mySportCar.DoorCount)
	println(mySportCar.Car.DoorCount)
	println(mySportCar.Car.DoorCount)
	mySportCar.Move()
	mySportCar.Car.Move()
	mySportCar.CarMethod()
	mySportCar.MyA.Amethod() // not have mySportCar.Amethod() method on hint list.
}

// 封裝(Encapsulation)
/*
通過限制只有特定類別的物件可以存取這一特定類別的成員
存取權限：public, private
*/

// 繼承(extends, override)

// 多型(Polymorphism
/*
abstract, interface
*/
