package main

type driveSpaceSegment struct {
	output chan string
}

func newDriveSpaceSegment() (segment *driveSpaceSegment) {
	segment = new(driveSpaceSegment)
	segment.output = make(chan string)
	return
}

func (segment *driveSpaceSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *driveSpaceSegment) Run() {
}
