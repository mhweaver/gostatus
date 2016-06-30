package main

import (
	"bufio"
	"log"
	"os"
)

type stdinSegment struct {
	Segment
	output chan string
	reader *bufio.Reader
}

func newStdinSegment() (segment *stdinSegment) {
	segment = new(stdinSegment)
	segment.output = make(chan string)
	segment.reader = bufio.NewReader(os.Stdin)
	return
}

func (segment *stdinSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *stdinSegment) Run() {
	for {
		line, _, err := segment.reader.ReadLine()
		if err != nil {
			log.Fatal("Unable to read line from stdin")
		}

		segment.output <- string(line[:])
	}
}
