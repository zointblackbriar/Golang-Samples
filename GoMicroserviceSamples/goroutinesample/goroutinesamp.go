package goroutinesample

import (
	"fmt"
	"time"
)

func String_func(sample_string string) {
	for i := 0; i < 3; i++ {
		fmt.Println(sample_string, ":", i)
	}
}

func Goroutinemain() {
	String_func("direct")

	go String_func("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("it has been done")
}
