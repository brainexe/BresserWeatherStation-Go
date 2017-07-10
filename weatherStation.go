package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	noise := flag.Int("noise", 700, "noise value")
	stopAtFirst := flag.Bool("stop-at-first", false, "noise value")
	flag.Parse()

	var ret = make(chan Result)

	reader := NewReader(os.Stdin, uint16(*noise), &ret)
	formatter := NewFormatter()

	go reader.Read()

	for result := range ret {
		fmt.Println(formatter.Format(&result))
		if *stopAtFirst {
			break
		}
	}
}
