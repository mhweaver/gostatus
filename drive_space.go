package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
	"log"
	"strconv"
	"time"
)

type driveSpaceSegment struct {
	Segment
	output    chan string
	formatter formatter
}

func newDriveSpaceSegment(formatter formatter) (segment *driveSpaceSegment) {
	segment = new(driveSpaceSegment)
	segment.output = make(chan string)
	segment.formatter = formatter
	return
}

func (segment *driveSpaceSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *driveSpaceSegment) Run() {
	for {
		segment.output <- segment.renderOutput()
		time.Sleep(1 * time.Minute)
	}
}

func bToGiB(bytes float64) float64 {
	return bytes / 1024 / 1024 / 1024
}

func (segment *driveSpaceSegment) renderOutput() string {
	disk, err := linuxproc.ReadDisk("/")
	if err != nil {
		log.Println("Unable to get stats on /")
	}
	freeSpace := float64(disk.Free) // in bytes
	freeGiB := bToGiB(freeSpace)
	allSpace := float64(disk.All) // in bytes
	allGiB := bToGiB(allSpace)
	output := strconv.FormatFloat(freeGiB, 'f', 2, 64) + "GiB / " +
		strconv.FormatFloat(allGiB, 'f', 2, 64) + "GiB (free)"

	return segment.formatter.Format(output)
}
