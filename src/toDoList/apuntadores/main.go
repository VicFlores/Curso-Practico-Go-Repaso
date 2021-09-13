package main

import "fmt"

func ChangeValue(a int) {
	fmt.Println(&a)
	a = 36
}

func main() {

	// Definimos una variable
	x := 25

	fmt.Println(x)

	// & nos dice en que espacio de la memoria se encuentra nuestra variable
	// & referencia y hace una copia por eso no cambia el valor
	fmt.Println(&x)

	y := &x

	fmt.Println(y)

	// * va al espacio en memoria y toma ese valor
	fmt.Println(*y)

}
