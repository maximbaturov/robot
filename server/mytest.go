package main

import (
	"fmt"
	"strconv"
)

func main() {

	t := 2
	channel := make(chan string, t)
	channel <- "test string zero"

	//go a(channel)
	//go a(channel)

	//fmt.Println(<-channel)

	for i := 0; i < t; i++ {
		go a(channel)
		channel <- "test string " + strconv.Itoa(i)
	}

	i := 0
	for str := range channel {
		fmt.Println(str)
		if i == t {
			//close(channel)
			break
		}
		i++
	}

	//fmt.Println(<-channel)
	close(channel)

}

func a(channel chan string) {
	fmt.Println(<-channel)
	channel <- "hello"
	//fmt.Println("1")
	//time.Sleep(1 * time.Second)
	//close(channel)
	//}
}
