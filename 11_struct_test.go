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
	ptr(s2.Field1, s2.Field2, s2.Field3, s2.count)

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
	ptr(c)
	// println(c.a) // error!

	c.b = 2
	ptr(c.b)
	c.b = 1.1
	ptr(c.b)
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
	ptr(string(b))

	// 取得 tag value
	field, ok := reflect.TypeOf(u).FieldByName("Name")
	if ok {
		fmt.Printf("field.tag: %s; Json value: %s\n", field.Tag, field.Tag.Get("json"))
	}

	ptr("===")
	uType := reflect.TypeOf(u)
	for i := 0; i < uType.NumField(); i++ {
		field := uType.Field(i)
		if field.Tag == "" {
			fmt.Printf("non tag, value: %v\n", field.Name)
			continue
		}

		if v := string(field.Tag.Get("json")); v != "" {
			ptr("has json tag, value:")
			ptrfTree("tag: %s", field.Tag)
			ptrfTree("value: %s", v)
		}

		if v := string(field.Tag.Get("gorm")); v != "" {
			ptr("has gorm tag, value:")
			ptrfTree("tag: %s", field.Tag)
			ptrfTree("value: %s", v)
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

	ptr("Empty size:", unsafe.Sizeof(Empty{}))
	ptr("A size:", unsafe.Sizeof(A{}))
	ptr("B size:", unsafe.Sizeof(B{}))
	ptr("C size:", unsafe.Sizeof(C{}))
	ptr("D size:", unsafe.Sizeof(D{}))
	ptr("E size:", unsafe.Sizeof(E{}))
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

	ptr(mySportCar.DoorCount)
	ptr(mySportCar.Car.DoorCount)
	ptr(mySportCar.Car.DoorCount)
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

// TODO: tidy.
/*
1. Field C.A1.g and method C.B.g collide, so they are both not promoted.
2. Method C.B.f gets promoted as C.f.
3. Method C.m overrides C.A1.m.
*/

type A1 struct {
	g int
}

func (A1) m() int {
	return 1
}

type B int

func (B) g() {}

func (B) f() {}

type C struct {
	A1
	B
}

func (C) m() int {
	return 9
}

func TestEmbed(t *testing.T) {
	var c interface{} = C{}
	//  Method C.B.f gets promoted as C.f.
	_, bf := c.(interface{ f() })
	// Field C.A1.g and method C.B.g collide, so they are both not promoted.
	_, bg := c.(interface{ g() })
	// Method C.m overrides C.A1.m.
	i := c.(interface{ m() int })
	ptr(bf, bg, i.m())
}
