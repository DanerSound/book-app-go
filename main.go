package main

import (
	"fmt"
	//"ticket-seller/helper" // se non aggiungi il nome della conferenza questo cerca di trovarlo nel suo path base non lo trova
)

// This is a definition of packege variable , so all function in packge main can see this variables
// and we dont have to pass them
const conferenceTickets int = 50

var conferenceName = "pycon 9"
var remainingTickets int = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets int

}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isValidEmail, isValidNumber := ValidateUserInput(firstName, lastName, email, userTicket) // devi mettere CAP la prima lettera altrimenti non funzionerà

		if isValidName && isValidEmail && isValidNumber {

			bookTicket(firstName, lastName, email, userTicket)

			sentTicket(firstName, lastName, userTicket, email )

			fmt.Printf("The names of the bookings are %v \n", currentbooked())

			if remainingTickets == 0 {
				// end the programm
				fmt.Println(" Our conference is booked, back the next year! ")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf(" first name or last name is too short\n")
			}
			if !isValidEmail {
				fmt.Printf(" mail is not valid!\n")
			}
			if !isValidName {
				fmt.Printf(" number of ticket is invalid! \n")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf(" Welcome to %v booking application\n", conferenceName)
	fmt.Printf(" we have %v we still have %v ! \n", conferenceTickets, remainingTickets)
	fmt.Println(" Get your tickets to attend! ")
}

func currentbooked() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTicket int

	fmt.Println(" Enter your firts name :")
	fmt.Scan(&firstName)
	fmt.Println(" Enter your last name :")
	fmt.Scan(&lastName)
	fmt.Println(" Enter your email :")
	fmt.Scan(&email)
	fmt.Println(" Enter the number of Thickets :")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(firstName string, lastName string, email string, userTicket int) {

	remainingTickets = remainingTickets - userTicket

	// create a map for each user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf(" The bookings : %v \n", bookings)

	fmt.Printf(" Thanks %v %v for booking %v Tickets, you will recive a conformation at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf(" %v remaning for %v \n", remainingTickets, conferenceName)

}

func sentTicket(firstName string, lastName string, userTicket int, email string){
	var ticket = fmt.Sprintf("%v tickets fot %v %v",userTicket, firstName, lastName)
	fmt.Printf(" ############# \n")
	fmt.Printf(" sending ticket:\n %v to email %v\n" , ticket, email)
	fmt.Printf(" ############# \n ")
}