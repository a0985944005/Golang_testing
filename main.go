package main

import (
	"fmt"
)

func main() {
	var studentAmount int = 50
	for number := 1; number <= studentAmount; number++ {
		fmt.Println(FizzBuzz(number))
	}
}

/*
Buzz
11
Fizz
13
14
FizzBuzz
16
*/
