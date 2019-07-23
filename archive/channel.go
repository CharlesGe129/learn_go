package archive

import (
	"fmt"
	"time"
)

func main() {
	var messages chan string = make(chan string, 1)
	go func(message string) {
		//time.Sleep(time.Second)
		messages <- message
		//messages <- "123"
		//messages <- "456"
		fmt.Println("Done")
	}("Ping!")
	//messages <- "123"
	time.Sleep(time.Second)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
