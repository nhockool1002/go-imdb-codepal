package main

import (
	"errors"
	"fmt"
	"time"
)

// License represents a software license.
type License struct {
	ID        int
	Name      string
	ExpiresAt time.Time
}

// LicenseManager manages software licenses.
type LicenseManager struct {
	licenses []License
}

// NewLicenseManager creates a new instance of LicenseManager.
func NewLicenseManager() *LicenseManager {
	return &LicenseManager{
		licenses: make([]License, 0),
	}
}

// AddLicense adds a new license to the LicenseManager.
func (lm *LicenseManager) AddLicense(license License) {
	lm.licenses = append(lm.licenses, license)
}

// GetLicenseByID retrieves a license by its ID from the LicenseManager.
func (lm *LicenseManager) GetLicenseByID(id int) (License, error) {
	for _, license := range lm.licenses {
		if license.ID == id {
			return license, nil
		}
	}
	return License{}, errors.New("license not found")
}

// IsLicenseValid checks if a license is valid based on the current time.
func (lm *LicenseManager) IsLicenseValid(license License) bool {
	return license.ExpiresAt.After(time.Now())
}

// Usage Examples for LicenseManager

func main() {
	// Create a new LicenseManager
	lm := NewLicenseManager()

	// Add licenses
	license1 := License{
		ID:        1,
		Name:      "License 1",
		ExpiresAt: time.Now().AddDate(1, 0, 0), // Expires in 1 year
	}
	lm.AddLicense(license1)

	license2 := License{
		ID:        2,
		Name:      "License 2",
		ExpiresAt: time.Now().AddDate(2, 0, 0), // Expires in 2 years
	}
	lm.AddLicense(license2)

	// Get license by ID
	license, err := lm.GetLicenseByID(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("License:", license.Name)
	}

	// Check if license is valid
	isValid := lm.IsLicenseValid(license)
	fmt.Println("Is license valid?", isValid)
}
