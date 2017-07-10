package main

import (
	"bufio"
	"encoding/binary"
	"math"
	"os"
	"strings"
)

type Reader struct {
	stream  *os.File
	scanner *bufio.Scanner
	noise   uint16
	parser  parser
	ret     *chan Result
}

func NewReader(stream *os.File, noise uint16, ret *chan Result) Reader {
	obj := Reader{}
	obj.stream = stream
	obj.noise = noise
	obj.parser = parser{}
	obj.ret = ret

	return obj
}

func (r Reader) readBool() (bool, error) {
	buff := make([]byte, 2)
	if _, err := r.stream.Read(buff); err != nil {
		return false, err
	}

	return binary.LittleEndian.Uint16(buff) > r.noise, nil
}

func (r Reader) Read() {
	r.scanner = bufio.NewScanner(r.stream)

	var samples []bool
	var sample bool
	var countPrevSamples uint16 = 0
	var prevSample bool = false

	for {
		samples = []bool{}

		// fast fetch first set bit
		for ; !sample; sample, _ = r.readBool() {
		}

		for {
			samples = append(samples, sample)
			sample, err := r.readBool()
			if err != nil {
				close(*r.ret)
				return
			}

			if sample == false && prevSample == false && countPrevSamples > 300 {
				break
			}

			if sample == false {
				countPrevSamples++
			} else {
				countPrevSamples = 0
			}

			prevSample = sample
		}

		if len(samples) > 1500 {
			r.processSamples(&samples)
		}
	}
}

func (r Reader) processSamples(samples *[]bool) {
	var buffer string
	var prevSample bool = false
	var countPrevSamples uint16 = 0

	for _, sample := range *samples {
		if sample == prevSample {
			countPrevSamples++
			continue
		}
		var rate int = int(math.Ceil(float64(countPrevSamples) / 6))
		if prevSample {
			buffer += strings.Repeat("1", rate)
		} else {
			buffer += strings.Repeat("0", rate)
		}

		prevSample = sample
		countPrevSamples = 0
	}

	buffer = strings.Trim(buffer, "0")
	if len(buffer) > 252 && len(buffer) < 264 {
		buffer += strings.Repeat("0", 264-len(buffer))
	}

	if len(buffer) != PACKAGE_LENGTHS {
		return
	}

	res, err := r.parser.parse(buffer)
	if err == nil {
		*r.ret <- res
	}
}
