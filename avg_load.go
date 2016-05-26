package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
	"log"
	"runtime"
	"strconv"
	"time"
)

type avgLoadSegment struct {
	Segment
	output chan string
	numCpu int
	color  string
}

func newAvgLoadSegment(color string) (segment *avgLoadSegment) {
	segment = new(avgLoadSegment)
	segment.output = make(chan string)
	segment.numCpu = runtime.NumCPU()
	segment.color = color
	return
}

func (segment *avgLoadSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *avgLoadSegment) Run() {
	for {
		loadAvg, err := linuxproc.ReadLoadAvg("/proc/loadavg")
		if err != nil {
			log.Fatal("unable to read /proc/loadavg")
		}

		segment.output <- segment.renderOutput(loadAvg)
		time.Sleep(1 * time.Second)
	}
}

func (segment *avgLoadSegment) renderOutput(loadAvg *linuxproc.LoadAvg) string {
	yellowThreshold := float64(segment.numCpu)
	redThreshold := float64(segment.numCpu * 2)

	return "%{F" + segment.color + "}%{F-} " + renderSingleLoad(loadAvg.Last1Min, yellowThreshold, redThreshold) +
		" " + renderSingleLoad(loadAvg.Last5Min, yellowThreshold, redThreshold) +
		" " + renderSingleLoad(loadAvg.Last15Min, yellowThreshold, redThreshold)
}

func renderSingleLoad(load, yellowThreshold, redThreshold float64) string {
	var color string
	switch {
	case load >= redThreshold:
		color = "#ff0000"
	case load >= yellowThreshold:
		color = "#ffff00"
	default:
		color = "#00ff00"
	}
	return "%{F" + color + "}" + strconv.FormatFloat(load, 'f', 2, 64) + "%{F-}"
}
