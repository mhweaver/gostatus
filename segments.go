package main

func loadSegments() (segments []Segment) {
	return []Segment{
		newMailSegment(),
		newDriveSpaceSegment("#999999"),
		newTempSegment(),
		newNetworkSegment(),
		newCpuPercentSegment("#0073BA"),
		newAvgLoadSegment("#0073BA"),
		newMemSegment(),
		newVolumeSegment(),
		newDateSegment(),
		newTimeSegment(),
		//newTestSegment(),
	}
}
