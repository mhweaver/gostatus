package main

import (
	"time"
)

type dateSegment struct {
	Segment
	output    chan string
	formatter formatter
}

func newDateSegment(formatter formatter) (s *dateSegment) {
	s = new(dateSegment)
	s.output = make(chan string)
	s.formatter = formatter
	return
}

func (s *dateSegment) GetOutputBuffer() chan string {
	return s.output
}

func (s *dateSegment) Run() {
	for {
		s.output <- s.buildOutput(time.Now())
		time.Sleep(15 * time.Minute)
	}
}

func (s *dateSegment) buildOutput(t time.Time) string {
	dayOfWeek := t.Format("Mon")
	date := t.Format("2006-01-02")
	return s.formatter.Format(dayOfWeek + " " + date)
}
