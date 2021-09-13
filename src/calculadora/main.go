package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculadora struct {
}

func ReadEntry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func (Calculadora) Operate(entrada, operador string) int {
	cleanEntry := strings.Split(entrada, operador)

	operador1 := Parsear(cleanEntry[0])
	operador2 := Parsear(cleanEntry[1])

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println("Operacion no valida D:")
		return 0
	}
}

func main() {

	entrada := ReadEntry()
	operador := ReadEntry()

	cal := Calculadora{}
	fmt.Println(cal.Operate(entrada, operador))
}
