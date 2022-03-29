package main

import (
	"log"

	"github.com/omekov/golang-interviews/internal/salecar"
)

func main() {
	if err := salecar.Run(); err != nil {
		log.Fatal(err)
	}
}
