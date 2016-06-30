package main

import (
	"time"
)

type timeSegment struct {
	Segment
	output    chan string
	formatter formatter
}

func newTimeSegment(formatter formatter) (s *timeSegment) {
	s = new(timeSegment)
	s.output = make(chan string)
	s.formatter = formatter
	return
}

func (s *timeSegment) GetOutputBuffer() chan string {
	return s.output
}

func (s *timeSegment) Run() {
	for {
		s.output <- s.buildOutput(time.Now())
		time.Sleep(1 * time.Second)
	}
}

func (s *timeSegment) buildOutput(t time.Time) string {
	now := t.Format("03:04:05 PM")
	return s.formatter.Format(now)
}
