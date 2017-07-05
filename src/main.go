package main

import (
	"os"
	"fmt"
	"flag"
)

func main () {
	noise := flag.Int("noise", 700, "noise value")
	flag.Parse()

	fmt.Printf("Started with noise %d\n", uint16(*noise))
	var ret = make(chan Result)

	reader := NewReader(os.Stdin, uint16(*noise), &ret)
	formatter := NewFormatter()

	go reader.read()

	for result := range ret {
		fmt.Println(formatter.Format(result))
	}

}