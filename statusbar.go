package main

import (
	"fmt"
)

type segmentUpdate struct {
	output  string
	segment segment
}

func main() {
	segments := loadSegments()

	// Set up a channel to watch for updates. This makes it so we can block on
	// a single channel to wait for updates from an arbitrary number of
	// segments/channels without segments needing to know about a shared channel.
	updatedSegmentBuffer := make(chan segmentUpdate)
	for _, segment := range segments {
		go func() {
			for output := range segment.getOutputBuffer() {
				updatedSegmentBuffer <- segmentUpdate{output, segment}
			}
		}()
	}

	segmentOutputs := make(map[segment]string)
	for {
		// Block until one of the segments updates
		update := <-updatedSegmentBuffer

		segmentOutputs[update.segment] = update.output
		printStatus(segments, segmentOutputs)
	}
}

func loadSegments() (segments []segment) {
	segments = make([]segment, 0)
	return
}

// Output each segment's output in the order each segment occurs in segments
func printStatus(segments []segment, outputs map[segment]string) {
	for _, segment := range segments {
		fmt.Print(outputs[segment])
	}
	fmt.Println()
}
