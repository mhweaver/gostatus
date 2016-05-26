package main

func loadSegments() (segments []Segment) {
	return []Segment{
		newMailSegment(),
		newDriveSpaceSegment(),
		newTempSegment(),
		newNetworkSegment(),
		newCpuPercentSegment(),
		newAvgLoadSegment(),
		newMemSegment(),
		newVolumeSegment(),
		newDateSegment(),
		newTimeSegment(),
		//newTestSegment(),
	}
}
