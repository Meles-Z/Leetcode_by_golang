package main

import (
	"fmt"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func sortByPrice(flights []Flight) []Flight {
	n := len(flights)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare current flight with the next one
			if flights[j].Price < flights[j+1].Price {
				// Swap flights[j] and flights[j+1]
				flights[j], flights[j+1] = flights[j+1], flights[j]
			}
		}
	}
	return flights
}

func main() {
	flights := []Flight{
		{"New York", "London", 500},
		{"Paris", "Tokyo", 700},
		{"Mumbai", "Dubai", 200},
		{"Los Angeles", "Sydney", 1200},
	}

	sortedFlights := sortByPrice(flights)

	fmt.Println("Flights sorted by price (highest to lowest):")
	for _, flight := range sortedFlights {
		fmt.Printf("From %s to %s - $%d\n", flight.Origin, flight.Destination, flight.Price)
	}
}
