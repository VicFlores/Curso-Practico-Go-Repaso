package main

import "fmt"

type Animal interface {
	move() string
}

type Dog struct{}

type Fish struct{}

type Bird struct{}

func (Dog) move() string {
	return "Soy un perro y estoy caminando"
}

func (Fish) move() string {
	return "Soy un pez y estoy nadando"
}

func (Bird) move() string {
	return "Soy un pajaro y estoy volando"
}

func moveAnimal(a Animal) {
	fmt.Println(a.move())
}

func main() {

	dog := Dog{}
	moveAnimal(dog)

	fish := Fish{}
	moveAnimal(fish)

	bird := Bird{}
	moveAnimal(bird)
}
