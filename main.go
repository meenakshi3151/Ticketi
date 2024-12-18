package main

import (
	"fmt"
	"strings"
)

// Greet users and introduce the conference
func greetUsers(conferenceName string) {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
}

// Display initial ticket information
func initialTicketInfo(totalTickets int, remainingTickets uint) {
	fmt.Printf("We have a total of %v tickets, and %v tickets are currently available.\n", totalTickets, remainingTickets)
	fmt.Println("Book your tickets now!")
}

// Display the current ticket status
func displayCurrentTicketsInfo(remainingTickets uint, bookedTickets int) {
	fmt.Printf("Currently, %v tickets are booked, and %v tickets are remaining.\n", bookedTickets, remainingTickets)
}

// Collect and validate user details
func collectUserInfo(firstName *string, lastName *string, email *string, city *string, ticketsToBook *uint, remainingTickets uint) {
	fmt.Print("Enter your first name: ")
	fmt.Scan(firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(lastName)
	if !validateName(*firstName, *lastName) {
		return
	}

	fmt.Print("Enter your email address: ")
	fmt.Scan(email)
	if !validateEmail(*email) {
		return
	}

	fmt.Print("Enter the number of tickets you'd like to book: ")
	fmt.Scan(ticketsToBook)
	if !validateTicketsToBook(*ticketsToBook) || !checkAvailabilityOfTickets(*ticketsToBook, remainingTickets) {
		return
	}

	fmt.Print("Enter your city: ")
	fmt.Scan(city)
}

// Validate the user's name
func validateName(firstName string, lastName string) bool {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	if !isValidName {
		fmt.Println("Please enter a valid name with at least 2 characters.")
		return false
	}
	return true
}

// Validate the user's email
func validateEmail(email string) bool {
	isValidEmail := strings.Contains(email, "@gmail.com")
	if !isValidEmail {
		fmt.Println("Please enter a valid Gmail address.")
		return false
	}
	return true
}

// Validate the number of tickets to book
func validateTicketsToBook(ticketsToBook uint) bool {
	if ticketsToBook <= 0 {
		fmt.Println("Please enter a positive number of tickets.")
		return false
	}
	return true
}

// Check if enough tickets are available
func checkAvailabilityOfTickets(ticketsToBook uint, remainingTickets uint) bool {
	if ticketsToBook > remainingTickets {
		fmt.Printf("Sorry, we only have %v tickets left. Please try again.\n", remainingTickets)
		return false
	}
	return true
}

// Update ticket counts after a successful booking
func updateTicketCounts(remainingTickets *uint, bookedTickets *int, ticketsToBook uint) {
	*remainingTickets -= ticketsToBook
	*bookedTickets += int(ticketsToBook)
}

// Store booking details in the bookings list
func storeBookingDetails(firstName, lastName string, bookingsList *[]string) {
	bookingDetails := firstName + " " + lastName
	*bookingsList = append(*bookingsList, bookingDetails)
}

// Print a confirmation message after booking
func printConfirmationMessage(firstName, lastName string, ticketsToBook uint, email string, remainingTickets uint) {
	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v.\n", firstName, lastName, ticketsToBook, email)
	fmt.Printf("Tickets remaining: %v\n", remainingTickets)
}

// Display all users who have booked tickets
func displayBookings(bookingsList []string) {
	fmt.Printf("List of all bookings: %v\n", bookingsList)
}

// Display the first names of all attendees
func displayFirstNames(bookingsList []string) {
	var firstNames []string
	for _, booking := range bookingsList {
		firstName := strings.Fields(booking)[0] // Extract the first name
		firstNames = append(firstNames, firstName)
	}
	fmt.Printf("First names of all attendees: %v\n", firstNames)
}

// Check if all tickets are sold out
func checkSoldOut(remainingTickets uint) bool {
	if remainingTickets == 0 {
		fmt.Println("All tickets are sold out! See you at the conference.")
		return true
	}
	return false
}

// Validate the city entered by the user
func validateCity(city string) bool {
	// Define valid cities for the conference
	validCities := map[string]bool{
		"New York":     true,
		"Singapore":    true,
		"Berlin":       true,
		"London":       true,
		"Hong Kong":    true,
		"Mexico":       true,
	}

	// Check if the entered city is valid
	if validCities[city] {
		return true
	}
	// If city is invalid, return false
	return false
}

// Handle the city choice and provide feedback
func handleCityChoice(city *string) {
	if !validateCity(*city) {
		fmt.Println("The city you entered does not have a conference available. Defaulting to 'New York'.")
		*city = "New York" // Default to a valid city
	}
}

// Main booking logic
func bookTickets(remainingTickets *uint, bookingsList *[]string, bookedTickets *int) {
	for *remainingTickets > 0 && len(*bookingsList) < 50 {
		displayCurrentTicketsInfo(*remainingTickets, *bookedTickets)

		// Collect user details
		var firstName, lastName, email, city string
		var ticketsToBook uint
		collectUserInfo(&firstName, &lastName, &email, &city, &ticketsToBook, *remainingTickets)

		// Handle city choice (default to New York if invalid)
		handleCityChoice(&city)

		// Update ticket counts
		updateTicketCounts(remainingTickets, bookedTickets, ticketsToBook)

		// Store booking details
		storeBookingDetails(firstName, lastName, bookingsList)

		// Print confirmation message
		printConfirmationMessage(firstName, lastName, ticketsToBook, email, *remainingTickets)

		// Display current bookings
		displayBookings(*bookingsList)

		// Display first names of attendees
		displayFirstNames(*bookingsList)

		// Check if tickets are sold out
		if checkSoldOut(*remainingTickets) {
			break
		}
	}
}

func main() {
	// Conference details
	var conferenceName = "GoConference"
	const totalTickets = 100
	var remainingTickets uint = 100
	var bookedTickets = 0

	// Slice to store bookings dynamically
	var bookingsList []string

	// Start the application
	greetUsers(conferenceName)
	initialTicketInfo(totalTickets, remainingTickets)
	bookTickets(&remainingTickets, &bookingsList, &bookedTickets)
}
