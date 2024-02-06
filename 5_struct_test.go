package gotutorial

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type StructA struct {
	// Nested Struct
	StructB
	Field1 string
	Field2 string
}

type StructB struct {
	count int
}

func (s *StructB) Count() {
	s.count++
}

func (s StructB) Count2() {
	// ineffective assignment to field MySecondStructure.count
	s.count++
}

func TestStruct(t *testing.T) {
	// Creating an instance of a struct
	var s1 StructA
	s1.Field1 = "v"

	// Creating an instance using struct literate
	s2 := StructA{
		StructB: StructB{
			count: 1,
		},
		Field2: "v",
	}

	// Creating an instance using the new keyword
	_ = new(StructA)

	// Creating an instance using the pointer address operator
	_ = &StructA{Field1: "x"}

	// Compare struct instances
	ptr("s1 == s2", s1 == s2)
}

func TestDuplicatedField(t *testing.T) {
	type A struct{ a, c int }
	type B struct{ a, b int }

	type S1 struct {
		A
		B
		b float32
	}
	var s S1
	s.A.a = 1
	s.B.a = 1
	ptr(s) // {1, 0}, {1, 0}, 0
	// ptr(c.a) // failed: ambiguous selector c.a

	s.b = 2
	s.c = 3
	ptr(s.b, s.B.b, s.c, s.A.c) // 2, 0, 3, 3
	s.b = 4
	s.A.c = 5
	ptr(s.b, s.B.b, s.c, s.A.c) // 2, 0
}

func TestStructTag(t *testing.T) {
	type UserA struct {
		// json tag
		ID   int64  `json:"iD"`
		Name string `json:"name1"`
		// non tag
		Age int
		// 3rd tag
		Gender bool `gorm:"gender"`
	}

	var user = UserA{
		ID:   1,
		Name: "user1",
		Age:  10,
	}

	// To generate json string from an object.
	b, _ := json.MarshalIndent(&user, "", "\t")
	ptr(string(b))

	// To generate a object from json string.
	var user2 UserA
	_ = json.Unmarshal(b, &user2)

	// Get tag value by reflect package.
	field, ok := reflect.TypeOf(user).FieldByName("Name")
	if ok {
		fmt.Printf("field.tag: %s; Json value: %s\n", field.Tag, field.Tag.Get("json"))
	}

	ptr("===")
	uType := reflect.TypeOf(user)
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

func (a *A) MethodA() {
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
	mySportCar.MyA.MethodA() // not have mySportCar.MethodA() method on hint list.
}
