package main

func loadSegments() (segments []Segment) {
	return []Segment{
		newDriveSpaceSegment("#999999"),
		newNetworkSegment("#008079"),
		newCpuPercentSegment("#0073BA"),
		newAvgLoadSegment("#0073BA"),
		newMemSegment("#BA4700"),
		newDateSegment("#00BAB1"),
		newTimeSegment("#66BA00"),
		//newTestSegment(),
	}
}
