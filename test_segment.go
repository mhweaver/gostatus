package main

import "time"
import "strconv"

type testSegment struct {
	Segment
	output chan string
	count  int
	color  string
}

func newTestSegment() (segment *testSegment) {
	segment = new(testSegment)
	segment.output = make(chan string)
	segment.count = 0
	return
}

func (segment *testSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *testSegment) Run() {
	for {
		time.Sleep(1 * time.Second)
		segment.count++
		segment.output <- strconv.Itoa(segment.count)
	}
}

func (segment *testSegment) GetColor() string {
	return segment.color
}
