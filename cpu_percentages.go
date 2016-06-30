package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
	"log"
	"strconv"
	"strings"
	"time"
)

type cpuPercentSegment struct {
	Segment
	output    chan string
	formatter formatter
	redFg     formatter
	orangeFg  formatter
	defaultFg formatter
}

func newCpuPercentSegment(formatter formatter) (segment *cpuPercentSegment) {
	segment = new(cpuPercentSegment)
	segment.output = make(chan string)
	segment.formatter = formatter
	bare := formatter.Bare()
	segment.redFg = bare.WrapFgColor("#aa0000")
	segment.orangeFg = bare.WrapFgColor("#ee7600")
	segment.defaultFg = bare.WrapFgColor(bare.GetDefaultColor())
	return
}

func (segment *cpuPercentSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *cpuPercentSegment) Run() {
	for {
		// Based on http://stackoverflow.com/questions/11356330/getting-cpu-usage-with-golang/17783687#17783687
		idle0, total0 := segment.getSample()
		time.Sleep(3 * time.Second)
		idle1, total1 := segment.getSample()
		segment.output <- segment.renderOutput(getPercentages(idle0, total0, idle1, total1))
	}
}

func (segment *cpuPercentSegment) renderOutput(percentages []float64) (s string) {
	percentageStrings := make([]string, len(percentages))
	for i, percentage := range percentages {
		var percentageFormatter formatter
		switch {
		case percentage >= 85:
			percentageFormatter = segment.redFg
		case percentage >= 50:
			percentageFormatter = segment.orangeFg
		default:
			percentageFormatter = segment.defaultFg
		}
		percentageStrings[i] = percentageFormatter.Format(strconv.FormatFloat(percentage, 'f', 0, 64)) + "%"
	}
	return segment.formatter.Format(strings.Join(percentageStrings[:4], " "))
}

func (segment *cpuPercentSegment) getSample() (idle, total []uint64) {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Print("stat read failed")
	}
	total = make([]uint64, len(stat.CPUStats))
	idle = make([]uint64, len(stat.CPUStats))
	for i, cpuStat := range stat.CPUStats {
		total[i] += cpuStat.User + cpuStat.Nice + cpuStat.System + cpuStat.IOWait + cpuStat.IRQ + cpuStat.SoftIRQ + cpuStat.Steal + cpuStat.Guest + cpuStat.GuestNice + cpuStat.Idle
		idle[i] = cpuStat.Idle
	}
	return
}

func getPercentages(idle0, total0, idle1, total1 []uint64) (percentages []float64) {
	total := make([]float64, len(idle0))
	idle := make([]float64, len(idle0))
	percentages = make([]float64, len(idle0))
	for i := range idle0 {
		total[i] = float64(total1[i] - total0[i])
		idle[i] = float64(idle1[i] - idle0[i])
		percentages[i] = 100 * (total[i] - idle[i]) / total[i]
	}
	return
}
