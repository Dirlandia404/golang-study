package main

import (
    "fmt"
)

func calcula(op string, num1 float64, num2 float64) float64 {
    switch op {
    case "soma":
        return num1 + num2
    case "subtração":
        return num1 - num2
    case "multiplicação":
        return num1 * num2
    case "divisão":
        if num2 == 0 {
            return 0
        }
        return num1 / num2
    default:
        return 0
    }
}

func main() {
    fmt.Println("Bem-vindo(a) à calculadora!")

    for {
        fmt.Println("Digite a operação desejada (soma, subtração, multiplicação ou divisão): ")
        op := ""
        fmt.Scanln(&op)

        fmt.Println("Digite o primeiro número: ")
        num1 := 0.0
        fmt.Scanln(&num1)

        fmt.Println("Digite o segundo número: ")
        num2 := 0.0
        fmt.Scanln(&num2)

        result := calcula(op, num1, num2)
        fmt.Println("O resultado é:", result)

        fmt.Println("Deseja continuar? (s/n): ")
        continueCalc := ""
        fmt.Scanln(&continueCalc)
        if continueCalc == "n" {
            break
        }
    }
}
