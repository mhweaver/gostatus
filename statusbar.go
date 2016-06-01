package main

import (
	"fmt"
	"strings"
)

type segmentUpdate struct {
	output  string
	segment Segment
}

func main() {
	rightSegments := loadSegments()
	leftSegments := []Segment{newStdinSegment("-")}

	// Set up a channel to watch for updates. This makes it so we can block on
	// a single channel to wait for updates from an arbitrary number of
	// segments/channels without segments needing to know about a shared channel.
	updatedSegmentBuffer := make(chan segmentUpdate)
	addToUpdateBuffer := func(segment Segment) {
		for output := range segment.GetOutputBuffer() {
			updatedSegmentBuffer <- segmentUpdate{output, segment}
		}
	}
	for _, segment := range append(rightSegments, leftSegments...) {
		go addToUpdateBuffer(segment)
		go segment.Run()
	}

	segmentOutputs := make(map[Segment]string)
	for {
		// Block until one of the segments updates
		update := <-updatedSegmentBuffer

		segmentOutputs[update.segment] = update.output
		printStatus(leftSegments, rightSegments, segmentOutputs)
	}
}

// Output each segment's output in the order each segment occurs in segments
func printStatus(leftSegments, rightSegments []Segment, outputs map[Segment]string) {
	buildOrderedOutputs := func(segments []Segment) []string {
		orderedOutputs := make([]string, 0)
		for _, segment := range segments {
			if outputs[segment] != "" {
				segmentOutput := "%{U" + segment.GetColor() + "}%{+o}" + outputs[segment] + "%{-o}%{U-}"
				orderedOutputs = append(orderedOutputs, segmentOutput)
			}
		}
		return orderedOutputs
	}
	orderedRightOutputs := buildOrderedOutputs(rightSegments)
	orderedLeftOutputs := buildOrderedOutputs(leftSegments)

	separator := "  "
	outputFromRightSegments := "%{r}" + strings.Join(orderedRightOutputs, separator)
	outputFromLeftSegments := "%{l}" + strings.Join(orderedLeftOutputs, separator)

	output := outputFromLeftSegments + outputFromRightSegments
	fmt.Println("%{Sf}" + output + "%{Sl}" + output)
}
