package demo

import (
	"fmt"
	"runtime"
	"testing"
)

func TestDemoSliceRef1(t *testing.T) {

	PrintMemUsage()
	s := DemoSliceRef1()
	PrintMemUsage()
	runtime.GC()
	fmt.Printf("s1: len=%d cap=%d %v\n", len(s), cap(s), s)
	PrintMemUsage()
}

func TestDemoSliceRef2(t *testing.T) {

	PrintMemUsage()
	s := DemoSliceRef2()
	PrintMemUsage()
	runtime.GC()
	fmt.Printf("s1: len=%d cap=%d %v\n", len(s), cap(s), s)
	PrintMemUsage()
}
