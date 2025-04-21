package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

func GenerateHash(block Block) string {
    record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}
