package main

import (
    "fmt"
)

func main() {
    fmt.Println("VAIRA blockchain inicializada.")

    // Criar blockchain e adicionar bloco de exemplo
    blockchain := NewBlockchain()
    blockchain.AddBlock("Primeira transação ética da VAIRA")
    for _, block := range blockchain.Blocks {
        fmt.Printf("Bloco #%d - Dados: %s\n", block.Index, block.Data)
    }
}
