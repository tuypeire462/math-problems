 package main

import (
    "math/rand"
    "time"
)

func GenerateRandomInt(min int, max int) int {
    rand.Seed(time.Now().UnixNano())
    return min + rand.Intn(max-min+1)
}
