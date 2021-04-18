package demo

type BigData struct {
	data [1000000]int
}

func DemoSliceRef1() []*BigData {
	s := []*BigData{new(BigData), new(BigData), new(BigData), new(BigData)}
	// do something with s ...

	return s[1:3:3]
}

func DemoSliceRef2() []*BigData {
	s := []*BigData{new(BigData), new(BigData), new(BigData), new(BigData)}
	// do something with s ...

	// Reset pointer values.
	s[0], s[len(s)-1] = nil, nil
	return s[1:3:3]
}
