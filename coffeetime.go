package main

import (
	"fmt"
	"time"
)

func main() {
	for timer := 5; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Coffee Time")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}
