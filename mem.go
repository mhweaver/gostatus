package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
	"log"
	"strconv"
	"time"
)

type memSegment struct {
	Segment
	output                chan string
	formatter             formatter
	lowUsageFormatter     formatter
	highUsageFormatter    formatter
	higherUsageFormatter  formatter
	highestUsageFormatter formatter
}

func newMemSegment(formatter formatter) (segment *memSegment) {
	segment = new(memSegment)
	segment.output = make(chan string)
	segment.formatter = formatter
	bare := formatter.Bare()
	segment.lowUsageFormatter = bare.WrapFgColor(bare.GetDefaultColor())
	segment.highUsageFormatter = bare.WrapFgColor("#fff600")
	segment.higherUsageFormatter = bare.WrapFgColor("#ffae00")
	segment.highestUsageFormatter = bare.WrapFgColor("#ff0000")
	return
}

func (segment *memSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *memSegment) Run() {
	for {
		segment.output <- segment.formatter.Format(segment.renderOutput(getMemInfo()))
		time.Sleep(5 * time.Second)
	}
}

func (segment *memSegment) renderOutput(free, used, total int64) string {
	percentUsed := 100 * float64(used) / float64(total)
	var usageFormatter formatter
	switch {
	case percentUsed > 95:
		usageFormatter = segment.highestUsageFormatter
	case percentUsed > 90:
		usageFormatter = segment.higherUsageFormatter
	case percentUsed > 80:
		usageFormatter = segment.highUsageFormatter
	default:
		usageFormatter = segment.lowUsageFormatter
	}

	return usageFormatter.Format(strconv.FormatFloat(float64(used)/1024/1024, 'f', 2, 64)+"GiB") +
		"/ " + strconv.FormatFloat(float64(total)/1024/1024, 'f', 2, 64) +
		"GiB (" + usageFormatter.Format(strconv.FormatFloat(percentUsed, 'f', 2, 64)+"%") + ")"
}

func getMemInfo() (free, used, total int64) {
	memInfo, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Println("failed to read /proc/meminfo")
	}

	total = int64(memInfo.MemTotal)
	free = int64(memInfo.MemFree + memInfo.Buffers + memInfo.Cached)
	used = int64(total - free)
	return
}
