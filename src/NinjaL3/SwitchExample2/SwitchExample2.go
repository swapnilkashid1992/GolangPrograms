package main

import "fmt"

func main() {
	favSport := "cricket"
	switch favSport {
	case "FootBall":
		fmt.Println("Football")
	case "cricket":
		fmt.Println("cricket")
	default:
		fmt.Println("No Sport")
	}
}
