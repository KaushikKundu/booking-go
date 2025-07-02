package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = []User{}

type User struct {
	firstName string
	lastName string
	email string
	ticketsBooked uint
}

func main() {

	greetUsers()
	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidFirstName, isValidLastName, isValidEmail, isTicketValid := validateUserInput(firstName, lastName, email, userTickets)
		if !isValidFirstName || !isValidLastName {
			fmt.Println("Names invalid  or too short.")
			continue
		}
		if !isTicketValid {
			fmt.Println("We don't have that much tickets. tickets left:",remainingTickets)
			continue
		}
		if !isValidEmail {
			fmt.Println("Invalid email")
			continue
		}
		bookTickets(userTickets,firstName,lastName,email)
		fmt.Println(bookings)

	}

}
func bookTickets(userTickets uint, firstName, lastName, email string) {
	remainingTickets -= userTickets
	var user = User {
		firstName: firstName,
		lastName: lastName,
		email: email,
		ticketsBooked:userTickets,
	}
	bookings = append(bookings,user)
	fmt.Println("tickets left:", remainingTickets)
}
func printNames() {
	for _, booking := range bookings {
		var firstName = booking.firstName
		fmt.Println(firstName)
	}
}
func validateUserInput(firstName, lastName, email string, tickets uint) (bool, bool, bool, bool) {
	isValidFirstName := len(strings.TrimSpace(firstName)) > 0
	isValidLastName := len(strings.TrimSpace(lastName)) > 0 
	isValidEmail := len(strings.TrimSpace(email)) > 0 && strings.Contains(email,"@")
	isTicketValid := tickets > 0 && tickets <= remainingTickets
	return isValidFirstName, isValidLastName, isValidEmail, isTicketValid
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func greetUsers() {
	var conferenceName = "Go Conference"
	fmt.Println("Welcome to", conferenceName, "booking application")
}
