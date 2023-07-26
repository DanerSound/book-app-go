package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTicket int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidNumber := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidNumber
}
