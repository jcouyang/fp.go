package chans

import (
	"fmt"
	"time"
)
func strLen(a1 string) int {
		return len(a1)
	}
func ExampleMap() {

	// Sized channel
	wordChan := make(chan string, 3)
	wordLenChan := Map(strLen)(wordChan)

	wordChan <- "cat"
	wordChan <- "catfish"
	wordChan <- "caterpillar"
	
	fmt.Println(
		<-wordLenChan,
		<-wordLenChan,
		<-wordLenChan,
	)

	// Ping pong
	ping := make(chan string)
	pong := Map(func(i string) string {return "pong"})(ping)

	go func(){
		for i:=0;i<10;i++ {
			ping <- "ping"
		}
	}()

	go func() {
		for i:=0;i<10;i++ {
			fmt.Println(<-pong)
		}
	}()

	time.Sleep(200 * time.Millisecond)

	// Output: 3 7 11
	// pong
	// pong
	// pong
	// pong
	// pong
	// pong
	// pong
	// pong
	// pong
	// pong
}
