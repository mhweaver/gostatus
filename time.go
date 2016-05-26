package main

import (
	"time"
)

type timeSegment struct {
	Segment
	output chan string
	color  string
}

func newTimeSegment(color string) (s *timeSegment) {
	s = new(timeSegment)
	s.output = make(chan string)
	s.color = color
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

func (s *timeSegment) GetColor() string {
	return s.color
}

func (s *timeSegment) buildOutput(t time.Time) string {
	now := t.Format("03:04:05 PM")
	return "%{U" + s.color + "}%{+o}%{F" + s.color + "} %{F-}" + now + "%{-o}%{U-}"
}
