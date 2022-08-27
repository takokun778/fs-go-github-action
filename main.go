package main

import (
	"log"
	"os"
)

func main() {
	env1 := os.Getenv("ENV1")
	env2 := os.Getenv("ENV2")
	log.Printf("env1: %s\n", env1)
	log.Printf("env2: %s\n", env2)
}
