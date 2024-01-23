package main

import (
	"LeetCodeGo/TopInterview150/medium"
	"log"
)

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7}
	medium.Rotate(test, 3)
	log.Println(test)
}
