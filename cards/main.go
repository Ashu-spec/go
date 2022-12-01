package main

import "fmt"

func main() {
	cards := []string{"Even Newer String", "Hell Yah", newCard()}
	cards = append(cards, "new card")

	for i, cards := range cards {
		fmt.Println(i, cards)
	}
}
func newCard () string {
	return "New String"
}