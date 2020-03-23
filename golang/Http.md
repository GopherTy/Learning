# `Http`协议

### 一. HTTP 协议

1.它是应用层协议，工作于客户端-服务端架构上，浏览器作为HTTP客户端通过URL向HTTP服务端即WEB服务器发送所有请求。Web服务器根据接收到的请求后，向客户端发送响应信息。

2.偷个懒看到一篇还不错的文章

[关于HTTP协议，一篇就够了](https://www.jianshu.com/p/80e25cb1d81a)

### 二. Go中的HTTP服务器。

```go
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
```

对上面的代码分析：

1.http.HandleFunc 作用是将某个函数与一个路由进行绑定，当用户访问指定路由时，所绑定的函数就会被执行。它的两个参数分别是路由，处理函数。

2.http.HandleFunc函数第二个参数的符合 函数签名func(http.ResponseWriter, *http.Request)其中，第一个参数是响应对象，包括了响应码、响应头、响应体，第二个参数是请求对象，包括请求头（Request Header）、请求体（Request Body）和其它相关的内容。

3.http.ListenAndServe的作用是启动 HTTP 服务器，并监听发送到指定地址和端口的HTTP请求。

4.在 `http.HandleFunc` 函数内部会将我们传入的绑定函数转化为类型 `http.HandlerFunc` 且这个对象实现了 http.Handler 接口

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

5.`http.HandleFunc` 的根本作用是将一个函数转化为一个实现了 `http.Handler` 接口的类型（`http.HandlerFunc`）。所以我们可以自定义类型实现 http.Handler 接口。

```go
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
```

6.服务复用器（ServeMux）,即 http.ListenAndServe 的第二个参数

`func ListenAndServe(addr string, handler Handler) error {...}` 就是一个 handler 对象所以上面的代码可以用我们自定义的实现了 Handler 接口的类型，但是实际开发中很少用到纯粹的该对象，因为它缺失了和路由进行绑定的功能。

```go
// MyHTTPServer 自定义类型实现 http.Handler 接口
func MyHTTPServer() {
    
    err := http.ListenAndServe(":4000", &myserver{})
	if err != nil {
		log.Fatal(err)
	}
}

type myserver struct {
}

func (myserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
```

7.`http.ServeMux` 是Go 语言标准库实现的一个带有基本路由功能的服务复用器（Multiplexer），`http.Handle` `http.HandleFunc` 方法使用的是 `http.DefaultServeMux` 对象默认封装的一个复用器。我们可以通过 [`http.NewServeMux`](https://gowalker.org/net/http#NewServeMux) 创建一个新的 `http.ServeMux` 对象.

```go
// MyHTTPServer 自定义类型实现 http.Handler 接口
func MyHTTPServer() {
    mux := http.NewServeMux()
    mux.Handle("/", &myserver{})
    
    err := http.ListenAndServe(":4000", &myserver{})
	if err != nil {
		log.Fatal(err)
	}
}

type myserver struct {
}

func (myserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
```

以上便是大多数 web 框架的底层用法了，其实就是实现了一个带有路由层的 `http.Handler`,以此为基础提供大量便利的辅助函数。

8.http.Server对象：

```go
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

type myserver struct {
}

func (myserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
```

9.停止服务器：

```go
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

type myserver struct {
}

func (myserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
```

通过捕捉 `os.Interrupt` 信号（Ctrl+C）然后调用 `server.Shutdown` 方法告知服务器应停止接受新的请求并在处理完当前已接受的请求后关闭服务器。为了与普通错误相区别，标准库提供了一个特定的错误类型 `http.ErrServerClosed`，我们可以在代码中通过判断是否为该错误类型来确定服务器是正常关闭的还是意外关闭的。

