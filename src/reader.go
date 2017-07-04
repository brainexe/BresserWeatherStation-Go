package main

import (
	"os"
	"bufio"
	"encoding/binary"
	"math"
	"strings"
)


type reader struct {
	stream *os.File
	noise int16
	scanner *bufio.Scanner
}

func NewReader(stream *os.File, noise int16) reader {
	obj := reader{}
	obj.stream = stream
	obj.noise = noise

	return obj
}

func (r reader) read_byte() int16 {
	var i int16

	binary.Read(r.stream, binary.LittleEndian, &i)

	return i
}

func (r reader) read_samples(ret chan []int16) {
	r.scanner = bufio.NewScanner(r.stream)

	var samples []int16
	var sample int16
	var count_prev_samples int = 0
	var prev_sample int16 = 0
	count_prev_samples++

	for {
		sample = r.read_byte()
		for sample <= r.noise {
			sample = r.read_byte()
		}

		for {
			samples = append(samples, sample)
			sample = r.read_byte()

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
		//ret <- samples
		r.process_samples(samples)
		samples = []int16{}
	}
}


func (r reader) process_samples(sample []int16) {
	var packet []bool
	var buffer string

	var prev_sample bool = false
	var count_prev_samples int = 0

	for _, byte := range sample {
		packet = append(packet, byte > 0)
	}

	for _, sample := range packet {
		if sample == prev_sample {
			count_prev_samples++
			continue
		}
		var rate int = int(math.Ceil(float64(count_prev_samples) / 6))
		if prev_sample {
			buffer += strings.Repeat("1", rate)
		} else {
			buffer += strings.Repeat("0", rate)
		}

		prev_sample = sample
		count_prev_samples = 0
	}

	buffer = strings.Trim(buffer, "0")
	if len(buffer) > 252 && len(buffer) < 264 {
		buffer += strings.Repeat("0", (264 - len(buffer)))
	}

	if len(buffer) != 264 {
		return
	}

	r.process_packet(buffer)
}

func (r reader) process_packet(buffer string) {
	parser := parser{}
	parser.parse(buffer)
}

	/* def process_signal(self, samples):

        buffer = ""

        #Normalising the samples
        for index, sample in enumerate(samples):
            if samples[index] > 0:
                #print "%d -> 1" % samples[index]
                samples[index] = 1
            else:
                #print "%d -> 0" % samples[index]
                samples[index] = 0

        prev_sample = 0
        count_prev_samples = 0

        #Reducing the samples
        for sample in samples:
            if sample == prev_sample:
                count_prev_samples += 1
                continue

            #6 is the rate for a 48Khz sampling
            rate = math.ceil(float(count_prev_samples) / 6)
            buffer+=str(prev_sample) * int(rate)

            prev_sample = sample
            count_prev_samples = 0

        #Stripping zeros from the signal
        buffer = buffer.strip("0")

        #If the buffer size is < 256 and > 252, probably the original packet
        #ends with 0s which has been stripped, so let's compensate
        if len(buffer) > 252 and len(buffer) < 264:
            buffer += "0" * (264 - len(buffer))

        self.process_packet(buffer)
        */
