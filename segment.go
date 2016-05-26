package main

type Segment interface {
	Run()
	GetOutputBuffer() chan string
	GetColor() string
}
