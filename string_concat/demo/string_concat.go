package demo

import (
	"bytes"
	"strings"
)

func RepeatTextWithStringConcat(text string, count int) string {
	var s string
	for i := 0; i < count; i++ {
		s += text
	}
	return s
}

func RepeatTextWithBytesBuffer(text string, count int) string {
	var buffer bytes.Buffer
	for i := 0; i < count; i++ {
		buffer.WriteString(text)
	}
	return buffer.String()
}

func RepeatTextWithBytesCopy(text string, count int) string {
	bs := make([]byte, len(text)*count)
	bl := 0
	for n := 0; n < count; n++ {
		bl += copy(bs[bl:], text)
	}
	return string(bs)
}

func RepeatTextWithStringsBuilder(text string, count int) string {
	var sb strings.Builder
	for i := 0; i < count; i++ {
		sb.WriteString(text)
	}
	return sb.String()
}
