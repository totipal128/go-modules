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