package main

import "time"

type Block struct {
    Index     int
    Timestamp string
    Data      string
    PrevHash  string
    Hash      string
}

func CreateBlock(index int, data string, prevHash string) Block {
    block := Block{
        Index:     index,
        Timestamp: time.Now().String(),
        Data:      data,
        PrevHash:  prevHash,
    }
    block.Hash = GenerateHash(block)
    return block
}
