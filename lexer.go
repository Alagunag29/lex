package main

import (
	"bufio"
	"fmt"
	"os"
)

// lexema: tipo variable {int, string,...}
// token : 0 = palabra reservada; 1 : signo reservada
type lex struct {
	lexema string
	token  int
}

var (
	lenguaje = []lex{
		lex{lexema: "string", token: 0},
		lex{lexema: "int", token: 0},
		lex{lexema: "(", token: 1},
		lex{lexema: ")", token: 1},
		lex{lexema: "[", token: 1},
		lex{lexema: "]", token: 1},
	}
)

func main() {

	archivo, error := os.Open("./hola.txt")
	if error != nil {
		fmt.Println("Hubo error")
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		expression := scanner.Text()
		String(expression)
	}
}

func estadoToken(token int) string {
	switch token {
	case 0:
		return "Palabra Reservada"
	case 1:
		return "Signo Reservado"
	default:
		return "Error No Definida"
	}
}

func String(palabra string) {
	token := -1
	for _, i := range lenguaje {
		if i.lexema == palabra {
			token = i.token
			break
		}
	}
	itemToken := estadoToken(token)
	fmt.Printf("Lexema: %s\t Token: %s \n", palabra, itemToken)
}
