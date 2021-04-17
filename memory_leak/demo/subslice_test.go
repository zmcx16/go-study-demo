package demo

import (
	"fmt"
	"runtime"
	"testing"
)

func TestDemoSubslice1(t *testing.T) {
	s := make([]int, 1000000)

	PrintMemUsage()
	DemoSubslice1(s)
	runtime.GC()
	PrintMemUsage()
}

func TestDemoSubslice2(t *testing.T) {
	s := make([]int, 1000000)

	PrintMemUsage()
	DemoSubslice2(s)
	runtime.GC()
	PrintMemUsage()
}

func TestSliceAppend(t *testing.T) {

	// Note. The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
	// https://stackoverflow.com/questions/35276022/unexpected-slice-append-behaviour
	// https://golang.org/pkg/builtin/#append

	var s1 []int
	var s2 []int

	fmt.Printf("s1: len=%d cap=%d %v\n", len(s1), cap(s1), s1)
	fmt.Printf("s2: len=%d cap=%d %v\n", len(s2), cap(s2), s2)
	fmt.Printf("\n")

	s1 = append(s1, 0, 1, 2, 3, 4, 5, 6)
	fmt.Printf("s1: len=%d cap=%d %v\n", len(s1), cap(s1), s1)
	fmt.Printf("s2: len=%d cap=%d %v\n", len(s2), cap(s2), s2)
	fmt.Printf("\n")

	s2 = append(s1, 100)
	fmt.Printf("s1: len=%d cap=%d %v 0:addr=%p\n", len(s1), cap(s1), s1, &s1[0])
	fmt.Printf("s2: len=%d cap=%d %v 0:addr=%p\n", len(s2), cap(s2), s2, &s2[0])
	fmt.Printf("\n")

	s1 = append(s1, 7)
	fmt.Printf("s1: len=%d cap=%d %v 0:addr=%p\n", len(s1), cap(s1), s1, &s1[0])
	fmt.Printf("s2: len=%d cap=%d %v 0:addr=%p\n", len(s2), cap(s2), s2, &s2[0])
	fmt.Printf("\n")

	s1 = append(s1, 8, 9, 10)
	fmt.Printf("s1: len=%d cap=%d %v 0:addr=%p\n", len(s1), cap(s1), s1, &s1[0])
	fmt.Printf("s2: len=%d cap=%d %v 0:addr=%p\n", len(s2), cap(s2), s2, &s2[0])
	fmt.Printf("\n")

	s2 = append(s2, 101, 102, 103)
	fmt.Printf("s1: len=%d cap=%d %v 0:addr=%p\n", len(s1), cap(s1), s1, &s1[0])
	fmt.Printf("s2: len=%d cap=%d %v 0:addr=%p\n", len(s2), cap(s2), s2, &s2[0])
	fmt.Printf("\n")
}
