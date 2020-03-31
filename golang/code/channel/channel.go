package channel

import (
	"fmt"
	"runtime"
	"time"
)

// Test1 channel的基本实现
func Test1() {
	ch := make(chan int)

	go func() {
		time.Sleep(10 * time.Second)
		ch <- 1
	}()

	fmt.Println(<-ch)
}

// Test2 测试使用 runtime 包中的 GOMAXPROCS
func Test2() {
	runtime.GOMAXPROCS(8) // 通过该函数设置并行运行的协程数量。
	go func() {
		x := 1
		for i := 0; i < 1000; i++ {
			x *= x + 1
		}

		fmt.Print(x)
	}()

	go func() {
		x := 1
		for i := 0; i < 1000; i++ {
			x *= x + 1
		}

		fmt.Print(x)
	}()
	go func() {
		x := 1
		for i := 0; i < 1000; i++ {
			x *= x + 1
		}

		fmt.Print(x)
	}()
	go func() {
		x := 1
		for i := 0; i < 1000; i++ {
			x *= x + 1
		}
		fmt.Print(x)
	}()

}

// Test3 测试从 通道获取下一个值
func Test3() {
	ch1 := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
	}()

	if <-ch1 == 0 {
		fmt.Println("0")
	} else if <-ch1 == 1 {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}

	fmt.Println("-------")
}

// Test4 测试无缓存通道是否会阻塞后面的代码执行
func Test4() {
	ch := make(chan int)
	go func(chan int) {
		ch <- 1
		fmt.Println("send 1")

		ch <- 2
		fmt.Println("send 2")

		ch <- 3
		fmt.Println("send 3")
	}(ch)

	time.Sleep(time.Second * 1) // 这里的睡眠可以让上面的协程展示出阻塞效果
	fmt.Println(<-ch)           // 当没有收到 ch 里面的内容 send 1 不会打印
	fmt.Println(<-ch)
}

// Test5 测试带缓存的 channel
func Test5() {
	ch := make(chan int, 10)
	go func() {
		ch <- 1
		fmt.Println("step 1")

		ch <- 2
		fmt.Println("step 2")

		ch <- 3
		fmt.Println("step 3")
	}()
	// 上面的内容应该会全部打印
	time.Sleep(time.Second * 1)
	// fmt.Println(<-ch)
}

// Test6 信号量模式
func Test6() {
	ch := make(chan int)
	go func(chan int) {
		ch <- func() int {
			return 1
		}()
	}(ch)

	ch1 := make(chan bool)

	go func() {
		fmt.Println("some IO operate")
		time.Sleep(time.Second * 1)
		ch1 <- true
	}()

	<-ch1
	fmt.Println("step ----")
	fmt.Println(<-ch)

}

// Test7 测试 sync 包的信号量实现同步
func Test7() {
	ch := make(semaphore)
	go ch.Wait(3)
	go ch.P(3)
	fmt.Println("1")

}

// Empty test7 的测试用例
type Empty interface{}
type semaphore chan Empty

func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}
