package main

func loadSegments(formatter formatter) (segments []Segment) {
	return []Segment{
		newDriveSpaceSegment(basicFormatter(formatter, "#999999", " ")),
		newNetworkSegment(basicFormatter(formatter, "#008079", " ")),
		newCpuPercentSegment(basicFormatter(formatter, "#0073BA", " ")),
		newAvgLoadSegment(basicFormatter(formatter, "#0073BA", " ")),
		newMemSegment(basicFormatter(formatter, "#BA4700", " ")),
		newDateSegment(basicFormatter(formatter, "#00BAB1", " ")),
		newTimeSegment(basicFormatter(formatter, "#66BA00", " ")),
		//newTestSegment(),
	}
}

func basicFormatter(formatter formatter, themeColor, icon string) formatter {
	iconStr := formatter.WrapFgColor(themeColor).Format(icon)
	return formatter.PrependInner(iconStr).WrapUnderlineColor(themeColor).WrapOverline()
}
