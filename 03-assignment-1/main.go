package main

import "fmt"

func main() {
	s := make([]int, 11)
	for i := 0; i < 11; i++ {
		s[i] = i
	}

	for _, v := range s {
		if v%2 == 0 {
			fmt.Printf("%v is even\n", v)
		} else {
			fmt.Printf("%v is odd\n", v)
		}
	}

}
