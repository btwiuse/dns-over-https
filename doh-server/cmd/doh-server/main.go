package main

import (
	"log"
	"os"

	"dohserver"
)

func main() {
	args := os.Args[1:]
	err := dohserver.Run(args)
	if err != nil {
		log.Fatalln(err)
	}
}
