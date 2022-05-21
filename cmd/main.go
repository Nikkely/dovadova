package main

import (
	"log"

	"github.com/Nikkely/dovadova/fetcher"
)

func main() {
	if err := fetcher.Fetch(); err != nil {
		log.Fatalln(err)
	}
}
