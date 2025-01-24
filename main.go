package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worked %d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Microsecond
		time.Sleep(dur)
	}
}

func main() {
	say(1, "Golang is the best progg!")
	say(2, "What is going on")
}
