package demo

import (
	"fmt"
	"runtime"
)

func DemoNoFinalizer() {

	type T struct {
		v [1 << 20]int
		t *T
	}

	var x, y T

	x.t, y.t = &y, &x
}

func DemoFinalizerLeak() {

	type T struct {
		v [1 << 20]int
		t *T
	}

	var finalizer = func(t *T) {
		fmt.Println("finalizer called")
	}

	var x, y T

	// The SetFinalizer call makes x escape to heap.
	runtime.SetFinalizer(&x, finalizer)

	// The following line forms a cyclic reference
	// group with two members, x and y.
	// This causes x and y are not collectable.
	x.t, y.t = &y, &x // y also escapes to heap.
}
