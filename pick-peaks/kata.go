package kata

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	peaks := PosPeaks{
		Pos:   make([]int, 0),
		Peaks: make([]int, 0),
	}
	peakIndex := 0

	for i := 1; i < len(array); i++ {

		if array[i] > array[i-1] {
			peakIndex = i
			continue
		}

		if array[i] < array[i-1] && peakIndex > 0 {
			peaks.Pos = append(peaks.Pos, peakIndex)
			peaks.Peaks = append(peaks.Peaks, array[peakIndex])
			peakIndex = 0
		}
	}
	return peaks
}
