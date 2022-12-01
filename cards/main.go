package main

import "fmt"

func main() {
	cards := []string{"Even Newer String", "Hell Yah", newCard()}
	cards = append(cards, "new card")
	fmt.Println(cards)
}
func newCard () string {
	return "New String"
}