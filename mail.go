package main

type mailSegment struct {
	Segment
	output chan string
}

func newMailSegment() (segment *mailSegment) {
	segment = new(mailSegment)
	segment.output = make(chan string)
	return
}

func (segment *mailSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *mailSegment) Run() {
}
