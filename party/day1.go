package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "unicode"
)

// Função para encontrar o valor de calibração em uma string.
func valorCalibracao(linha string) (int, error) {
    var primeiroDigito, ultimoDigito rune
    primeiroEncontrado := false

    for _, char := range linha {
        if unicode.IsDigit(char) {
            if !primeiroEncontrado {
                primeiroDigito = char
                primeiroEncontrado = true
            }
            ultimoDigito = char
        }
    }

    if !primeiroEncontrado {
        return 0, fmt.Errorf("nenhum dígito encontrado")
    }

    valor, err := strconv.Atoi(string(primeiroDigito) + string(ultimoDigito))
    if err != nil {
        return 0, err
    }

    return valor, nil
}

func main() {
    arquivo, err := os.Open("lista.txt")
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }
    defer arquivo.Close()

    somaTotal := 0
    scanner := bufio.NewScanner(arquivo)
    for scanner.Scan() {
        linha := scanner.Text()
        valor, err := valorCalibracao(linha)
        if err != nil {
            fmt.Printf("Erro na linha '%s': %v\n", linha, err)
        } else {
            somaTotal += valor
            fmt.Printf("Valor de calibração na linha '%s': %d\n", linha, valor)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Erro ao ler o arquivo:", err)
    }

    fmt.Println("Soma total dos valores de calibração:", somaTotal)
}
