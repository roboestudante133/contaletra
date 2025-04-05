package main

import (
	"bufio"
	"fmt"
	"os"
)

func input(daf string) string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(daf)
	text, _ := reader.ReadString('\n')
	return text

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Insira a letra que deseja contar")
	fmt.Print("\n")
	wletter, _, _ := reader.ReadRune()

	text := input("Insira o conteúdo")

	var num int

	for _, letter := range text {

		if letter == wletter {

			num = num + 1

		}

	}

	fmt.Printf("existe %d letras %c no texto\n", num, wletter)

}
