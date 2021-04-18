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

func TestDemoGoroutineLeak2(t *testing.T) {

	PrintMemUsage()
	DemoGoroutineLeak2()
	runtime.GC()
	PrintMemUsage()
}
