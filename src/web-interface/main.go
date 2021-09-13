package main

import (
	"fmt"
	"io"
	"net/http"
)

type Escritor struct{}

func (Escritor) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {

	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println(err)
	}

	e := Escritor{}

	io.Copy(e, res.Body)

}
