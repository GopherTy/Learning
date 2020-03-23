package function

import (
	"fmt"
	"time"
)

// 函数

// Test1  递归
func Test1() {
	r := fibonacci1(4)
	fmt.Println(r)
}

// 斐波那契数列，可以返回位置。
func fibonacci(n int) (res, position int) {

	position = n
	// 递归栈溢出
	if n <= 1 {
		res = 1
	} else {
		x, _ := fibonacci(n - 1)
		y, _ := fibonacci(n - 2)

		res = x + y
	}

	return
}

// 原始写法，从第零个位置开始
func fibonacci1(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		// 闭包的写法
		func() int {
			a, b := 1, 1 // 第 0，1个位置的数
			for i := 2; i <= n; i++ {
				res = a + b
				a = b
				b = res
			}
			return res
		}()
	}

	return
}

// Test2 回调 。
// 函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调。
func Test2() {
	callback(3, add)
}

func add(a, b int) {
	fmt.Println(a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

// Test3 闭包
// 匿名函数同样被称之为闭包（函数式语言的术语）：它们被允许调用定义在其它环境下的变量。
// 闭包可使得某个函数捕捉到一些外部状态，例如：
// 函数被创建时的状态。另一种表示方式为：一个闭包继承了函数所声明时的作用域。
// 这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁
//  判断一段代码的执行时间 用 time 包
func Test3() {
	a := func(a, b int) {
		fmt.Println(a + b)
	}
	a(1, 2) // 匿名函数，a为变量存放的是匿名函数的地址，此时变量的签名是能够找到的。

	fmt.Println(f()) // 结果是2，因为 ret++ 是在执行 return 1 语句后发生的。

	// 三次调用函数 f 的过程中函数 Adder() 中变量 delta 的值分别为：1、20 和 300。
	// 我们可以看到，在多次调用中，变量 x 的值是被保留的，即 0 + 1 = 1，然后 1 + 20 = 21，最后 21 + 300 = 321
	// ：闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
	// 可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数，是函数式语言的特点。
	// 闭包在 Go 语言中非常常见，常用于 goroutine 和管道操作

	ad := adder()
	fmt.Println(ad(1))
	fmt.Println(ad(20))
	fmt.Println(ad(300))
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func adder() func(int) int {
	var x int
	return func(delta int) int {
		fmt.Println(delta)
		x += delta
		return x
	}
}

// Test4 使用内存缓存
func Test4() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci2(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

//LIM 限制
const LIM = 41

var fibs [LIM]uint64

// 使用内存缓存提升效率
func fibonacci2(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs[n] = res
	return
}

// Test5 变长参数 在传递过程中应该传递 切片的 slice... 形式。
//  实际上变长参数的形式是 []type 形式
func Test5() {
	args([]int{1, 2, 3}...)
}

func args(arg ...int) {
	fmt.Println(arg) // var arg []int 变长参数的形式
}
