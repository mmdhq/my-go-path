package main

import (
	"fmt"
	"sync"
)

//多协程执行后使用channael收集结果,来源程序员在囧途

func main2() {
	megCh := make(chan any)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			megCh <- i + 23
		}(i)
	}
	go func() {
		wg.Wait()
		defer close(megCh)
	}()
	for mes := range megCh {
		fmt.Println("协程中的消息", mes)
	}

}
