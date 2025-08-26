package check

import (
	"testing"
)

type DummyStruct struct {
    A int
    B string
}

// Test cases for CheckType and related functions
// ...existing code...

func TestSliceStruct(t *testing.T) {
    ds := []DummyStruct{{A: 1, B: "test"}}
    if !SliceStruct(ds) {
        t.Errorf("SliceStruct should return true for slice of struct")
    }

    notStruct := []int{1, 2, 3}
    if SliceStruct(notStruct) {
        t.Errorf("SliceStruct should return false for slice of non-struct")
    }

    notSlice := DummyStruct{A: 1, B: "test"}
    if SliceStruct(notSlice) {
        t.Errorf("SliceStruct should return false for non-slice")
    }
}

func TestSlice(t *testing.T) {
    slice := []int{1, 2, 3}
    if !Slice(slice) {
        t.Errorf("Slice should return true for slice")
    }

    array := [3]int{1, 2, 3}
    if !Slice(array) {
        t.Errorf("Slice should return true for array")
    }

    notSlice := 123
    if Slice(notSlice) {
        t.Errorf("Slice should return false for non-slice/array")

	}
}


// Benchmark tests for CheckType and related functions
func BenchmarkCheckType_Struct(b *testing.B) {
    ds := DummyStruct{A: 1, B: "test"}
    for i := 0; i < b.N; i++ {
        CheckType(ds)
    }
}

func BenchmarkCheckType_SliceStruct(b *testing.B) {
    ds := []DummyStruct{{A: 1, B: "test"}}
    for i := 0; i < b.N; i++ {
        CheckType(ds)
    }
}

func BenchmarkSliceStruct(b *testing.B) {
    ds := []DummyStruct{{A: 1, B: "test"}}
    for i := 0; i < b.N; i++ {
        SliceStruct(ds)
    }
}

func BenchmarkSlice(b *testing.B) {
    ds := []int{1, 2, 3}
    for i := 0; i < b.N; i++ {
        Slice(ds)
    }
}