package main

import (
    "fmt"
)

type Item struct {
    name string
    price float64
    qty int
}

type Cart struct {
    items []Item
}

func main() {
    cart := Cart{}

    fmt.Println("Bem-vindo(a) ao carrinho de compras!")

    for {
        fmt.Println("Digite o nome do produto: ")
        name := ""
        fmt.Scanln(&name)

        if name == "" {
            break
        }

        fmt.Println("Digite o preço do produto: ")
        price := 0.0
        fmt.Scanln(&price)

        fmt.Println("Digite a quantidade do produto: ")
        qty := 0
        fmt.Scanln(&qty)

        item := Item{name, price, qty}
        cart.addItem(item)
    }

    fmt.Println("Seu carrinho de compras contém os seguintes itens:")
    for _, item := range cart.items {
        fmt.Println("* Nome:", item.name)
        fmt.Println("* Preço:", item.price)
        fmt.Println("* Quantidade:", item.qty)
    }

    fmt.Println("O valor total do seu carrinho é:", cart.total())
}

func (c Cart) addItem(item Item) {
    var qty int
    fmt.Scanln(&qty)

    if qty <= 0 {
        fmt.Println("Quantidade inválida")
        return
    }

    item.qty = qty
    c.items = append(c.items, item)
}


func (c Cart) total() float64 {
    total := 0.0
    for _, item := range c.items {
        total += item.price * float64(item.qty)
    }
    return total
}
