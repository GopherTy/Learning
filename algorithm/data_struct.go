package algorithm

import "errors"

// Stack stack operate
type Stack []interface{}

// Len  stack length
func (s Stack) Len() int {
	return len(s)
}

// Cap stack cap
func (s Stack) Cap() int {
	return cap(s)
}

// IsEmpty is stack empty?
func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

// Push add data to stack
func (s *Stack) Push(e interface{}) {
	*s = append(*s, e)
}

// Top get stack top data
func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s[len(s)-1], nil
}

// Pop pop stack data
func (s *Stack) Pop() (interface{}, error) {
	stk := *s
	if len(stk) == 0 {
		return nil, errors.New("stack is empty")
	}
	top := stk[len(stk)-1]
	*s = stk[:len(stk)-1]
	return top, nil
}

// 除了下面的队列实现方式还可以直接使用 slice 和标准库中的 list.New() 双向链表充当队列。
// 但是 slice 这种实现方式会造成内存泄漏。
// 此队列实现方式比较浪费内存空间，可以使用循环队列进行优化处理。

// Queue .
type Queue struct {
	Element []interface{}
	Index   int
}

// EnQueue enqueue
func (q *Queue) EnQueue(val interface{}) {
	if q.Element == nil {
		q.Element = make([]interface{}, 0)
	}

	q.Element = append(q.Element, val)
}

// Front return first Element from queue
func (q *Queue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.Element[q.Index]
}

// DeQueue dequeue
func (q *Queue) DeQueue() (ok bool) {
	if q.IsEmpty() {
		return false
	}

	q.Index++
	return true
}

// IsEmpty .
func (q *Queue) IsEmpty() bool {
	if q.Element == nil || q.Index >= len(q.Element) {
		return true
	}
	return false
}

// CycleQueue 循环队列
// 循环队列：入栈 tail 游标移动，出栈 head 游标移动。
// 栈满：(tail + 1) % len(element) == head (此处踩了坑)
// 栈空：head,tail  = -1,-1
type CycleQueue struct {
	element    []interface{}
	size       int // 循环队列大小
	head, tail int // 循环队列指针，初始值为 -1
}

// Constructor 构造循环队列
func Constructor(size int) CycleQueue {
	return CycleQueue{
		element: make([]interface{}, size),
		head:    -1,
		tail:    -1,
	}
}

// EnQueue 入队
func (c *CycleQueue) EnQueue(value interface{}) bool {
	// 判断队列是否满
	if c.IsFull() {
		return false
	}

	// 判断是否是第一次入队
	if c.head == -1 {
		c.head++
	}
	// 不管是否为第一次入队将游标加1，因为下标从 -1 开始
	c.tail++
	// 判断是否超过了下标长度,超过了游标归零。
	if c.tail == len(c.element) {
		c.tail = 0
	}

	c.element[c.tail] = value
	return true
}

// DeQueue 出队
func (c *CycleQueue) DeQueue() bool {
	if c.IsEmpty() {
		return false
	}

	// 先删除队列中元素
	c.element[c.head] = nil
	// 当两个游标指向同一块位置时，游标归为 -1。
	if c.head == c.tail {
		c.head = -1
		c.tail = -1
		return true // 出队成功
	}

	// 移动游标
	c.head++
	// 判断是否超过了下标长度,超过了游标归零。
	if c.head == len(c.element) {
		c.head = 0
	}
	return true
}

// Front 返回队头的元素
func (c *CycleQueue) Front() interface{} {
	if c.IsEmpty() {
		return nil
	}
	return c.element[c.head]
}

// End 返回队尾的元素
func (c *CycleQueue) End() interface{} {
	if c.IsEmpty() {
		return nil
	}
	return c.element[c.tail]
}

// IsEmpty 队列是否为空
func (c *CycleQueue) IsEmpty() bool {
	if c.head == -1 && c.tail == -1 {
		return true
	}
	return false
}

// IsFull 队列是否已满
func (c *CycleQueue) IsFull() bool {
	if (c.tail+1)%len(c.element) == c.head {
		return true
	}
	return false
}
