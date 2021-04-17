package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemo_Demo_Pass(t *testing.T) {
	assert.True(t, true)
}

func BenchmarkDemo(b *testing.B) {
	b.ResetTimer()
	b.StopTimer()
}

func BenchmarkRepeatTextWithStringConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RepeatTextWithStringConcat("If You Work, You Lose!", 100000)
	}
}

func BenchmarkRepeatTextWithBytesBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RepeatTextWithBytesBuffer("If You Work, You Lose!", 100000)
	}
}

func BenchmarkRepeatTextWithBytesCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RepeatTextWithBytesCopy("If You Work, You Lose!", 100000)
	}
}

func BenchmarkRepeatTextWithStringsBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RepeatTextWithStringsBuilder("If You Work, You Lose!", 100000)
	}
}
