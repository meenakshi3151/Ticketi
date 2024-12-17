package main

import (
	"fmt"
	"strings"
)

func main() {
	// Conference details
	var conferenceName = "GoConf"
	const totalTickets = 100
	var remainingTickets uint = 100
	var bookedTickets = 0

	// Slice to store bookings dynamically
	var bookingsList []string
    
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are currently available.\n", totalTickets, remainingTickets)
	fmt.Println("Book your tickets now!")

	// Start the booking process
	for remainingTickets > 0 && len(bookingsList) < 50 {
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("Currently, %v tickets are booked, and %v tickets are remaining.\n", bookedTickets, remainingTickets)

		// Collect user details
		var firstName, lastName, email string
		var ticketsToBook uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Print("Enter the number of tickets you'd like to book: ")
		fmt.Scan(&ticketsToBook)

		//validate name
		isValidName := len(firstName) > 2 && len(lastName) > 2
		if !isValidName {
			fmt.Println("Please enter name correctly with atleast 2 characters")
			continue
		}

		//validate email address
		isValidEmail := strings.Contains(email, "@gmail.com")

		if !isValidEmail {
			fmt.Println("Please enter a valid email address")
			continue
		}
        
		//valid tickets
		isValidTickets := ticketsToBook>0
		if !isValidTickets{
			fmt.Println("Enter positive number of tickets")
		}

		// Check ticket availability
		if ticketsToBook > remainingTickets {
			fmt.Printf("Sorry, we only have %v tickets left. Please try again.\n", remainingTickets)
			
			continue
		}

		// Update ticket counts
		remainingTickets -= ticketsToBook
		bookedTickets += int(ticketsToBook)

		// Store booking details
		bookingDetails := firstName + " " + lastName
		bookingsList = append(bookingsList, bookingDetails)

		// Confirmation message
		fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n", firstName, lastName, ticketsToBook, email)
		fmt.Printf("Tickets remaining: %v\n", remainingTickets)

		// Display all users who have booked tickets
		fmt.Printf("List of all bookings: %v\n", bookingsList)

		// Extract and display first names of all users
		var firstNames []string
		for _, booking := range bookingsList {
			firstName := strings.Fields(booking)[0] // Extract first name
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("First names of all attendees: %v\n", firstNames)

		// End booking if no tickets are left
		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out! See you at the conference.")
			break
		}
	}
}
