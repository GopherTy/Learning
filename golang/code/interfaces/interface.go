package interfaces

import (
	"fmt"
)

// 每个 interface {} 变量在内存中占据两个字长：一个用来存储它包含的类型
// 另一个用来存储它包含的数据或者指向数据的指针。

// Test1 practice 1
// The way to Go 接口练习
func Test1() {
	// 测试接口类型变量在内存中的存储接口，
	// 1.接收的对应实例的值。
	// 2.自动生成对应实例方法表的指针。
	var a Nameable
	fmt.Println(a)
	a = &Name{1, 2, 3}
	name := a.Name("aa")
	age := a.Age(1)
	fmt.Printf("a:%v name: %v age: %v\n", a, name, age)

	// 类型实现接口
	var o valueable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
}

// Test2 practice 2
// 类型断言 判断接口是什么类型
// 接口就是一个动态类型，即在运行时才检查它的实际类型。
func Test2() {
	var value valueable
	value = new(car)

	if t, ok := value.(car); ok {
		c := t.getValue()
		fmt.Println(c)
	}
}

// Test3 practice 3
// 类型判断 type-switch
// type-switch 不允许有 fallthrough
func Test3() {
	var v valueable
	switch t := v.(type) {
	case car:
		fmt.Println("car")
	case stockPosition:
		fmt.Println(" other ")
	default:
		fmt.Printf("%T\n", t)
	}
}

// Test4 类型断言中的特例，判断该值的类型是否实现了接口。
func Test4() {
	// type Stringer interface {
	// 	String() string
	// }

	// if sv, ok := v.(Stringer); ok {
	// 	fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	// }
}

// Test5 作用于接口变量上的方法和作用于变量上的方法是不同的。
// 接口变量中存储的值是不可以寻址的。
func Test5() {
	var lst List
	// CountInto(lst, 0, 10) error

	if LongEnough(lst) {
		fmt.Println("---------")
	}

	plst := new(List)
	CountInto(plst, 1, 10) // valid : identical (相同的) receiver type
	if LongEnough(plst) {  // *List 指针会被自动解引用
		fmt.Println("sss")
	}
}

// Test6 sorter 接口
func Test6() {
	data := []int{12, 15, 20, 3, 1, 8}
	fmt.Println(data)
	a := IntArray(data)
	Sort(a)

	fmt.Println(data)
}

//Test 1

// Nameable name interface
type Nameable interface {
	Name(name string) string
	Age(age int) int
}

// Name 实现 Nameable 接口看下接口类型的变量在内存中的存储
// 是否是按照对应类型的值和自动反射出来的方法
type Name []int

// Name Nameable function
func (n Name) Name(name string) string {

	return name
}

// Age Nameable function
func (n Name) Age(age int) int {
	return age
}

// Test2

type stockPosition struct {
	tricker    string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valueable interface {
	getValue() float32
}

func showValue(asset valueable) {
	fmt.Printf("value of asset is %f \n", asset.getValue())
}

// Test5

// List .
type List []int

// Len .
func (l List) Len() int {
	return len(l)
}

// Append .
func (l *List) Append(val int) {
	*l = append(*l, val)
}

// Appender .
type Appender interface {
	Append(int)
}

// CountInto .
func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

// Lener .
type Lener interface {
	Len() int
}

// LongEnough .
func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}
