package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	log.Printf("%v: %v\n", "Время работы программы", time.Since(start))

}
