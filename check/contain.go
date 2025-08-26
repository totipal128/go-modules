package check

import (
	"fmt"
	"sort"
	"time"
)

// check condition if array equals str
func ContainString(str string, val []string) bool {
	for _, elm := range val {
		if str == elm {
			return true
		}
	}
	return false
}

// check condition if array equals str
func ContainInt64(str int64, val []int64) bool {
	for _, elm := range val {
		if str == elm {
			return true
		}
	}
	return false
}

// check condition if array equals str
func ContainInt(str int, val []int) bool {
	for _, elm := range val {
		if str == elm {
			return true
		}
	}
	return false
}

// check condition if array equals str
func ContainUint(str uint, val []uint) bool {
	for _, elm := range val {
		if str == elm {
			return true
		}
	}
	return false
}

// IsConsecutiveDates checks if there are N consecutive dates in the given slice.
// dates: slice of date strings (format "2006-01-02")
// n: jumlah hari berturut-turut yang ingin dicek
func IsConsecutiveDates(dates []string, n int) bool {
	if len(dates) < n {
		return false
	}

	var tDates []time.Time
	for _, d := range dates {
		t, err := time.Parse("2006-01-02", d)
		if err == nil {
			tDates = append(tDates, t)
		}
	}
	if len(tDates) < n {
		return false
	}

	sort.Slice(tDates, func(i, j int) bool {
		return tDates[i].Before(tDates[j])
	})

	streak := 1
	for i := 1; i < len(tDates); i++ {
		if tDates[i].Sub(tDates[i-1]) == 24*time.Hour {
			streak++
			if streak >= n {
				return true
			}
		} else if !tDates[i].Equal(tDates[i-1]) {
			streak = 1
		}
	}
	return false
}

// IsTrendingWeeksStreak menerima slice tanggal (format "2006-01-02")
func IsConsecutiveWeeksDates(dates []string, n int) bool {
	return GetConsecutiveTrendingWeeks(dates) >= n
}

// GetConsecutiveTrendingWeeks menerima slice tanggal (format "2006-01-02")
// dan mengembalikan jumlah minggu berturut-turut terpanjang yang memiliki data.
func GetConsecutiveTrendingWeeks(dates []string) int {
	if len(dates) == 0 {
		return 0
	}

	weekMap := make(map[string]struct{})
	for _, d := range dates {
		t, err := time.Parse("2006-01-02", d)
		if err == nil {
			year, week := t.ISOWeek()
			key := fmt.Sprintf("%04d-%02d", year, week)
			weekMap[key] = struct{}{}
		}
	}

	// Kumpulkan semua minggu unik, lalu urutkan
	var weeks []string
	for w := range weekMap {
		weeks = append(weeks, w)
	}
	sort.Strings(weeks)

	// Hitung streak minggu berturut-turut
	maxStreak, streak := 1, 1
	for i := 1; i < len(weeks); i++ {
		prevYear, prevWeek := parseWeekKey(weeks[i-1])
		currYear, currWeek := parseWeekKey(weeks[i])
		if (currYear == prevYear && currWeek == prevWeek+1) ||
			(currYear == prevYear+1 && prevWeek == 52 && currWeek == 1) {
			streak++
			if streak > maxStreak {
				maxStreak = streak
			}
		} else {
			streak = 1
		}
	}
	return maxStreak
}

// Helper untuk parsing key minggu
func parseWeekKey(key string) (int, int) {
	var year, week int
	fmt.Sscanf(key, "%d-%d", &year, &week)
	return year, week
}

func IsConsecutiveMonthsStreak(dates []string, n int) bool {
	return GetConsecutiveTrendingMonths(dates) >= n
}

// GetConsecutiveTrendingMonths menerima slice tanggal (format "2006-01-02")
// dan mengembalikan jumlah bulan berturut-turut terpanjang yang memiliki data.
func GetConsecutiveTrendingMonths(dates []string) int {
	if len(dates) == 0 {
		return 0
	}

	monthMap := make(map[string]struct{})
	for _, d := range dates {
		t, err := time.Parse("2006-01-02", d)
		if err == nil {
			key := fmt.Sprintf("%04d-%02d", t.Year(), int(t.Month()))
			monthMap[key] = struct{}{}
		}
	}

	// Kumpulkan semua bulan unik, lalu urutkan
	var months []string
	for m := range monthMap {
		months = append(months, m)
	}
	sort.Strings(months)

	// Hitung streak bulan berturut-turut
	maxStreak, streak := 1, 1
	for i := 1; i < len(months); i++ {
		prevYear, prevMonth := parseMonthKey(months[i-1])
		currYear, currMonth := parseMonthKey(months[i])
		if (currYear == prevYear && currMonth == prevMonth+1) ||
			(currYear == prevYear+1 && prevMonth == 12 && currMonth == 1) {
			streak++
			if streak > maxStreak {
				maxStreak = streak
			}
		} else {
			streak = 1
		}
	}
	return maxStreak
}

// Helper untuk parsing key bulan
func parseMonthKey(key string) (int, int) {
	var year, month int
	fmt.Sscanf(key, "%d-%d", &year, &month)
	return year, month
}
