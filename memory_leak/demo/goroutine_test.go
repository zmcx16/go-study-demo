package demo

import (
	"runtime"
	"testing"
)

func TestDemoGoroutineLeak1(t *testing.T) {

	PrintMemUsage()
	DemoGoroutineLeak1()
	runtime.GC()
	PrintMemUsage()
}

func TestDemoGoroutineNoLeak(t *testing.T) {

	PrintMemUsage()
	DemoGoroutineNoLeak()
	runtime.GC()
	PrintMemUsage()
}
