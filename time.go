package main

import (
	"time"
)

type timeSegment struct {
	output chan string
}

func newTimeSegment() (s *timeSegment) {
	s = new(timeSegment)
	s.output = make(chan string)
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
	return "%{U#66BA00}%{+o}%{F#66BA00}⌚ %{F-}" + now + "%{-o}%{U-}"
}
