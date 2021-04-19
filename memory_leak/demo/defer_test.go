package demo

import (
	"testing"
)

func TestDemoDefer0(t *testing.T) {
	DemoDefer1()
}

func TestDemoDefer1(t *testing.T) {
	DemoDefer2()
}
