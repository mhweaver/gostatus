package main

type volumeSegment struct {
	Segment
	output chan string
	color  string
}

func newVolumeSegment(color string) (segment *volumeSegment) {
	segment = new(volumeSegment)
	segment.output = make(chan string)
	segment.color = color
	return
}

func (segment *volumeSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *volumeSegment) Run() {
}

func (segment *volumeSegment) GetColor() string {
	return segment.color
}
