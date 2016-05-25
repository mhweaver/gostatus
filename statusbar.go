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
	segments := loadSegments()

	// Set up a channel to watch for updates. This makes it so we can block on
	// a single channel to wait for updates from an arbitrary number of
	// segments/channels without segments needing to know about a shared channel.
	updatedSegmentBuffer := make(chan segmentUpdate)
	for _, segment := range segments {
		go func(segment Segment) {
			for output := range segment.GetOutputBuffer() {
				updatedSegmentBuffer <- segmentUpdate{output, segment}
			}
		}(segment)
		go segment.Run()
	}

	segmentOutputs := make(map[Segment]string)
	for {
		// Block until one of the segments updates
		update := <-updatedSegmentBuffer

		segmentOutputs[update.segment] = update.output
		printStatus(segments, segmentOutputs)
	}
}

// Output each segment's output in the order each segment occurs in segments
func printStatus(segments []Segment, outputs map[Segment]string) {
	orderedOutputs := make([]string, 0)
	for _, segment := range segments {
		if outputs[segment] != "" {
			orderedOutputs = append(orderedOutputs, outputs[segment])
		}
	}
	fmt.Println(strings.Join(orderedOutputs, "    "))
}
