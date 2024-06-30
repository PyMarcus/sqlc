package util

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

const alphab = "abcdefghijklmnopqrstuvwxyz"

func init() {
    rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
    return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
    var strb strings.Builder
    for i := 0; i < n; i++ {
        strb.WriteByte(alphab[rand.Intn(len(alphab))])
    }
    return strb.String()
}

func RandomOwner() string {
    return RandomString(6)
}

func RandomMoney() string {
    return fmt.Sprintf("%d", RandomInt(0, 1000))
}

func RandomCurrency() string {
    currency := []string{"USD", "BRL"}
    return currency[rand.Intn(len(currency))]
}
