package demo

import (
	"fmt"
	"runtime"
	"testing"
)

func TestDemoSubstring1(t *testing.T) {
	PrintMemUsage()
	DemoSubstring1()
	runtime.GC()
	fmt.Printf("substring0: len=%d\n", len(substring0))
	PrintMemUsage()
}

func TestDemoSubstring2(t *testing.T) {
	PrintMemUsage()
	DemoSubstring2()
	runtime.GC()
	fmt.Printf("substring0: len=%d\n", len(substring0))
	PrintMemUsage()
}
