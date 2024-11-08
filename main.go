package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func generateNitroLinks(quantity int, linkType string) {
    baseUrl := "https://discord.gift/"
    length := 16

    if linkType == "p" {
        baseUrl = "https://discord.com/billing/promotions/"
        length = 24
    }

    rand.Seed(time.Now().UnixNano())
    for i := 1; i <= quantity; i++ {
        code := make([]byte, length)
        for j := range code {
            code[j] = letters[rand.Intn(len(letters))]
        }
        formattedCode := string(code)

        if linkType == "p" {
            formattedCode = insertDashes(formattedCode)
        }

        fmt.Printf("%d: %s%s\n", i, baseUrl, formattedCode)
    }
}

func insertDashes(code string) string {
    var parts []string
    for i := 0; i < len(code); i += 4 {
        end := i + 4
        if end > len(code) {
            end = len(code)
        }
        parts = append(parts, code[i:end])
    }
    return strings.Join(parts, "-")
}

func main() {
    fmt.Print("Promo (p) or Gift (g)?: ")
    linkType := ""
    fmt.Scan(&linkType)

    if linkType != "p" && linkType != "g" {
        fmt.Println("有効な値を入力してください")
        return
    }

    fmt.Print("Count: ")
    quantity := 0
    if _, err := fmt.Scan(&quantity); err != nil || quantity <= 0 {
        fmt.Println("正の整数を入力してください")
        return
    }

    generateNitroLinks(quantity, linkType)
}
