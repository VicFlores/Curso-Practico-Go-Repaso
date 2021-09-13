package main

import (
	"fmt"
	"net/http"
	"time"
)

func CheckServer(servers string) {
	_, err := http.Get(servers)

	if err != nil {
		fmt.Println(servers, "no disponible :(")
	} else {
		fmt.Println(servers, "funcionando correctamente")
	}
}

func main() {

	start := time.Now()

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		CheckServer(server)
	}

	finished := time.Since(start)

	fmt.Printf("Tiempo de ejecucion %s\n", finished)

}
