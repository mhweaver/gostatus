package main

func loadSegments() (segments []Segment) {
	return []Segment{
		newMailSegment("-"),
		newDriveSpaceSegment("#999999"),
		newTempSegment("-"),
		newNetworkSegment("#008079"),
		newCpuPercentSegment("#0073BA"),
		newAvgLoadSegment("#0073BA"),
		newMemSegment("#BA4700"),
		newVolumeSegment("-"),
		newDateSegment("#00BAB1"),
		newTimeSegment("#66BA00"),
		//newTestSegment(),
	}
}
