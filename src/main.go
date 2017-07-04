package main

import (
	"os"
	"fmt"
	"flag"
)

func main () {
	noise := flag.Int("noise", 700, "noise value")
	flag.Parse()

	reader := NewReader(os.Stdin, int16(*noise))

	fmt.Printf("Started with noise %d", int16(*noise))
	var ret = make(chan []int16, 2)

	go reader.read_samples(ret)

	for samples := range ret {
		fmt.Println(samples)
	}

	fmt.Println("done")
}