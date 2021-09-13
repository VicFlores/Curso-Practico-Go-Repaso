package main

import "fmt"

func main() {

	// break => rompe el ciclo por completo
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// continue => rompe el ciclo implicitamente donde le digamos
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
