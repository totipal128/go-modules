package check

import (
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