package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference" // var conferenceName = "Go Conference" would be the same
const conferenceTickets uint = 50

var remainingTickets = conferenceTickets

// var bookings [conferenceTickets]string
var bookings []string

func main() {
	welcome()
	for {
		firstName, lastName, email, userTickets := askUserData()
		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets -= userTickets

		outputResult(email, userTickets)

		if remainingTickets == 0 {
			fmt.Println("Selling stops. No available remaining tickets")
			break
		}
	}
}

func welcome() {
	fmt.Printf("Type of variables:\nconferenceName: %T, conferenceTickets: %T, remainingTickets: %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(fullNames []string) []string {
	var firstNames []string
	for _, fullName := range bookings {
		firstNames = append(firstNames, strings.Fields(fullName)[0])
	}
	return firstNames
}

func outputResult(email string, userTickets uint) {
	fmt.Printf("bookings: %v\n", bookings)
	fmt.Printf("Type of bookings: %T\n", bookings)
	fmt.Printf("Length of bookings: %v\n", len(bookings))
	fmt.Printf("Thank you %v for boocking %v tickets. You'll receive a confirmation email at %v.\n", bookings[len(bookings)-1], userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("There are all names of bookings: %v\n", getFirstNames(bookings))
}

func askUserData() (string, string, string, uint) {
	var (
		firstName, lastName, email string
		userTickets                uint
	)
	firstNameCorrect, lastNameCorrect, emailCorrect, userTicketsCorrect := false, false, false, false

	for !firstNameCorrect {
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		firstNameCorrect = len(firstName) >= 2

		if !firstNameCorrect {
			fmt.Println("First name is not valid. Try again, please")
		}
	}

	for !lastNameCorrect {
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		lastNameCorrect = len(lastName) >= 2

		if !lastNameCorrect {
			fmt.Println("Last name is not valid. Try again, please")
		}
	}

	for !emailCorrect {
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		emailCorrect = strings.Contains(email, "@")

		if !emailCorrect {
			fmt.Println("Email is not valid. Try again, please")
		}
	}

	for !userTicketsCorrect {
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)
		userTicketsCorrect = remainingTickets >= userTickets
		if !userTicketsCorrect {
			fmt.Printf("You can not book %v number of tickets. It are only %v tickets available\n", remainingTickets, userTickets)
		}
	}

	return firstName, lastName, email, userTickets
}
