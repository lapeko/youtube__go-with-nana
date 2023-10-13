package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" // var conferenceName = "Go Conference" would be the same
	const conferenceTickets uint = 50
	var remainingTickets = conferenceTickets
	//var bookings [conferenceTickets]string
	var bookings []string
	var firstNames []string

	fmt.Printf("Type of variables:\nconferenceName: %T, conferenceTickets: %T, remainingTickets: %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		firstNameCorrect := len(firstName) >= 2

		if !firstNameCorrect {
			fmt.Println("First name is not valid. Try again, please")
			continue
		}

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		lastNameCorrect := len(lastName) >= 2

		if !lastNameCorrect {
			fmt.Println("Last name is not valid. Try again, please")
			continue
		}

		fmt.Println("Enter your email")
		fmt.Scan(&email)
		emailCorrect := strings.Contains(email, "@")

		if !emailCorrect {
			fmt.Println("Email is not valid. Try again, please")
			continue
		}

		for {
			fmt.Println("Enter number of tickets")
			fmt.Scan(&userTickets)
			userTicketsCorrect := remainingTickets >= userTickets
			if userTicketsCorrect {
				remainingTickets -= userTickets
				break
			} else {
				fmt.Printf("You can not book %v number of tickets. It are only %v tickets available\n", remainingTickets, userTickets)
				continue
			}
		}

		fmt.Printf("bookings: %v\n", bookings)
		fmt.Printf("Type of bookings: %T\n", bookings)
		fmt.Printf("Length of bookings: %v\n", len(bookings))

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v for boocking %v tickets. You'll receive a confirmation email at %v.\n", bookings[0], userTickets, email)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

		for _, fullName := range bookings {
			firstNames = append(firstNames, strings.Fields(fullName)[0])
		}

		fmt.Printf("There are all names of bookings: %v\n", firstNames)
		firstNames = nil

		if remainingTickets == 0 {
			fmt.Println("Selling stops. No available remaining tickets")
			break
		}
	}
}
