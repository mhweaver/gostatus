package main

func loadSegments() (segments []Segment) {
	return []Segment{
		newMailSegment(),
		newDriveSpaceSegment(),
		newTempSegment(),
		newNetworkSegment(),
		newCpuPercentSegment("#FF9500"),
		newAvgLoadSegment(),
		newMemSegment(),
		newVolumeSegment(),
		newDateSegment(),
		newTimeSegment(),
		//newTestSegment(),
	}
}
