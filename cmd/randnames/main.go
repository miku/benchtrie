package main

import (
	"flag"
	"fmt"

	"github.com/miku/benchtrie"
)

var n = flag.Int("n", 10, "number of hostnames to generate")

func main() {
	flag.Parse()
	for _, name := range benchtrie.GenerateHosts(*n) {
		fmt.Println(name)
	}
}
