package main

import (
	"fmt"
	"net/http"
	"time"
)

func CheckServer(servers string, channel chan<- string) {
	_, err := http.Get(servers)

	if err != nil {
		channel <- servers + " no se encuentra disponible"
	} else {
		channel <- servers + " esta funcionando "
	}
}

func main() {

	start := time.Now()
	channel := make(chan string)

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for {
		if i > 2 {
			break
		}

		for _, server := range servers {
			go CheckServer(server, channel)
		}

		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
		i++
	}

	// Hacemos que lea la info pero solo espera una respuesta para terminar y no espera las demas
	// go routines

	// fmt.Println(<-channel)

	// aca solucionamos este problema, usando for hacemos que termine cuando todas las go routines
	// acaben

	/* 	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	} */

	finished := time.Since(start)

	fmt.Printf("Tiempo de ejecucion %s\n", finished)

}
