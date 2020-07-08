/*
Go routines and channels are a seperate line of code execution that can be used to handle blocking code.

The purpose of a channel is to communicate between go routines.

*/

//Website Status Checker

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//*Therefore we need channel
	c := make(chan string)

	//bug: child go routine cant run when main routine exits *
	for _, link := range links {
		go checkLink(link, c)
	}

	//blocking channel
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	//1 So we'll put this in a for loop
	for i := 0; i < len(links); i++ {
		//fmt.Println(<-c)
		go checkLink(<-c, c)
	}

	//Or an infinite loop
	//2 hence for receiving messages while pinging continuously multiple times

	/*for {
		go checkLink(<-c, c)
	}*/

	//Alternate loop syntax
	for l := range c {
		go checkLink(l, c)
	}

	//Sleeping a routine
	for l := range c {
		//function literal- unnamed function (like lambda function in java)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
