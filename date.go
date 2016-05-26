package main

import (
	"time"
)

type dateSegment struct {
	Segment
	output chan string
	color  string
}

func newDateSegment(color string) (s *dateSegment) {
	s = new(dateSegment)
	s.output = make(chan string)
	s.color = color
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

func (s *dateSegment) GetColor() string {
	return s.color
}

func (s *dateSegment) buildOutput(t time.Time) string {
	dayOfWeek := t.Format("Mon")
	date := t.Format("2006-01-02")
	return "%{U" + s.color + "}%{+o}%{F" + s.color + "} %{F-}" + dayOfWeek + " " + date + "%{-o}%{U-}"
}
