package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const cofrenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings = []string{}
	fmt.Printf("Welcome to %s booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining ticketes \n", cofrenceTickets, remainingTickets)
	fmt.Println("Get your tickets herre to attend")

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets
		if isValidName && isValidEmail && isValidTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v/n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, v := range bookings {
				var names = strings.Fields(v)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are :%v \n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our confrence is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("please enter firt name or last name valid !!")
			}
			if !isValidEmail {
				fmt.Println("please enter a valid email address with contain @ sign !!")
			}
			if !isValidTickets {
				fmt.Println("please enter a valid number of tickets !!")
			}

		}

	}

}
func getUser() {

}
