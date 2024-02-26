package utils

import (
	"github.com/google/uuid"
)

// GetCode is a function that returns a unique code for a file.
//
// Returns:
// - string: A unique code for a file.
//
// Example Usage:
//
//	code := GetCode()
//
// // use code to identify a file
func GetCode() string {
	uuid, _ := uuid.NewRandom()

	return uuid.String()
}
