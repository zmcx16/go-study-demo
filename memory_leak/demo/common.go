package demo

import (
	"fmt"
	"runtime"
)

func RepeatTextWithBytesCopy(text string, count int) string {
	bs := make([]byte, len(text)*count)
	bl := 0
	for n := 0; n < count; n++ {
		bl += copy(bs[bl:], text)
	}
	return string(bs)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
