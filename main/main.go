package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to %v and \n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "and", remainingTickets, "are still available")
	fmt.Println("Thanks Nana")

	var userName string
	var userTickets int
	var bookings = [50]string{"name1", "name2", "name3"}
	fmt.Scan(&userName)

	fmt.Printf("Usr %v booked%v\n", userName, userTickets)

}
