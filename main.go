package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 1. Peça ao usuário para digitar duas strings e retorne a concatenação das duas.

	var str0 string
	var str1 string

	fmt.Print("Usuário, digite um texto, por favor: ")
	fmt.Scan(&str0)

	fmt.Print("Usuário, digite outro texto, por favor: ")
	fmt.Scan(&str1)

	fmt.Printf(str0 + str1)

	// 2. Peça ao usuário para digitar uma string e retorne o número de caracteres nessa string.

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Usuário, digite um texto, por favor: ")
	scanner.Scan()
	str := scanner.Text()

	fmt.Println(len(str), "é o número de caracteres que tem no texto.")

	// 3. Peça ao usuário para digitar uma string e um caractere e retorne o número de ocorrências desse caractere na string.

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Usuário, digite um texto, por favor: ")
	scanner.Scan()
	str := scanner.Text()

	fmt.Print("Usuário, digite um caractere, por favor: ")
	scanner.Scan()
	r := scanner.Text()

	if strings.Contains(str, r) {

		fmt.Printf("O texto contém o caractere %s.", r)

	} else {

		fmt.Printf("O texto não contém o caractere %s.", r)

	}

	// 4. Peça ao usuário para digitar uma string e retorne a mesma string com as letras em maiúsculo.

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Usuário, digite um texto, por favor: ")
	scanner.Scan()
	str := scanner.Text()

	fmt.Println(strings.ToUpper(str))

	// 5. Peça ao usuário para digitar uma string e um número n e retorne a mesma string com as n primeiras letras em maiúsculo.

	var resultado string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Usuário, digite um texto, por favor: ")
	scanner.Scan()
	str := scanner.Text()

	var x int
	fmt.Print("Usuário, digite o número até quanto você quer deixar as letras em maíusculo, por favor: ")
	fmt.Scan(&x)

	for i := 0; i <= x; i++ {

		resultado = resultado + strings.ToUpper(string(str[i]))

	}

	for i := x + 1; i < len(str); i++ {

		resultado = resultado + string(str[i])

	}

	fmt.Println(resultado)

}
