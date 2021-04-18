package demo

// more goroutine example: https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html

func DemoGoroutineLeak1() {

	for i := 0; i < 10000; i++ {
		ch := make(chan int)
		go func() {
			_ = <-ch
		}()
	}
}

func DemoGoroutineLeak2() {

	for i := 0; i < 10000; i++ {
		ch := make(chan int)
		go func() {
			_ = <-ch
		}()

		ch <- i
	}
}
