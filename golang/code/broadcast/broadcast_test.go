package broadcast

import (
	"fmt"
	"sync"
	"testing"
)

func TestBroadcast(t *testing.T) {
	wg := sync.WaitGroup{}

	b := NewBroadCast(0)
	b.Run()
	defer b.Close()

	for i := 0; i < 5; i++ {
		wg.Add(1)

		cch := make(chan interface{})

		b.Observer(cch)
		go func(i int) {
			if msg, ok := <-cch; ok {
				fmt.Println(i, "------>", msg)
			}
			wg.Done()
		}(i)
	}

	b.Submit(1)
	b.Submit(2)
	wg.Wait()
}
