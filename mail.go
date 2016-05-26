package main

type mailSegment struct {
	Segment
	output chan string
	color  string
}

func newMailSegment(color string) (segment *mailSegment) {
	segment = new(mailSegment)
	segment.output = make(chan string)
	segment.color = color
	return
}

func (segment *mailSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *mailSegment) Run() {
}

func (segment *mailSegment) GetColor() string {
	return segment.color
}
