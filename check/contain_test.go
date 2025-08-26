package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsConsecutiveDates(t *testing.T) {
	tests := []struct {
		name     string
		dates    []string
		n        int
		expected bool
	}{
		{
			name:     "3 hari berturut-turut",
			dates:    []string{"2025-08-15", "2025-08-16", "2025-08-17"},
			n:        3,
			expected: true,
		},
		{
			name:     "tidak berturut-turut",
			dates:    []string{"2025-08-15", "2025-08-17", "2025-08-18"},
			n:        3,
			expected: false,
		},
		{
			name:     "ada 4 tanggal, 3 berturut-turut",
			dates:    []string{"2025-08-15", "2025-08-16", "2025-08-17", "2025-08-19"},
			n:        3,
			expected: true,
		},
		{
			name:     "kurang dari n tanggal",
			dates:    []string{"2025-08-15", "2025-08-16"},
			n:        3,
			expected: false,
		},
		{
			name:     "tanggal duplikat",
			dates:    []string{"2025-08-15", "2025-08-15", "2025-08-16", "2025-08-17"},
			n:        3,
			expected: true,
		},
		{
			name:     "tanggal acak tidak urut",
			dates:    []string{"2025-08-17", "2025-08-15", "2025-08-16"},
			n:        3,
			expected: true,
		},
		{
			name:     "semua tanggal sama",
			dates:    []string{"2025-08-15", "2025-08-15", "2025-08-15"},
			n:        3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsConsecutiveDates(tt.dates, tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsTrendingWeeksStreak(t *testing.T) {
	tests := []struct {
		name   string
		dates  []string
		n      int
		expect bool
	}{
		{
			name:   "empty slice always false",
			dates:  []string{},
			n:      1,
			expect: false,
		},
		{
			name:   "single week, n=1 true",
			dates:  []string{"2025-08-01"},
			n:      1,
			expect: true,
		},
		{
			name:   "single week, n=2 false",
			dates:  []string{"2025-08-01"},
			n:      2,
			expect: false,
		},
		{
			name: "two consecutive weeks, n=2 true",
			dates: []string{
				"2025-08-04", // week 32
				"2025-08-12", // week 33
			},
			n:      2,
			expect: true,
		},
		{
			name: "three consecutive weeks, n=3 true",
			dates: []string{
				"2025-07-28", // week 31
				"2025-08-05", // week 32
				"2025-08-12", // week 33
			},
			n:      3,
			expect: true,
		},
		{
			name: "non consecutive weeks, n=2 false",
			dates: []string{
				"2025-08-04", // week 32
				"2025-08-19", // week 34
			},
			n:      2,
			expect: false,
		},
		{
			name: "cross year consecutive weeks, n=2 true",
			dates: []string{
				"2024-12-30", // ISO week 1 of 2025
				"2025-01-06", // ISO week 2 of 2025
			},
			n:      2,
			expect: true,
		},
		{
			name: "invalid date ignored, still counts",
			dates: []string{
				"invalid-date",
				"2025-08-01",
			},
			n:      1,
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsConsecutiveWeeksDates(tt.dates, tt.n)
			if got != tt.expect {
				t.Errorf("IsTrendingWeeksStreak(%v, %d) = %v; want %v",
					tt.dates, tt.n, got, tt.expect)
			}
		})
	}
}

func TestGetConsecutiveTrendingWeeks(t *testing.T) {
	tests := []struct {
		name   string
		dates  []string
		expect int
	}{
		{
			name:   "empty slice",
			dates:  []string{},
			expect: 0,
		},
		{
			name:   "single date",
			dates:  []string{"2025-08-01"},
			expect: 1,
		},
		{
			name: "two dates in same week",
			dates: []string{
				"2025-08-04", // Monday
				"2025-08-05", // Tuesday
			},
			expect: 1, // tetap satu minggu
		},
		{
			name: "two consecutive weeks",
			dates: []string{
				"2025-08-04", // week 32
				"2025-08-12", // week 33
			},
			expect: 2,
		},
		{
			name: "three consecutive weeks with duplicates",
			dates: []string{
				"2025-07-28", // week 31
				"2025-08-01", // week 31 duplicate
				"2025-08-05", // week 32
				"2025-08-12", // week 33
			},
			expect: 3,
		},
		{
			name: "non consecutive weeks",
			dates: []string{
				"2025-08-04", // week 32
				"2025-08-19", // week 34
			},
			expect: 1,
		},
		{
			name: "cross year consecutive weeks",
			dates: []string{
				"2024-12-30", // ISO week 1 of 2025 or week 1? careful ISOWeek: 2025-01-01 is week 1
				"2025-01-06", // week 2 of 2025
			},
			expect: 2,
		},
		{
			name: "non parsable date ignored",
			dates: []string{
				"invalid-date",
				"2025-08-01",
			},
			expect: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetConsecutiveTrendingWeeks(tt.dates)
			if got != tt.expect {
				t.Errorf("GetConsecutiveTrendingWeeks(%v) = %d; want %d", tt.dates, got, tt.expect)
			}
		})
	}
}

func TestParseWeekKey(t *testing.T) {
	tests := []struct {
		name string
		key  string
		year int
		week int
	}{
		{
			name: "normal week",
			key:  "2025-32",
			year: 2025,
			week: 32,
		},
		{
			name: "week with leading zero",
			key:  "2025-03",
			year: 2025,
			week: 3,
		},
		{
			name: "week 52 end of year",
			key:  "2024-52",
			year: 2024,
			week: 52,
		},
		{
			name: "week 1 new year",
			key:  "2026-01",
			year: 2026,
			week: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotYear, gotWeek := parseWeekKey(tt.key)
			if gotYear != tt.year || gotWeek != tt.week {
				t.Errorf("parseWeekKey(%q) = (%d, %d); want (%d, %d)",
					tt.key, gotYear, gotWeek, tt.year, tt.week)
			}
		})
	}
}

// Benchmark tests for IsConsecutiveDates

func BenchmarkIsConsecutiveDates(b *testing.B) {
	ds := []string{"2025-08-15", "2025-08-15", "2025-08-16", "2025-08-17"}
	for i := 0; i < b.N; i++ {
		IsConsecutiveDates(ds, 3)
	}
}
