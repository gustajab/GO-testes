package main

import "fmt"

//var nome string = "gustavo"
func main() {
	var num1, num2 int

    fmt.Print("Digite o primeiro número: ")
    fmt.Scanln(&num1)

    fmt.Print("Digite o segundo número: ")
    fmt.Scanln(&num2)

    resultado := num1 + num2

    fmt.Printf("A soma de %d e %d é igual a %d\n", num1, num2, resultado)

	//len("ALO MUNDO")
	//fmt.Println("ola meu nome é", nome)

}
