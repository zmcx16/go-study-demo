package demo

var subslice0 []int

func DemoSubslice1(s1 []int) {
	// Assume the length of s1 is much larger than 30.
	subslice0 = s1[len(s1)-30:]
}

func DemoSubslice2(s1 []int) {
	subslice0 = append([]int(nil), s1[len(s1)-30:]...)
	// Now, the memory block hosting the elements
	// of s1 can be collected if no other values
	// are referencing the memory block.
}