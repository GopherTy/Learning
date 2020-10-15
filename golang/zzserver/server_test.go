package zzserver

import (
	"io"
	"reflect"
	"sync"
	"testing"
)

func TestRequst(t *testing.T) {
	// run server
	go Server()

	// 一个客户端请求多次
	s := struct {
		in       []string
		reqTimes int // request times
		out      map[int][]bool
	}{
		in:       []string{"hello", "world", "zz", "11", "22", "33"},
		reqTimes: 6,
		out: map[int][]bool{
			0: {false, false, false, false, false, false},
			1: {true, true, true, true, true, true},
			2: {true, true, true, true, true, true},
			3: {true, true, true, true, true, true},
			4: {true, true, true, true, true, true},
			5: {true, true, true, true, true, true},
		},
	}

	for i := 0; i < s.reqTimes; i++ {
		rs, err := Clinet1(s.in)
		if err != nil && err != io.EOF {
			t.Fatalf("test fail: dail server error ---> %v\n", err)
			return
		}

		if !reflect.DeepEqual(s.out[i], rs) {
			t.Fatalf("test fail: input is ---> %v, result ---> %v\n", s.out[i], rs)
			return
		}
	}

	// 多客户端同时请求
	s1 := struct {
		in     []string
		client int // client numbers
	}{
		in:     []string{"hello", "world", "11"},
		client: 20,
	}

	var wg sync.WaitGroup
	for i := 0; i < s1.client; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			rs, err := Clinet1(s.in)
			if err != nil && err != io.EOF {
				t.Fatalf("test fail: dail server error ---> %v\n", err)
				return
			}

			m.mu.Lock()
			defer m.mu.Unlock()
			// 判断之前请求的字符串是否存在
			for _, v := range s1.in {
				if !m.data[v] {
					t.Fatalf("test fail: input is ---> %v, result ---> %v\n", s1.in, rs)
					return
				}
			}
		}(i)
	}

	wg.Wait()
}
