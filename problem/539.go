package problem

import (
	"math"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	ans := math.MaxInt32
	for i := 0; i < len(timePoints)-1; i++ {
		for j := i + 1; j < len(timePoints); j++ {
			time := timeSub(timePoints[i], timePoints[j])
			if time == 0 {
				return 0
			}
			if time < ans {
				ans = time
			}
		}
	}
	return ans
}

func timeSub(time1S, time2S string) int {
	if time1S == time2S {
		return 0
	}
	time1 := hToM(time1S)
	time2 := hToM(time2S)
	if abs(time1-time2) > 1440-abs(time1-time2) {
		return 1440 - abs(time1-time2)
	}
	return abs(time1 - time2)
}

func hToM(timeS string) int {
	times := strings.Split(timeS, ":")
	timeH, _ := strconv.Atoi(times[0])
	time1M, _ := strconv.Atoi(times[1])
	return timeH*60 + time1M
}
