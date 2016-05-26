package main

func loadSegments() (segments []Segment) {
	return []Segment{
		newMailSegment(),
		newDriveSpaceSegment(),
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
