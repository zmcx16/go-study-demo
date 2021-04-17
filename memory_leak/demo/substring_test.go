package demo

import (
	"runtime"
	"testing"
)

func TestDemoSubstring1(t *testing.T) {
	PrintMemUsage()
	DemoSubstring1()
	runtime.GC()
	PrintMemUsage()
}

func TestDemoSubstring2(t *testing.T) {
	PrintMemUsage()
	DemoSubstring2()
	runtime.GC()
	PrintMemUsage()
}
