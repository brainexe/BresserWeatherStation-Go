package main

import (
	"strconv"
	"fmt"
)

type parser struct {
}

func (p parser) getDigit (hex string, digit int) uint64 {
	number, _ := strconv.ParseUint(hex[digit:digit+1], 16, 10)

	return number
}

func (p parser) parse (raw string) Result {
	hex := ""
	for i := 0; i < 66*4; i+=4 {
		tmp, _ := strconv.ParseInt(raw[i:i+4], 2, 64)
		hex += fmt.Sprintf("%x", tmp)
	}

	result := Result{}


	if hex[0:10] != "aaaaaaaaaa" {
		return result
	}

	fmt.Println(raw)
	fmt.Println(hex)

	//for n in range(0,26):
	//	if ord(self.stream[14+n:15+n]) ^ 0xf != ord(self.stream[40+n:41+n]):
	//	return 4

	temp_digit_2 := p.getDigit(hex, 54)
	temp_decimal := p.getDigit(hex, 55)
	//temp_sign := int(hex[65:66])
	temp_digit_1 := p.getDigit(hex, 57)

	result.temperature = float32(temp_digit_1 * 10 + temp_digit_2 + (temp_decimal/10))

	fmt.Println(result.temperature)
	fmt.Println(result)

	return result
}

