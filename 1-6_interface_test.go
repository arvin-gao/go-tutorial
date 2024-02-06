package gotutorial

import (
	"testing"
)

/*
1. hiding implement detail
*/

type Storage interface {
	Upload(file *File) error
}

type File struct {
	Content string
}

type AwsStorage struct {
}

func (a *AwsStorage) Upload(file *File) error {
	// upload to aws storage.
	return nil
}

type LocalStorage struct {
}

func (l *LocalStorage) Upload(file *File) error {
	// upload to aws storage.
	return nil
}

func TestInterface(t *testing.T) {
	uploadTextFile := func(s Storage, content string) {
		// do something...
		file := &File{Content: content}
		s.Upload(file)
	}

	awsStorage := &AwsStorage{}
	localStorage := &LocalStorage{}

	content := "hello world"

	uploadTextFile(awsStorage, content)
	uploadTextFile(localStorage, content)
}

type Eater interface {
	Eat()
}

type Runner interface {
	Run()
}

// 介面繼承
type Animal interface {
	Eater
	Runner
}

type Dog struct {
	Name string
}

func (d *Dog) Eat() {
	ptrlnf("dog %s is eating\n", d.Name)
}

func (d *Dog) Run() {
	ptrlnf("dog %s is running\n", d.Name)
}

type Cat struct {
	Name string
}

func (d *Cat) Eat() {
	ptrlnf("cat %s is eating\n", d.Name)
}

func (d *Cat) Run() {
	ptrlnf("cat %s is running\n", d.Name)
}

func ShowEat(animal Animal) {
	animal.Eat()
}

func ShowRun(animal Animal) {
	animal.Run()
}

func ShowEat2(eater Eater) {
	eater.Eat()
}

func ShowRun2(runner Runner) {
	runner.Run()
}

func TestInterface2(t *testing.T) {
	var animal Animal        // 介面
	dog := Dog{Name: "dog1"} // 物件
	cat := Cat{Name: "cat1"} // 物件
	animal = &dog
	animal = &cat
	pass(animal)

	ShowEat(&dog)
	ShowRun(&dog)
	ShowEat2(&dog)
	ShowRun2(&dog)
}
