package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func printText(text string) {
	for _, i := range text {
		fmt.Print(string(i))
		time.Sleep(time.Millisecond * 40)
	}
	fmt.Print("\n")
}

func main() {
	printText("Random bald generator!")
	time.Sleep(time.Second)
	fmt.Println()
	fmt.Println()
	var bald [5]int
	var res int = 0
	for i := range bald {
		bald[i] = rand.Intn(100)
		res += bald[i]
	}
	printText("You are " + strconv.Itoa(bald[0]) + "% bald")
	printText("Your sister is " + strconv.Itoa(bald[1]) + "% bald")
	printText("Your brother is " + strconv.Itoa(bald[2]) + "% bald")
	printText("Your dog is " + strconv.Itoa(bald[3]) + "% bald")
	printText("Your cat is " + strconv.Itoa(bald[4]) + "% bald")

	printText("My rating: " + strconv.Itoa(res) + "%!")
	fmt.Println()
	fmt.Println()
}
