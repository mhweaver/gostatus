package main

type memSegment struct {
	output chan string
}

func newMemSegment() (segment *memSegment) {
	segment = new(memSegment)
	segment.output = make(chan string)
	return
}

func (segment *memSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *memSegment) Run() {
}
