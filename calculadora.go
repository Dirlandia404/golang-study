package main

import (
    "fmt"
)

func main() {
    fmt.Println("Bem-vindo(a) à calculadora!")

    for {
        fmt.Println("Digite a operação desejada (soma, subtração, multiplicação ou divisão): ")
        op := ""
        fmt.Scanln(&op)

        if op == "soma" {
            fmt.Println("Digite o primeiro número: ")
            num1 := 0
            fmt.Scanln(&num1)

            fmt.Println("Digite o segundo número: ")
            num2 := 0
            fmt.Scanln(&num2)

            result := num1 + num2
            fmt.Println("O resultado é:", result)
        } else if op == "subtração" {
            fmt.Println("Digite o primeiro número: ")
            num1 := 0
            fmt.Scanln(&num1)

            fmt.Println("Digite o segundo número: ")
            num2 := 0
            fmt.Scanln(&num2)

            result := num1 - num2
            fmt.Println("O resultado é:", result)
        } else if op == "multiplicação" {
            fmt.Println("Digite o primeiro número: ")
            num1 := 0
            fmt.Scanln(&num1)

            fmt.Println("Digite o segundo número: ")
            num2 := 0
            fmt.Scanln(&num2)

            result := num1 * num2
            fmt.Println("O resultado é:", result)
        } else if op == "divisão" {
            fmt.Println("Digite o primeiro número: ")
            num1 := 0
            fmt.Scanln(&num1)

            fmt.Println("Digite o segundo número: ")
            num2 := 0
            fmt.Scanln(&num2)

            if num2 == 0 {
                fmt.Println("Erro: divisão por zero")
            } else {
                result := num1 / num2
                fmt.Println("O resultado é:", result)
            }
        } else {
            fmt.Println("Operação inválida")
        }

        fmt.Println("Deseja continuar? (s/n): ")
        continueCalc := ""
        fmt.Scanln(&continueCalc)
        if continueCalc == "n" {
            break
        }
    }
}
