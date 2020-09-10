package broadcast

import (
	"errors"
	"fmt"
)

// Broadcaster .
type Broadcaster interface {
	Observer(chan interface{})
	CancelOberserver(chan interface{})
	Submit(v interface{})
	Run()
	Close() error
}

type broadcaster struct {
	input     chan interface{}
	observers map[chan interface{}]bool
	done      chan struct{}
}

func (b *broadcaster) Run() {
	if b == nil {
		b = &broadcaster{
			input:     make(chan interface{}),
			observers: make(map[chan interface{}]bool),
			done:      make(chan struct{}),
		}
	}
	go func() {
		for {
			select {
			case msg := <-b.input:
				// 分发消息到订阅者中
				for obs := range b.observers {
					obs <- msg
				}
			case <-b.done:
				fmt.Println("broadcaser closed")
				return
			}
		}
	}()
}

func (b *broadcaster) Close() error {
	if b == nil {
		return errors.New("broadcaster is nil")
	}
	// if _, ok := <-b.done; !ok {
	// 	return errors.New("broadcaster is closed")
	// }
	// select{
	// case <-b.d:
	// }
	close(b.done)
	return nil
}

func (b *broadcaster) Submit(v interface{}) {
	if b == nil {
		b = &broadcaster{
			input:     make(chan interface{}),
			observers: make(map[chan interface{}]bool),
			done:      make(chan struct{}),
		}
	}
	// 发送消息到广播中
	b.input <- v
}

func (b *broadcaster) Observer(ch chan interface{}) {
	if b == nil {
		b = &broadcaster{
			input:     make(chan interface{}),
			observers: make(map[chan interface{}]bool),
			done:      make(chan struct{}),
		}
	}
	select {
	case <-b.done:
		return
	default:
		// 将订阅者增加到列表中
		b.observers[ch] = true
	}
}

func (b *broadcaster) CancelOberserver(ch chan interface{}) {
	if b == nil {
		return
	}
	select {
	case <-b.done:
		return
	default:
		close(ch)
		delete(b.observers, ch)
	}
}

// NewBroadCast .
func NewBroadCast(buf int) Broadcaster {
	b := &broadcaster{
		input:     make(chan interface{}, buf),
		observers: make(map[chan interface{}]bool),
		done:      make(chan struct{}),
	}
	return b
}
