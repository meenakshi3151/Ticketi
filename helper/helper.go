package helper
import (
"fmt"
"strings"
)
//we can start with cap letter to directly export the function
// Validate the user's name
func ValidateName(firstName string, lastName string) bool {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	if !isValidName {
		fmt.Println("Please enter a valid name with at least 2 characters.")
		return false
	}
	return true
}

// Validate the user's email
func ValidateEmail(email string) bool {
	isValidEmail := strings.Contains(email, "@gmail.com")
	if !isValidEmail {
		fmt.Println("Please enter a valid Gmail address.")
		return false
	}
	return true
}

// Validate the number of tickets to book
func ValidateTicketsToBook(ticketsToBook uint) bool {
	if ticketsToBook <= 0 {
		fmt.Println("Please enter a positive number of tickets.")
		return false
	}
	return true
}

// Check if enough tickets are available
func CheckAvailabilityOfTickets(ticketsToBook uint, remainingTickets uint) bool {
	if ticketsToBook > remainingTickets {
		fmt.Printf("Sorry, we only have %v tickets left. Please try again.\n", remainingTickets)
		return false
	}
	return true
}