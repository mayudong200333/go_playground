package main

import "fmt"

func main() {
	s := []int{}
	for i := 0; i <= 10; i++ {
		s = append(s, i)
	}
	for _, v := range s {
		if v%2 == 0 {
			fmt.Println(v, "is", "even")
		} else {
			fmt.Println(v, "is", "odd")
		}
	}
}
