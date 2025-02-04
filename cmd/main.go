package main

import (
	"log"
	"pachca-bot/cmd/apiserver"
)

func main() {
	config := apiserver.NewConfig()

	// consumer
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
