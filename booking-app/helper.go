package main

import "strings"

func validateUserInput(firstName string, email string, userTickets uint) (bool, bool, bool) {
	isValidFirstName := len(firstName) >= 1
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidFirstName, isValidEmail, isValidTicketNumber
}
