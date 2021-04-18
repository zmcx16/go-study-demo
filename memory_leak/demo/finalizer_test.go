package demo

import (
	"runtime"
	"testing"
)

func TestDemoNoFinalizer(t *testing.T) {

	PrintMemUsage()
	DemoNoFinalizer()
	runtime.GC()
	PrintMemUsage()
}

func TestDemoFinalizerLeak(t *testing.T) {

	PrintMemUsage()
	DemoFinalizerLeak()
	runtime.GC()
	PrintMemUsage()
}
