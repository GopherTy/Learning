package rw

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
