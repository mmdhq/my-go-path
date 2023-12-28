package main

//来源：程序员在囧途
//优化点：使用了泛型

import (
	"errors"
	"fmt"
)

// 总结，对于结构体和函数来说我们需要指定[T any],
// 对于方法来说，只需要在前面写[T]
type Iterator[T any] struct {
	data  []T
	index int
}

// 用于初始化值
func newIterator[T any](data []T) *Iterator[T] {
	return &Iterator[T]{
		data: data,
	}
}

// 验证是否有下一位，判断下一位，下标不能有len(it.data)
func (it *Iterator[T]) HasNext() bool {
	return it.index < len(it.data)-1
}

// 这里Next函数里必须把HasNext加上,否则它不能单独调用
func (it *Iterator[T]) Next() (T, error) {
	if !it.HasNext() {
		//我之前返回一个空，不对应该返回当前值
		ret := it.data[it.index]
		return ret, errors.New("final of slice")
	}
	//我这里逻辑开始有问题，他应该是返回当前并前往下一个
	ret := it.data[it.index]
	it.index++
	return ret, nil
}

// 重置迭代器的索引
func (it *Iterator[T]) Reset() {
	it.index = 0
}

// 查看下一个元素而不改变迭代器
func (it *Iterator[T]) Peek() (T, error) {
	if !it.HasNext() {
		var zero T
		return zero, errors.New("end of board")
	}
	return it.data[it.index+1], nil
}

// 查看是否有前一个函数
func (it *Iterator[T]) HasPrevious() bool {
	return it.index > 0
}

// 向前迭代
func (it *Iterator[T]) Previous() (T, error) {
	if !it.HasPrevious() {
		value := it.data[it.index]
		return value, errors.New("begin of board")
	}
	value := it.data[it.index]
	it.index--
	return value, nil
}

func main() {
	//data := []int{1, 2, 3, 4, 5}
	data := []string{"a", "b", "c"}
	it := newIterator(data)
	//实现有下一个的话，输出这个，并前往下一位
	for it.HasNext() {
		value, _ := it.Next()
		fmt.Println(value)
	}

}
