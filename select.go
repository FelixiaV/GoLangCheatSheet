package main

import (
	"fmt"
	"time"
)

func main() {


	c1 := make(chan string)
	c2 := make(chan string) 


	go func(){
		for {
			c2 <- "Every Two Second"
			time.Sleep(time.Second * 2)
		}
	}()


	go func(){
		for{
			c1 <- "Every 500 mas"
			time.Sleep(time.Millisecond*500)
			}
		}()
	


	for {
		select{
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}


	}
}

