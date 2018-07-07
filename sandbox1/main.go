package main

import (
	"fmt"
	"time"
)

func producer(min, max, step int, channel chan string) {
	for i := min; i < max; i += step {
		channel <- fmt.Sprintf("producer[%d,%d,%d] = %d", min, max, step, i)
		time.Sleep(500 * time.Millisecond)
	}
	close(channel)
	fmt.Println("*** producer exited!")
}

func consumer(channel chan string) {
	val, ok := <-channel
	for ok {
		fmt.Printf("consumer read \"%s\"\n", val)
		time.Sleep(30 * time.Millisecond)
		val, ok = <-channel
	}
	fmt.Println("*** consumer exited!")
}

func selector(channel chan string) {
	for {
		select {
		case v, ok := <-channel:
			if !ok {
				fmt.Println("*** selector exited!")
				return
			}
			fmt.Printf("selector received \"%s\"\n", v)
		}
	}
}

type Vertex struct {
	x, y int
}

func main() {
	// a1 := data.Create(123, "King St.", "Arlington", "VA", 22204)
	// fmt.Println(a1)
	// a1 = data.Create(321, "Jefferson Rd.", "Charlottesvill", "VA", 12345)
	// fmt.Println(a1)

	// c := make(chan string)
	// go producer(4, 40, 2, c)
	// go consumer(c)
	// selector(c)

	var v1 Vertex
	var v2 *Vertex

	v1 = Vertex{1, 2}
	v2 = &v1

	fmt.Printf("v1 = %v\n", v1)
	fmt.Printf("v2 = %v\n", v2)
	fmt.Println("---- modify ----")
	v2.x = 4
	fmt.Printf("v1 = %v\n", v1)
	fmt.Printf("v2 = %v\n", v2)
}
