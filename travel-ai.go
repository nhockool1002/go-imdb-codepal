package main

import (
	"fmt"
	"time"
)

// Travel represents a travel itinerary.
type Travel struct {
	Source      string
	Destination string
	Departure   time.Time
	Arrival     time.Time
}

// NewTravel creates a new Travel instance.
//
// Parameters:
// source (string): The source location of the travel.
// destination (string): The destination location of the travel.
// departure (time.Time): The departure time of the travel.
// arrival (time.Time): The arrival time of the travel.
//
// Returns:
// *Travel: A pointer to the newly created Travel instance.
func NewTravel(source, destination string, departure, arrival time.Time) *Travel {
	return &Travel{
		Source:      source,
		Destination: destination,
		Departure:   departure,
		Arrival:     arrival,
	}
}

// Duration returns the duration of the travel.
//
// Returns:
// time.Duration: The duration of the travel.
func (t *Travel) Duration() time.Duration {
	return t.Arrival.Sub(t.Departure)
}

// TravelManager represents a travel management application.
type TravelManager struct {
	Travels []*Travel
}

// NewTravelManager creates a new TravelManager instance.
//
// Returns:
// *TravelManager: A pointer to the newly created TravelManager instance.
func NewTravelManager() *TravelManager {
	return &TravelManager{
		Travels: []*Travel{},
	}
}

// AddTravel adds a new travel to the travel manager.
//
// Parameters:
// travel (*Travel): The travel to be added.
func (tm *TravelManager) AddTravel(travel *Travel) {
	tm.Travels = append(tm.Travels, travel)
}

// GetTotalDuration returns the total duration of all travels in the travel manager.
//
// Returns:
// time.Duration: The total duration of all travels.
func (tm *TravelManager) GetTotalDuration() time.Duration {
	totalDuration := time.Duration(0)

	for _, travel := range tm.Travels {
		totalDuration += travel.Duration()
	}

	return totalDuration
}

func main() {
	// Create a new travel manager
	travelManager := NewTravelManager()

	// Add some travels
	travel1 := NewTravel("New York", "London", time.Date(2022, time.January, 1, 10, 0, 0, 0, time.UTC), time.Date(2022, time.January, 1, 15, 0, 0, 0, time.UTC))
	travel2 := NewTravel("London", "Paris", time.Date(2022, time.January, 2, 8, 0, 0, 0, time.UTC), time.Date(2022, time.January, 2, 10, 0, 0, 0, time.UTC))
	travel3 := NewTravel("Paris", "Rome", time.Date(2022, time.January, 3, 12, 0, 0, 0, time.UTC), time.Date(2022, time.January, 3, 15, 0, 0, 0, time.UTC))

	travelManager.AddTravel(travel1)
	travelManager.AddTravel(travel2)
	travelManager.AddTravel(travel3)

	// Get the total duration of all travels
	totalDuration := travelManager.GetTotalDuration()
	fmt.Printf("Total duration of all travels: %s\n", totalDuration)
}
