package problem

var week = []string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}

func DayOfTheWeek(day int, month int, year int) string {
	days := day

	for i := 1; i < month; i++ {
		switch i {
		case 4, 6, 9, 11:
			days += 30
		case 2:
			if year%400 == 0 || year%4 == 0 && year%100 != 0 {
				days += 29
			} else {
				days += 28
			}
		case 1, 3, 5, 7, 8, 10, 12:
			days += 31
		}
	}

	for i := 1971; i < year; i++ {
		if i%400 == 0 || i%4 == 0 && i%100 != 0 {
			days += 366
			continue
		}
		days += 365
	}

	return week[days%7]
}
