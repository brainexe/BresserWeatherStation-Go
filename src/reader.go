package main

import (
	"os"
	"bufio"
	"encoding/binary"
	"math"
	"strings"
	"io"
)

type Reader struct {
	stream *os.File
	scanner *bufio.Scanner
	noise uint16
	parser parser
	ret chan Result
}

func NewReader(stream *os.File, noise uint16, ret chan Result) Reader {
	obj := Reader{}
	obj.stream = stream
	obj.noise = noise
	obj.parser = parser{}
	obj.ret = ret

	return obj
}

func (r Reader) readByte() uint16 {
	var i int16

	e := binary.Read(r.stream, binary.LittleEndian, &i)
	if e == io.EOF {
		panic("End of file");
	}

	return uint16(i) // todo matze
}

func (r Reader) read() {
	r.scanner = bufio.NewScanner(r.stream)

	var samples []uint16
	var sample uint16
	var count_prev_samples uint16 = 0
	var prev_sample uint16 = 0
	count_prev_samples++

	for {
		samples = []uint16{}
		sample = r.readByte()
		for sample <= r.noise {
			sample = r.readByte()
		}

		for {
			samples = append(samples, sample)
			sample = r.readByte()

			if sample <= r.noise {
				sample = 0
			}

			if sample == 0 && prev_sample == 0 && count_prev_samples > 300 {
				break
			}

			if sample == prev_sample {
				count_prev_samples += 1
			} else {
				count_prev_samples = 0
			}

			prev_sample = sample
		}

		if len(samples) > 1500 {
			r.processSamples(samples)
		}
	}
}

func (r Reader) processSamples(sample []uint16) {
	var packet []bool
	var buffer string

	var prevSample bool = false
	var countPrevSamples uint16 = 0

	for _, byte := range sample {
		packet = append(packet, byte > 0)
	}

	for _, sample := range packet {
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
		buffer += strings.Repeat("0", (264 - len(buffer)))
	}

	if len(buffer) != 264 {
		return
	}

	res := r.parser.parse(buffer)
	if res.stationId != ""  {
		r.ret <- res
	}
}
