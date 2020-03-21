package mapmd

import (
	"fmt"
	"sort"
)

// Map 练习

// map 引用类型。未初始化的 map 的值是 nil。
// 在声明的时候不需要知道 map 的长度，map 是可以动态增长的。
//数组、切片和结构体不能作为 key,但是指针和接口类型可以。
//(含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的）

// map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。
// 通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍

// Test1 不要使用 new，永远用 make 来构造 map
// 当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。
func Test1() {
	var m map[int]interface{}
	m = make(map[int]interface{})
	v, ok := m[0]
	if ok {
		fmt.Println("present", v)
	}
	fmt.Println(v, m, m[1])

	var mapList map[string]int

	var mapAssigned map[string]int

	mapList = map[string]int{"one": 1}
	mapCreated := make(map[string]float64)
	mapAssigned = mapList

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.0
	// mapAssigned["two"] = 3

	fmt.Println(&mapAssigned == &mapList) // mapAssigned 存的是 mapList 的引用

	fmt.Printf("Map literal at \"one\" is: %d\n", mapList["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapList["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapList["ten"])
}

// Test2 map的一些用法
func Test2() {
	m := make(map[int]int)

	// 用于判断键是否存在，存在则取出值，键不存在则值为空( <nil> )
	if v, ok := m[0]; ok {
		fmt.Println(v)
	}

	// 删除键 delete(map1,key1) key1 不存在不会产生错误。
	delete(m, 1)

	// for-range key ,value 仅内部可见。
	// 注意 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。
	for key, value := range m {
		fmt.Println(key, value)
	}

	// map 类型的切片分配内存

	s := make([]map[int]int, 5)

	for i := range s {
		s[i] = make(map[int]int)
		s[i][0] = 1
	}

	// 以下方式中 items2 中的 map 并没有真的得到初始化，而是得到了items2 中的拷贝。
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)

	// items2[0][1] = 3 // error map 为空
}

// Test3 map 的排序
func Test3() {
	barVal := map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}

	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}

	keys := make([]string, len(barVal))

	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}
