package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
	"log"
	"strconv"
	"time"
)

type memSegment struct {
	Segment
	output chan string
}

func newMemSegment() (segment *memSegment) {
	segment = new(memSegment)
	segment.output = make(chan string)
	return
}

func (segment *memSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *memSegment) Run() {
	for {
		segment.output <- renderOutput(getMemInfo())
		time.Sleep(5 * time.Second)
	}
}

func renderOutput(free, used, total int64) string {
	percentUsed := 100 * float64(used) / float64(total)
	var color string
	switch {
	case percentUsed > 95:
		color = "#ff0000"
	case percentUsed > 90:
		color = "#ffae00"
	case percentUsed > 20:
		color = "#fff600"
	default:
		color = "#ffffff"
	}

	return " " + "%{F" + color + "}" +
		strconv.FormatFloat(float64(used)/1024/1024, 'f', 2, 64) +
		"GiB%{F-} / " + strconv.FormatFloat(float64(total)/1024/1024, 'f', 2, 64) +
		"GiB (%{F" + color + "}" + strconv.FormatFloat(percentUsed, 'f', 2, 64) + "%%{F-})"
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
