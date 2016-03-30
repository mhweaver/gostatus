package main

type networkSegment struct {
	output chan string
}

func newNetworkSegment() (segment *networkSegment) {
	segment = new(networkSegment)
	segment.output = make(chan string)
	return
}

func (segment *networkSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *networkSegment) Run() {
}
