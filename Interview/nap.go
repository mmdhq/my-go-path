package main

import "fmt"

func main1() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan struct{})
	defer func() {
		close(ch1)
		close(ch2)
		close(ch3)
	}()
	go func() {
		for i := 0; ; i += 2 {
			<-ch1
			fmt.Println(i)
			fmt.Println(i + 1)
			ch2 <- 1
		}
	}()
	go func() {
		for i := 97; i < 123; i++ {
			<-ch2
			fmt.Printf("%c\n", i)
			ch1 <- 1
		}
		ch3 <- struct{}{}
		return
	}()
	ch1 <- 1
	<-ch3
}
