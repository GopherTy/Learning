package ref

import (
	"fmt"
	"reflect"
)

// 反射包 reflect
// 反射是用程序检查其所拥有的结构，尤其是类型的一种能力；这是元编程的一种形式。
// 反射可以在运行时检查类型和变量，例如它的大小、方法和 动态 的调用这些方法

// Test1 方法和类型的反射
// 反射是通过检查一个接口的值，变量首先被转换成空接口。
func Test1() {
	var a int = 1

	aType := reflect.TypeOf(a)
	aValue := reflect.ValueOf(a)

	fmt.Printf("type: %v  value: %v \n", aType, aValue)
	fmt.Println("type:", aValue.Type())
	fmt.Println("kind:", aValue.Kind())
	fmt.Println("value:", aValue.Int())
	fmt.Println(aValue.Interface())
	fmt.Printf("value is %v\n", aValue.Interface())

	y := aValue.Interface().(int)
	fmt.Println(y)

}

// Test2 接口与动态类型
func Test2() {
	b := new(Bird)
	DuckDance(b)
}

// IDuck .
type IDuck interface {
	Quack()
	Walk()
}

// DuckDance .
func DuckDance(duck IDuck) {
	for i := 0; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

// Bird .
type Bird struct {
}

// Quack .
func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

// Walk .
func (b *Bird) Walk() {
	fmt.Println("I am walking!")
}
