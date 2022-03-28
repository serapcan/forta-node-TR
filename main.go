package main

import (
	"log"

	"github.com/forta-protocol/forta-node/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
