package main

func loadSegments() (segments []Segment) {
	return []Segment{
		newMailSegment(),
		newDriveSpaceSegment(),
		newTempSegment(),
		newNetworkSegment(),
		newCpuSegment(),
		newMemSegment(),
		newVolumeSegment(),
		newDateSegment(),
		newTimeSegment(),
		//newTestSegment(),
	}
}
