package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
	"log"
	"strconv"
	"time"
)

type networkSegment struct {
	Segment
	output     chan string
	lastSample []linuxproc.NetworkStat
	formatter  formatter
	whiteFg    formatter
}

func newNetworkSegment(formatter formatter) (segment *networkSegment) {
	segment = new(networkSegment)
	segment.output = make(chan string)
	segment.lastSample = nil
	segment.formatter = formatter
	segment.whiteFg = formatter.Bare().WrapFgColor(formatter.GetDefaultColor())
	return
}

func (segment *networkSegment) GetOutputBuffer() chan string {
	return segment.output
}

func (segment *networkSegment) Run() {
	interval := 2 * time.Second
	for {
		stats := segment.getSample()

		segment.output <- segment.renderOutput(interval, segment.lastSample, stats)
		segment.lastSample = stats
		time.Sleep(interval)
	}
}

func (segment *networkSegment) renderOutput(interval time.Duration, stats0, stats1 []linuxproc.NetworkStat) string {
	if stats0 == nil {
		return ""
	}
	var lastSample, currSample *linuxproc.NetworkStat
	for _, s := range stats0 {
		if s.Iface == "eth0" {
			lastSample = &s
			break
		}
	}
	for _, s := range stats1 {
		if s.Iface == "eth0" {
			currSample = &s
			break
		}
	}
	if lastSample == nil || currSample == nil {
		log.Println("Missing interface information")
		return ""
	}
	rxSpeedBps := float64(currSample.RxBytes-lastSample.RxBytes) / float64(interval/time.Second)
	txSpeedBps := float64(currSample.TxBytes-lastSample.TxBytes) / float64(interval/time.Second)
	rxIcon := segment.whiteFg.Format("")
	txIcon := segment.whiteFg.Format("")
	output := rxIcon + strconv.FormatFloat(rxSpeedBps/1024, 'f', 1, 64) + " KiB/s " +
		txIcon + strconv.FormatFloat(txSpeedBps/1024, 'f', 1, 64) + " KiB/s"
	return segment.formatter.Format(output)

}

func (segment *networkSegment) getSample() (networkStats []linuxproc.NetworkStat) {
	networkStats, err := linuxproc.ReadNetworkStat("/proc/net/dev")
	if err != nil {
		log.Println("failed to read /proc/net/dev")
	}
	return networkStats
}
