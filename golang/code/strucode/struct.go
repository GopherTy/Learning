package strucode

import (
	"fmt"
	"reflect"
)

/**

Go 通过类型别名（alias types）和结构体的形式支持用户自定义类型，或者叫定制类型。
一个带属性的结构体试图表示一个现实世界中的实体。结构体是复合类型（composite types），
当需要定义一个类型，它由一系列属性组成，每个属性都有自己的类型和值的时候，就应该使用结构体，它把数据聚集在一起。
然后可以访问这些数据，就好像它是一个独立实体的一部分。结构体也是值类型，因此可以通过 new 函数来创建。

组成结构体类型的那些数据称为 字段（fields）。每个字段都有一个类型和一个名字；在一个结构体中，字段名字必须是唯一的。

结构体的概念在软件工程上旧的术语叫 ADT（抽象数据类型：Abstract Data Type），
在一些老的编程语言中叫 记录（Record），比如 Cobol，在 C 家族的编程语言中它也存在，
并且名字也是 struct，在面向对象的编程语言中，跟一个无方法的轻量级类一样。
不过因为 Go 语言中没有类的概念，因此在 Go 中结构体有着更为重要的地位。
*/

// Test1 结构体中的一些细节，需要注意的地方。
//声明 var t T 也会给 t 分配内存， 并零值化内存，但是这个时候 t 是类型T。
//在这两种方式中，t 通常被称做类型 T 的一个实例（instance）或对象（object）。
func Test1() {
	a := new(testStruct)
	var b *testStruct
	var c testStruct
	// fmt.Println(b.filed1) // error b 是一个空指针，会出问题。
	// c 是一个实例
	fmt.Println(a, b, c.filed1)
}

type testStruct struct {
	filed1 int
	filed2 string
	filed3 float64
}

// Test2 带标签的结构体方法。
// 通常要是相比面向对象可以使用包中结构体不导出，然后通过 New 方法（工厂方法）返回一个结构体。
func Test2() {
	a := tagType{}
	fmt.Println(reflect.TypeOf(a).Field(0).Tag)
}

// tagType .
type tagType struct { // tags
	Field1 bool `json:"An important answer"`
	// field2 string `The name of the thing`
	// field3 int    `How much there are`
}

// Test3 匿名字段和嵌套结构体
//结构体可以包含一个或多个 匿名（或内嵌）字段，即这些字段没有显式的名字，
// 只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体。
// 在一个结构体中对于每一种数据类型只能有一个匿名字段。
func Test3() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5

	// 以下为匿名字段
	outer.int = 1

	outer.innerS.in1 = 2
	outer.in2 = 3

	fmt.Println(outer)

	// 结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{1, 2}}
	fmt.Println(outer2)

	test := test{3.2, 1, "hahaha"}
	fmt.Println(test)
}

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field  outerS 内嵌了 innerS
}

type test struct {
	price float32
	int
	string
}

/**
内嵌中出现的命名冲突：

	当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？
		1. 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
		2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。
		没有办法来解决这种问题引起的二义性，必须由程序员自己修正。
*/
