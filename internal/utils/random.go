package utils

import (
	"math/rand"
)

func GetRandomNumber() int {
	return rand.Intn(100)
}
