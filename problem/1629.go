package problem

func slowestKey(releaseTimes []int, keysPressed string) byte {
	max, k := releaseTimes[0], keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		if max < releaseTimes[i]-releaseTimes[i-1] {
			max = releaseTimes[i] - releaseTimes[i-1]
			k = keysPressed[i]

		} else if max == releaseTimes[i]-releaseTimes[i-1] {
			if k < keysPressed[i] {
				k = keysPressed[i]
			}
		}
	}
	return k
}
