package main

import (
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	if min == 0 {
		return rand.Intn(max)
	}
	return rand.Intn(max-min) + min
}
