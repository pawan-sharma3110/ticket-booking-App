package main

import (
	"fmt"
	"strconv"
	"strings"
)

const cofrenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 1)

func main() {

	getUser()
	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := validateUser(firstName, lastName, email, userTickets)
		if isValidName && isValidEmail && isValidTickets {
			bookTickets(userTickets, firstName, lastName, email)
			printFirstName()
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
				fmt.Printf("please enter a valid number of tickets. We have only %v tickets left. !!", remainingTickets)
			}

		}

	}

}
func getUser() {
	fmt.Printf("Welcome to %s booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v remaining ticketes \n", cofrenceTickets, remainingTickets)
	fmt.Println("Get your tickets herre to attend")
}
func printFirstName() {
	firstNames := []string{}
	for _, v := range bookings {

		firstNames = append(firstNames, v["firstName"])
	}

	fmt.Printf("The first names of bookings are :%v \n", firstNames)

}
func validateUser(firstName, lastName, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}
func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName, lastName, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v/n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
