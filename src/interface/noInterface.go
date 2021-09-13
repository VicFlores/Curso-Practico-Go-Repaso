package main

import "fmt"

type Dog struct {
}

type Fish struct {
}

type Bird struct {
}

func (Dog) Walk() string {
	return "Soy un perro y estoy caminando"
}

func (Fish) Swim() string {
	return "Soy un pez y estoy nadando"
}
func (Bird) Fly() string {
	return "Soy un pajaro y estoy volando"
}

func moveDog(d Dog) {
	fmt.Println(d.Walk())
}

func moveFish(f Fish) {
	fmt.Println(f.Swim())
}

func moveBird(b Bird) {
	fmt.Println(b.Fly())
}

func main() {

	dog := Dog{}
	moveDog(dog)

	fish := Fish{}
	moveFish(fish)

	bird := Bird{}
	moveBird(bird)
}
