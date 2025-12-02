package main

import "fmt"

func main() {
	var leagueName = "LaLiga"
	var firstname string
	var Lastname string
	var age uint
	var Total_players uint = 24

	fmt.Println("Welcome to the player registration portal")

	fmt.Println("Your firstname is: ")
	fmt.Scan(&firstname)

	fmt.Println("Your Lastname is: ")
	fmt.Scan(&Lastname)

	fmt.Println("Your age is: ")
	fmt.Scan(&age)

	if age < 16 {
		fmt.Printf("Sorry %v %v, you are not eligible to play in %v\n", firstname, Lastname, leagueName)
	} else {
		fmt.Printf("Welcome %v %v, you are eligible to play in %v\n", firstname, Lastname, leagueName)
	}
	fmt.Printf("Thanks for registering with us %v %v\n", firstname, Lastname)
	fmt.Printf("Now there are only %v players allowed in %v\n", Total_players-1, leagueName)
}
