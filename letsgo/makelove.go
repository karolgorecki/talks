package main

import "fmt"

func main() {
	ugly, beauty := getGirls()

	makeLove(beauty)

}

func getGirls() (ladyOne, ladyTwo string) {
	return "Miley Cyrus", "Beyonce"
}

func makeLove(lady string) {
	fmt.Println("Kissing", lady, "( ͡° ͜ʖ ͡°) ♥")
}
