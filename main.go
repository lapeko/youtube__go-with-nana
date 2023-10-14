package main

import (
	"booking-app/helpers"
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

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

func outputResult(email string, userTickets uint) {
	fmt.Printf("bookings: %v\n", bookings)
	fmt.Printf("Type of bookings: %T\n", bookings)
	fmt.Printf("Length of bookings: %v\n", len(bookings))
	fmt.Printf("Thank you %v for boocking %v tickets. You'll receive a confirmation email at %v.\n", bookings[len(bookings)-1], userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("There are all names of bookings: %v\n", helpers.GetFirstNames(bookings))
}

func askUserData() (string, string, string, uint) {
	var (
		firstName, lastName, email string
		userTickets                uint
	)

	for {
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		if len(firstName) >= 2 {
			break
		}

		fmt.Println("First name is not valid. Try again, please")
	}

	for {
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		if len(lastName) >= 2 {
			break
		}

		fmt.Println("Last name is not valid. Try again, please")
	}

	for {
		fmt.Println("Enter your email")
		fmt.Scan(&email)

		if strings.Contains(email, "@") {
			break
		}

		fmt.Println("Email is not valid. Try again, please")
	}

	for {
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		if remainingTickets >= userTickets {
			break
		}

		fmt.Printf("You can not book %v number of tickets. It are only %v tickets available\n", remainingTickets, userTickets)
	}

	return firstName, lastName, email, userTickets
}
