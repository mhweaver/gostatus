package main

type tempSegment struct {
	Segment
	output chan string
	color  string
}

func newTempSegment(color string) (segment *tempSegment) {
	segment = new(tempSegment)
	segment.output = make(chan string)
	segment.color = color
	return
}

func (segment *tempSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *tempSegment) Run() {
}

func (segment *tempSegment) GetColor() string {
	return segment.color
}
