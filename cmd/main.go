package main

import (
	"log"

	c "github.com/meles-z/test/configs"
	"github.com/meles-z/test/internal"
)

func main() {

	configs, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Could not load config", err)
	}
	internal.NewServer(*configs)

}