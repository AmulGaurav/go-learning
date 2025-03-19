package main

import (
	"fmt"
)

const conferenceName = "Go Conference"
const conferenceTickets uint = 50

var (
	remainingTickets uint = 50
	// var bookings [50]string // Array declaration of size 50 of type string
	bookings = make([]UserData, 0) // slice declaration
)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

func main() {
	greetUsers(conferenceName, conferenceTickets)

	for remainingTickets != 0 {
		firstName, lastName, email, userTickets := getUserInput()

		isValidFirstName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, email, userTickets)

		if isValidFirstName && isValidEmail && isValidTicketNumber {
			bookTicket(conferenceName, firstName, lastName, email, userTickets)
			sendTicket(firstName, lastName, email, userTickets)

			printFirstNames()
		} else {
			if !isValidFirstName {
				fmt.Println("First name cannot be empty.")
			}

			if !isValidEmail {
				fmt.Println("Invalid email address")
			}

			if !isValidTicketNumber {
				fmt.Printf("We only have %d tickets remaining, so you can't book %d tickets\n", remainingTickets, userTickets)
			}
		}
	}

	fmt.Printf("%s is booked out. Please come back next year.\n", conferenceName)
}

func greetUsers(confName string, confTickets uint) {
	fmt.Printf("Welcome to our %s booking application\n", confName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(confName string, firstName string, lastName string, email string, userTickets uint) {
	remainingTickets -= userTickets

	userData := UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTickets,
	}

	// bookings[0] = firstName + " " + lastName // Array
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %s\n", remainingTickets, confName)
}

func printFirstNames() {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	fmt.Printf("These are all our bookings: %v\n", firstNames)
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	var ticket = fmt.Sprintf("%d tickets for %s %s", userTickets, firstName, lastName)

	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n%s \nto email address %s\n", ticket, email)
	fmt.Println("####################")
}
