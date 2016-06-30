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
	output           chan string
	numCpu           int
	formatter        formatter
	redFormatter     formatter
	yellowFormatter  formatter
	defaultFormatter formatter
}

func newAvgLoadSegment(formatter formatter) (segment *avgLoadSegment) {
	segment = new(avgLoadSegment)
	segment.output = make(chan string)
	segment.numCpu = runtime.NumCPU()
	segment.formatter = formatter
	bare := formatter.Bare()
	segment.redFormatter = bare.WrapFgColor("#ff0000")
	segment.yellowFormatter = bare.WrapFgColor("#ffff00")
	segment.defaultFormatter = bare.WrapFgColor(bare.GetDefaultColor())
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

		segment.output <- segment.formatter.Format(segment.renderOutput(loadAvg))
		time.Sleep(1 * time.Second)
	}
}

func (segment *avgLoadSegment) renderOutput(loadAvg *linuxproc.LoadAvg) string {
	yellowThreshold := float64(segment.numCpu)
	redThreshold := float64(segment.numCpu * 2)

	return segment.renderSingleLoad(loadAvg.Last1Min, yellowThreshold, redThreshold) +
		" " + segment.renderSingleLoad(loadAvg.Last5Min, yellowThreshold, redThreshold) +
		" " + segment.renderSingleLoad(loadAvg.Last15Min, yellowThreshold, redThreshold)
}

func (segment *avgLoadSegment) renderSingleLoad(load, yellowThreshold, redThreshold float64) string {
	var thresholdFormatter formatter
	switch {
	case load >= redThreshold:
		thresholdFormatter = segment.redFormatter
	case load >= yellowThreshold:
		thresholdFormatter = segment.yellowFormatter
	default:
		thresholdFormatter = segment.defaultFormatter
	}
	return thresholdFormatter.Format(strconv.FormatFloat(load, 'f', 2, 64))
}
