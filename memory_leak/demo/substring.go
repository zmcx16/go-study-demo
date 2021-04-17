package demo

var substring0 string // a package-level variable

// A demo purpose function.
func doSubstring1(s1 string) {
	substring0 = s1[:50]
	// Now, s0 shares the same underlying memory block
	// with s1. Although s1 is not alive now, but s0
	// is still alive, so the memory block they share
	// couldn't be collected, though there are only 50
	// bytes used in the block and all other bytes in
	// the block become unavailable.
}

func DemoSubstring1() {
	s := RepeatTextWithBytesCopy("El Psy Congroo", 1000000)
	doSubstring1(s)
}

// A demo purpose function.
func doSubstring2(s1 string) {
	substring0 = string([]byte(s1[:50]))
}

func DemoSubstring2() {
	s := RepeatTextWithBytesCopy("El Psy Congroo", 1000000)
	doSubstring2(s)
}
