package main

import (
	"log"
	"time"

	"github.com/Jarover/gopractice/internal/app/config"
)

func main() {
	start := time.Now()
	log.Println(config.VersionStr())
	log.Printf("%v: %v\n", "Время работы программы", time.Since(start))

}
