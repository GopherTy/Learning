package rw

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

// 接收用户输入

// Input .
func Input() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}

// BufIORead .
func BufIORead() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	switch input {
	case "Philip\r\n":
		fmt.Println("Welcome Philip!")
	case "Chris\r\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\r\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch input {
	case "Philip\r\n":
		fallthrough
	case "Ivo\r\n":
		fallthrough
	case "Chris\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\r\n", "Ivo\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}

// 文件读写
func fileRW() {
	// 将整个文件的内容读取到一个字符串里
	buf, err := ioutil.ReadFile("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))

	// 带缓冲的读取
	b := make([]byte, 1024)
	var reader bufio.Reader
	reader.Read(b)

	// 按列读取文件
	// fmt.Fscanln(inputFIle,&v1)

	// 读取压缩文件
	// gzip.NewReader(inputFIle)

	// 写文件
	// outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	// if outputError != nil {
	// 	fmt.Printf("An error occurred with file opening or creation\n")
	// 	return
	// }
	// defer outputFile.Close()

	// outputWriter := bufio.NewWriter(outputFile)
	// outputString := "hello world!\n"

	// for i:=0; i<10; i++ {
	// 	outputWriter.WriteString(outputString)
	// }
	// outputWriter.Flush()

	// 打开文件，并读取内容
	inputFIle, inputError := os.Open("a.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFIle.Close()
	inputReader := bufio.NewReader(inputFIle)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}

}

type page struct {
	Title string
	Body  []byte
}

func (p *page) save() (err error) {
	f, err := os.OpenFile(p.Title, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	w := bufio.NewWriter(f)
	_, err = w.Write(p.Body)
	if err != nil {
		return
	}
	w.Flush()
	return
}

func (p *page) load(title string) (err error) {
	p.Title = title
	buf, err := ioutil.ReadFile(title)
	if err != nil {
		return
	}
	p.Body = buf
	return
}

// 文件拷贝
func filecopy(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// bufio读取
func bufRead(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *a {
			fmt.Fprintf(os.Stdout, "%d %s", n, buf)
			n++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}

	}
}

var n int = 0
var a = flag.Bool("n", false, "output line number")

func bufReadTest() {
	flag.Parse()
	if flag.NArg() == 0 {
		bufRead(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		bufRead(bufio.NewReader(f))
	}
}

func cat(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}

// 数据结构 --> 指定格式 = 序列化 或 编码（传输之前）
// 指定格式 --> 数据格式 = 反序列化 或 解码（传输之后）
// 编码也是一样的，只是输出一个数据流（实现了 io.Writer 接口）
// 解码是从一个数据流（实现了 io.Reader）输出到一个数据结构。
func jsonRW() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}

// Address .
type Address struct {
	Type    string
	City    string
	Country string
}

// VCard .
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

// json 包使用 map[string]interface{} 和 []interface{} 储存任意的 JSON 对象和数组；
// 其可以被反序列化为任何的 JSON blob 存储到接口值中。
func jsonUnmashalAny() {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f) // json 可以被反序列化为任何的 json blob 存储到接口值中。
	if err != nil {
		panic(err)
	}
	var m map[string]interface{}
	var ok bool
	if m, ok = f.(map[string]interface{}); !ok {
		return
	}

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)

		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don’t know how to handle")
		}
	}
}

// xml 编码解码
func xmlEnDecode() {
	var t xml.Token
	var err error

	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)

	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				// ...
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
			// ...
		default:
			// ...
		}
	}
}

// gobEnDecode 编码解码
func gobEnDecode() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
}

// P .
type P struct {
	X, Y, Z int
	Name    string
}

// Q .
type Q struct {
	X, Y *int32
	Name string
}
