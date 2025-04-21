package main

type Blockchain struct {
    Blocks []Block
}

func NewBlockchain() *Blockchain {
    genesis := CreateBlock(0, "Bloco Gênese da VAIRA", "")
    return &Blockchain{[]Block{genesis}}
}

func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := CreateBlock(prevBlock.Index+1, data, prevBlock.Hash)
    bc.Blocks = append(bc.Blocks, newBlock)
}
