package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		//diferent order func
		firtsName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firtsName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			firtsNames := getFirstNames(bookings)
			fmt.Printf("The first name of bookings are: %v\n", firtsNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Out conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email adress  you entered doesn't contain @ sign")

			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")

			}
		}
	}

}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking aplication\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firtsName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firtsName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber

}

func getUserInput() (string, string, string, uint) {
	var firtsName string
	var lastName string
	var email string
	var userTickets uint

	//asking for user input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firtsName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of ticket: ")
	fmt.Scan(&userTickets)

	return firtsName, lastName, email, userTickets

}

func bookTicket(
	remainingTickets uint,
	userTickets uint,
	bookings []string,
	firtsName string,
	lastName string,
	email string,
	conferenceName string) {

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firtsName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firtsName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
