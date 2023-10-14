package main

import (
	"booking-app/helpers"
	"booking-app/models"
	"fmt"
	"strings"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets = conferenceTickets
var orders []models.Order
var wg = sync.WaitGroup{}

func main() {
	welcome()
	for {
		order := askOrderData()
		orders = append(orders, order)
		remainingTickets -= order.UserTickets

		wg.Add(1)
		go outputResult(remainingTickets, orders, order)

		if remainingTickets == 0 {
			fmt.Println("Selling stops. No available remaining tickets")
			break
		}
	}
	wg.Wait()
}

func welcome() {
	fmt.Printf("Type of variables:\nconferenceName: %T, conferenceTickets: %T, remainingTickets: %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

func outputResult(remainingTickets uint, orders []models.Order, order models.Order) {
	time.Sleep(30 * time.Second)
	fmt.Printf("orders: %v\n", orders)
	fmt.Printf("Type of orders: %T\n", orders)
	fmt.Printf("Length of orders: %v\n", len(orders))
	fmt.Printf("Thank you %v for boocking %v tickets. You'll receive a confirmation email at %v.\n", orders[len(orders)-1], order.UserTickets, order.Email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("There are all names of bookings: %v\n", helpers.GetFirstNames(orders))
	wg.Done()
}

func askOrderData() models.Order {
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

	return models.Order{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		UserTickets: userTickets,
	}
}
