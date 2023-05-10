package gotutorial

import (
	"testing"
)

type Eater interface {
	Eat()
}

type Runner interface {
	Run()
}

type Animal interface {
	Eater
	Runner
}

type Dog struct {
	Name string
}

func (d *Dog) Eat() {
	pf("%s is eating\n", d.Name)
}

func (d *Dog) Run() {
	pf("%s is running\n", d.Name)
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
	dog := Dog{Name: "dog1"}
	ShowEat(&dog)
	ShowRun(&dog)
	ShowEat2(&dog)
	ShowRun2(&dog)
}
