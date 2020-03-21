package slicemd

import (
	"bytes"
	"fmt"
)

// 切片和数组

// Test1 一些基本操作
// 当声明数组时所有的元素都会被自动初始化为默认值 0
// 数组是值类型，可以通过 new 来创建。
func Test1() {
	a := make([]interface{}, 10)
	a[0] = "string"
	for key, v := range a {
		fmt.Println(key, v)
	}
	fmt.Println(a[1])

	var arr1 = new([5]int)
	// var arr2 = [5]int
	arr2 := *arr1
	arr2[2] = 100
	fmt.Println(arr1, arr2)

}

// Test2 切片的一些常规操作。
// 切片（slice）是对数组一个连续片段的引用（该数组我们称之为相关数组，通常是匿名的）
// ，所以切片是一个引用类型
// 注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针!!
func Test2() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	a := []int{1, 2, 3, 4, 5, 6}
	b := a[2:5]
	c := b[0:1]
	fmt.Println(len(b), cap(b), cap(c)) // 切片容量只跟相关数组有关。

	// 对 []byte 切片的处理
	var buffer bytes.Buffer
	for i := 0; i < 3; i++ {
		_, err := buffer.WriteString("asda")
		if err != nil {
			break
		}
	}
}

// Append 练习，此时可以看出GO中只有值传递。此时的 append 函数（接收引用类型）
// 但却是改变的 slice 的一个副本（拷贝）所以不会改变实参。
func Append(slice, data []byte) (b []byte) {
	// slice cap
	sliceCap := cap(slice)

	// len slice data
	sliceLen := len(slice)
	dataLen := len(data)

	if sliceCap < sliceLen+dataLen {
		b = make([]byte, sliceCap+dataLen)
	} else {
		b = make([]byte, sliceCap)
	}

	for i := 0; i < dataLen; i++ {
		slice = append(slice, data[i])
	}

	b = slice
	return
}

// Test3 关于切片的扩容和拷贝
func Test3() {
	a := []byte{97, 98}
	b := byte('c')
	c := AppendByte(a, b)
	fmt.Println(c)
}

// AppendByte .
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

// Test4 字符串、数组、切片
// 字符串本质上是一个字节数组
func Test4() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c) // s2 == "cello"
	fmt.Println(s2)
}
