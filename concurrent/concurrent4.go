package main
import (
	"fmt"
)
func main() {
	ch := make(chan int,1)
	for  {
		select {
			case ch<-1 :
			case ch<-0 :
		}
		i := <-ch
		fmt.Println("Value：", i)
	}
}