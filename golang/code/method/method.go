package method

import (
	"fmt"
	"strconv"
)

// 方法
// Go 方法是作用在接收者（receiver）上的一个函数，
// 接收者是某种类型的变量。因此方法是一种特殊类型的函数。

// 1. 在 Go 中，接受者（receiver）不能是接口类型，因为接口是一个抽象定义，
// 方法是具体实现。
// 2. 接受者不能是一个指针类型，但是它可以是其他允许类型的指针。
// 3. 类型和它的方法可以在不同源文件中,它们必须是同一个包的。
// 4.类型 T（或 *T）上的所有方法的集合叫做类型 T（或 *T）的方法集（method set）。

// 注意：
// 如果 recv 是 receiver 的实例，Method1 是它的方法名，
// 那么方法调用遵循传统的 object.name 选择器符号：recv.Method1()。

// 如果 recv 是一个指针，Go 会自动解引用. 如果方法不需要使用 recv 的值，可以用 _ 替换它

// Test1 .
func Test1() {
	e := employee{salary: 3.2}
	rs := e.giveRaise(5000)
	fmt.Println(rs)
}

type employee struct {
	salary float32
}

func (e employee) giveRaise(price float32) float32 {
	return price * (1 + price)
}

// Test2 .  具体只跟方法的接收者相关，如果接受者为指针，调用者(t) 是值类型
// Go 会自动转换 (&t).method()
// 可以通过指针接受者来修改一个结构体的外包不可见字段。
func Test2() {
	var b1 B
	b1.change()
	fmt.Println(b1)
	fmt.Println(b1.write())
	fmt.Println(b1)
}

// B .
type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 2
	fmt.Println(b)
}

func (b B) write() string {
	b.thing = 3
	return fmt.Sprint(b)
}

// Test3  类型 String() 方法和格式化描述
// 当为类型定义了 String 方法时，调用 print 方法会自动调用
// 该类型的 String 方法
func Test3() {
	two1 := &twoInts{1, 2}
	fmt.Println(two1)
	fmt.Printf("two is %v \n", two1)
}

type twoInts struct {
	a int
	b int
}

func (tn *twoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

// Test4 创建一个 employee 实例显示出它的ID
func Test4() {
	e := Employee{Person{"t", "yx", Base{1}}, 1.0}
	id := e.ID()
	fmt.Println(id)
}

// Base .
type Base struct {
	id int
}

// SetID .
func (b *Base) SetID(id int) {
	b.id = id
}

// ID .
func (b Base) ID() int {
	return b.id
}

// Person .
type Person struct {
	FirstName string
	LastName  string
	Base
}

// Employee .
type Employee struct {
	Person
	salary float64
}
