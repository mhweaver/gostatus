package main

type volumeSegment struct {
	Segment
	output chan string
}

func newVolumeSegment() (segment *volumeSegment) {
	segment = new(volumeSegment)
	segment.output = make(chan string)
	return
}

func (segment *volumeSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *volumeSegment) Run() {
}
