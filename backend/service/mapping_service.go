package service

import (
	"errors"

	"github.com/nayanmahajan642/jsonmapper/repository"
)

// MapService is responsible for handling the mapping logic at the service level.
// It validates the incoming data and calls the repository to perform the actual mapping.
//
// Parameters:
// - requestJson: the original nested JSON object sent by the user (parsed as map[string]interface{})
// - requestMapping: a flat key-to-path mapping (map[string]string) that describes how to extract values
//
// Returns:
// - A new JSON object (map[string]interface{}) after applying the mapping
// - An error if input is invalid
func MapService(requestJson map[string]interface{}, requestMapping map[string]string) (map[string]interface{}, error) {

	// Basic input validation to ensure both JSON and Mapping are provided
	if requestJson == nil || requestMapping == nil {
		return nil, errors.New("invalid input data") // return error if either input is missing
	}

	// Delegate the actual mapping logic to the repository layer
	// It returns the transformed/mapped JSON object
	return repository.DoMapping(requestJson, requestMapping), nil
}
