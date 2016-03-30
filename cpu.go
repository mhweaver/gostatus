package main

type cpuSegment struct {
	output chan string
}

func newCpuSegment() (segment *cpuSegment) {
	segment = new(cpuSegment)
	segment.output = make(chan string)
	return
}

func (segment *cpuSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *cpuSegment) Run() {
}
