package main

import (
    "fmt"
    "github.com/satori/go.uuid"
)

func main() {
    u, _ := uuid.NewV4()
    fmt.Printf("Hello, %s\n", u)
}
