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

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		//bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets -= userTickets

		fmt.Printf("bookings: %v\n", bookings)
		fmt.Printf("Type of bookings: %T\n", bookings)
		fmt.Printf("Length of bookings: %v\n", len(bookings))

		fmt.Printf("Thank you %v for boocking %v tickets. You'll receive a confirmation email at %v.\n", bookings[0], userTickets, email)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

		for _, fullName := range bookings {
			firstNames = append(firstNames, strings.Fields(fullName)[0])
		}

		fmt.Printf("There are all names of bookings: %v\n", firstNames)
		firstNames = nil
	}
}
