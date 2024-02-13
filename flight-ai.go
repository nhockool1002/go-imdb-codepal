package main

import (
	"fmt"
	"time"
)

// Flight represents a flight in the flight management application.
type Flight struct {
	FlightNumber string
	Departure    time.Time
	Arrival      time.Time
}

// NewFlight creates a new Flight instance with the given flight number, departure time, and arrival time.
func NewFlight(flightNumber string, departure, arrival time.Time) *Flight {
	return &Flight{
		FlightNumber: flightNumber,
		Departure:    departure,
		Arrival:      arrival,
	}
}

// Duration returns the duration of the flight.
func (f *Flight) Duration() time.Duration {
	return f.Arrival.Sub(f.Departure)
}

// FlightManager represents a flight management application.
type FlightManager struct {
	Flights []*Flight
}

// NewFlightManager creates a new FlightManager instance.
func NewFlightManager() *FlightManager {
	return &FlightManager{
		Flights: []*Flight{},
	}
}

// AddFlight adds a flight to the flight manager.
func (fm *FlightManager) AddFlight(flight *Flight) {
	fm.Flights = append(fm.Flights, flight)
}

// GetFlightByNumber returns the flight with the given flight number.
func (fm *FlightManager) GetFlightByNumber(flightNumber string) *Flight {
	for _, flight := range fm.Flights {
		if flight.FlightNumber == flightNumber {
			return flight
		}
	}
	return nil
}

// Usage Example for FlightManager

func main() {
	// Create a new flight manager
	flightManager := NewFlightManager()

	// Add flights to the flight manager
	flight1 := NewFlight("ABC123", time.Date(2022, time.January, 1, 8, 0, 0, 0, time.UTC), time.Date(2022, time.January, 1, 10, 0, 0, 0, time.UTC))
	flightManager.AddFlight(flight1)

	flight2 := NewFlight("DEF456", time.Date(2022, time.January, 2, 12, 0, 0, 0, time.UTC), time.Date(2022, time.January, 2, 14, 0, 0, 0, time.UTC))
	flightManager.AddFlight(flight2)

	// Get flight by flight number
	flightNumber := "ABC123"
	flight := flightManager.GetFlightByNumber(flightNumber)
	if flight != nil {
		fmt.Printf("Flight %s found\n", flightNumber)
		fmt.Printf("Duration: %s\n", flight.Duration())
	} else {
		fmt.Printf("Flight %s not found\n", flightNumber)
	}
}
