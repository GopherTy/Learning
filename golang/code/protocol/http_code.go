package protocol

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// HTTPServer Go语言中的http服务器
func HTTPServer() {
	// 注册回调
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	log.Println(" http server running ...")

	// 监听端口
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// MyHTTPServer 自定义类型实现 http.Handler 接口
func MyHTTPServer() {
	// 自定义类型实现 handler 接口
	http.Handle("/", &myserver{})

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type myserver struct {
}

func (myserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

// MyHTTPServer1 自定义 http.Handler
func MyHTTPServer1() {
	// 此种写法是不带路由管理的一个handler
	err := http.ListenAndServe(":4000", &myserver{})
	if err != nil {
		log.Fatal(err)
	}
}

// MyHTTPServer2 自定义 ServeMux
func MyHTTPServer2() {
	mux := http.NewServeMux()
	mux.Handle("/", &myserver{})

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

// MyHTTPServer3 自定义 http.Server对象,设置了一个超时函数
func MyHTTPServer3() {
	mux := http.NewServeMux()
	mux.Handle("/", &myserver{})
	mux.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("Timeout"))
	})

	server := &http.Server{
		Addr:         ":4000",
		Handler:      mux,
		WriteTimeout: 2 * time.Second,
	}

	server.ListenAndServe()
}

// Stop 如果服务器端要进行更新，此时还有请求没有完成
// 可以通过以下方式优雅的停止服务端
func Stop() {
	mux := http.NewServeMux()
	mux.Handle("/", &myserver{})

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	// 创建系统信号接收器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	log.Println("Starting HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}
