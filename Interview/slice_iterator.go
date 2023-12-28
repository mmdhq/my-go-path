package main

//来源：程序员在囧途
//优化点：使用了泛型

import "fmt"

// 总结，对于结构体和函数来说我们需要指定[T any],
// 对于方法来说，只需要在前面写[T]
type Iterator[T any] struct {
	data  []T
	index int
}

func newIterator[T any](data []T) *Iterator[T] {
	return &Iterator[T]{
		data: data,
	}
}

func (it *Iterator[T]) HasNext() bool {
	return it.index < len(it.data)
}

func (it *Iterator[T]) Next() T {
	ret := it.data[it.index]
	it.index++
	return ret
}

func main() {
	//data := []int{1, 2, 3, 4, 5}
	data := []string{"a", "b", "c", "d", "e"}
	it := newIterator(data)
	for it.HasNext() {
		fmt.Println(it.Next())
	}

}
