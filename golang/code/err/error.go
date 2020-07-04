package err

import "fmt"

// defer-panic-recover recover 必须在 defer 修饰的函数中使用。
func badcall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panicing", e)
		}
	}()

	badcall()
	fmt.Println("after bad call")
}
