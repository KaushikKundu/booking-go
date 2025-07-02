package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string,0)


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
		// printNames(bookings)

	}

}
func bookTickets(userTickets uint, firstName, lastName, email string) {
	remainingTickets -= userTickets
	var user = make(map[string]string)
	user["firstName"] = firstName
	user["lastName"] = lastName
	user["email"]= email
	user["userTickets"] = strconv.FormatUint(uint64(userTickets),10)
	fmt.Println("tickets left:", remainingTickets)
}
func printNames(bookings []string) {
	firstnames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}
}
func validateUserInput(firstName, lastName, email string, tickets uint) (bool, bool, bool, bool) {
	isValidFirstName := len(strings.TrimSpace(firstName)) > 0 && reflect.TypeOf(firstName) == reflect.TypeOf("")
	isValidLastName := len(strings.TrimSpace(lastName)) > 0 && reflect.TypeOf(lastName) == reflect.TypeOf("")
	isValidEmail := len(strings.TrimSpace(email)) > 0 && reflect.TypeOf(email) == reflect.TypeOf("")
	isTicketValid := tickets > 0 && tickets <= remainingTickets
	return isValidFirstName, isValidLastName, isValidEmail, isTicketValid
}
func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}
func greetUsers() {
	var conferenceName = "Go Conference"
	fmt.Println("Welcome to", conferenceName, "booking application")
}
