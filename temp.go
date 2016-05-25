package main

type tempSegment struct {
	Segment
	output chan string
}

func newTempSegment() (segment *tempSegment) {
	segment = new(tempSegment)
	segment.output = make(chan string)
	return
}

func (segment *tempSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *tempSegment) Run() {
}
