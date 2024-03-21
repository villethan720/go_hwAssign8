package main

import "fmt"

func main() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50 //uint is a small integer
	conferenceName := "Pioneer Conference"
	booking := []string{}

	fmt.Println("Welcome to %s booking application. \nWe have a total of %v are still available. \n Get your tickets here to attend. \n thank you!", conferenceName, remainingTickets)

	//Declare Data Types
	var firstname string
	var lastname string
	var email string
	var userTickets uint

	//collect data
	fmt.Println("Enter first name")
	fmt.Scanln(&firstname)
	fmt.Println("Enter last name")
	fmt.Scanln(&lastname)
	fmt.Println("Enter email")
	fmt.Scanln(&email)
	fmt.Println("How many tickets")
	fmt.Scanln(&userTickets)

	//logic for booking system
	remainingTickets = remainingTickets - userTickets
	booking = append(booking, firstname+" "+lastname)

	//Output
	fmt.Println("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v\n", firstname, lastname, userTickets, email)
	fmt.Println("There are %v tickets left for the %v\n", remainingTickets, conferenceName)
	fmt.Println("These are all the bookings. \n", booking)
}
