package main

import (
	"fmt"
	"strings"
)

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string
	
	fmt.Println("Welcome to", conferenceName, "booking application")
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets you want to book:")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry, we only have %d tickets remaining.\n", remainingTickets)
		}
		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)		
		
		fmt.Printf("Thanks you for booking %v tickets, You will receive mail shortly!\n", userTickets);
		
		firstnames := []string{}
		for _,booking := range bookings {
			var names = strings.Fields(booking)
			firstnames = append(firstnames, names[0])
		}
	}
}