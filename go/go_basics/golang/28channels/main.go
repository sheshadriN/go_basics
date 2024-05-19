package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	fmt.Println("Channels in golang-  LearnCodeOnline.in")

// 	myCh := make(chan int, 2)
// 	wg := &sync.WaitGroup{}

// 	// fmt.Println(<-myCh)
// 	// myCh <- 5
// 	wg.Add(2)
// 	// R ONLY
// 	go func(ch <-chan int, wg *sync.WaitGroup) {

// 		val, isChanelOpen := <-myCh

// 		fmt.Println(isChanelOpen)
// 		fmt.Println(val)

// 		//fmt.Println(<-myCh)

// 		wg.Done()
// 	}(myCh, wg)
// 	// send ONLY
// 	go func(ch chan<- int, wg *sync.WaitGroup) {
// 		myCh <- 0
// 		close(myCh)
// 		// myCh <- 6
// 		wg.Done()
// 	}(myCh, wg)

// 	wg.Wait()
// }

func main() {
	channel := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go myfunc(channel, wg)  //read
	go myfunc2(channel, wg) // write
	// channel <- 11
	// defer close(channel)
	defer wg.Wait()

}
func myfunc(channel chan int, wg *sync.WaitGroup) {
	for {
		value, ok := <-channel
		if !ok {
			break
		}
		fmt.Println(value)
	}
	defer wg.Done()
}
func myfunc2(channel chan int, wg *sync.WaitGroup) {
	for i := 0; i < 21; i++ {
		channel <- i
	}
	defer wg.Done()
	defer close(channel)
	
}
