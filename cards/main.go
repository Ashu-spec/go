package main

import "fmt"

func main() {
	cards := []string{"Even Newer String", "Hell Yah", newCard()}
	fmt.Println(cards)
}
func newCard () string {
	return "New String"
}