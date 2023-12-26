package main

import (
	"errors"
	"fmt"
)

func Add[T any](array []T, value T) []T {
	array = append(array, value)
	return array
}

func Insert[T any](array []T, index int, value T) ([]T, error) {
	if index < 0 || index > len(array) {
		//这里的errors不需要事先存储到一个变量中
		return array, errors.New("index out of range")
	}
	//预分配足够的容量
	newArray := make([]T, 0, len(array)+1)
	newArray = append(array[:index], value)
	//append这里可以使用可变参数
	newArray = append(newArray, array[index:]...)
	return newArray, nil
}

func Delete[T any](array []T, index int) ([]T, error) {
	if index < 0 || index >= len(array) {
		return array, errors.New("index out of range")
	}
	newArray := append(array[:index], array[index+1:]...)
	return newArray, nil
}

func main() {
	s := make([]int, 0, 4)
	fmt.Printf("s:%+v,len %d,cap %d\n", s, len(s), cap(s))
	s = Add(s, 1)
	s = Add(s, 2)
	s = Add(s, 84)
	fmt.Printf("s:%+v,len %d,cap %d\n", s, len(s), cap(s))
	s, err := Insert(s, 3, 7)
	s, err = Insert(s, 3, 8)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("s:%+v,len %d,cap %d\n", s, len(s), cap(s))
	s, _ = Delete(s, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("s:%+v,len %d,cap %d\n", s, len(s), cap(s))
}
