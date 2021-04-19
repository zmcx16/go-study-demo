package demo

import (
	"fmt"
)

func DemoDefer1() {

	for i := 0; i < 10; i++ {

		defer func(index int) {
			fmt.Printf("run defer func%d\n", index)
		}(i)

		fmt.Printf("loop %d end\n", i)
	}
}

func DemoDefer2() {

	for i := 0; i < 10; i++ {

		func() {
			defer func(index int) {
				fmt.Printf("run defer func%d\n", index)
			}(i)

			fmt.Printf("loop %d end\n", i)
		}()
	}
}
