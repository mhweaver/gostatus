package main

import (
	"fmt"
	"strings"
	"time"
)

type segmentUpdate struct {
	output  string
	segment Segment
}

func main() {
	formatter := NewLemonbarFormatter()
	rightSegments := loadSegments(formatter)
	leftSegments := []Segment{newStdinSegment()}

	// Set up a channel to watch for updates. This makes it so we can block on
	// a single channel to wait for updates from an arbitrary number of
	// segments/channels without segments needing to know about a shared channel.
	updatedSegmentBuffer := make(chan segmentUpdate)
	addToUpdateBuffer := func(segment Segment) {
		for output := range segment.GetOutputBuffer() {
			updatedSegmentBuffer <- segmentUpdate{output, segment}
		}
	}
	segmentStartDelay := 1 * time.Second / time.Duration(len(leftSegments)+len(rightSegments))
	go func() {
		for i, segment := range append(rightSegments, leftSegments...) {
			go addToUpdateBuffer(segment)
			// Delay the segment slightly, to space the segments out a little,
			// so they aren't fighting over the update buffer as much
			time.Sleep(time.Duration(i) * segmentStartDelay)
			go segment.Run()
		}
	}()

	firstScreenFormatter := formatter.SetMonitor("f")
	lastScreenFormatter := formatter.SetMonitor("l")
	alignLeftFormatter := formatter.AlignLeft()
	alignRightFormatter := formatter.AlignRight()
	segmentOutputs := make(map[Segment]string)
	for {
		// Block until one of the segments updates
		update := <-updatedSegmentBuffer

		segmentOutputs[update.segment] = update.output
		output := renderOutput(leftSegments, rightSegments, segmentOutputs, alignLeftFormatter, alignRightFormatter)
		fmt.Println(firstScreenFormatter.Format(output) + lastScreenFormatter.Format(output))
	}
}

// Output each segment's output in the order each segment occurs in segments
func renderOutput(leftSegments, rightSegments []Segment, outputs map[Segment]string, leftFormatter, rightFormatter formatter) string {
	buildOrderedOutputs := func(segments []Segment) []string {
		orderedOutputs := make([]string, 0)
		for _, segment := range segments {
			if outputs[segment] != "" {
				orderedOutputs = append(orderedOutputs, outputs[segment])
			}
		}
		return orderedOutputs
	}
	orderedRightOutputs := buildOrderedOutputs(rightSegments)
	orderedLeftOutputs := buildOrderedOutputs(leftSegments)

	separator := "  "
	outputFromLeftSegments := leftFormatter.Format(strings.Join(orderedLeftOutputs, separator))
	outputFromRightSegments := rightFormatter.Format(strings.Join(orderedRightOutputs, separator))

	return outputFromLeftSegments + outputFromRightSegments
}
