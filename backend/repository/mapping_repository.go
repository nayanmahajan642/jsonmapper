package repository

import (
	"strings"
)

// ExtractValueFromPath navigates a nested JSON using a path string like "customer.displayName"
// It returns the value at that nested path or nil if the path doesn't exist
func ExtractValueFromPath(data map[string]interface{}, path string) interface{} {
	// Split the path string into keys, e.g. "customer.displayName" → ["customer", "displayName"]
	keys := strings.Split(path, ".")

	// Start from the top-level input JSON (data)
	var current interface{} = data

	// Traverse through each key step-by-step
	for _, key := range keys {
		// If current is still a map, try to go one level deeper
		if m, ok := current.(map[string]interface{}); ok {
			current = m[key] // go to next level
		} else {
			// If we can't go deeper (e.g. path is invalid), return nil
			return nil
		}
	}

	// Return the final value found at the given path
	return current
}

// DoMapping constructs the output JSON by mapping values from input JSON using the mapping config
// input  = the original nested JSON from the user
// mapping = a flat JSON that says which field maps to what path in the input JSON
func DoMapping(input map[string]interface{}, mapping map[string]string) map[string]interface{} {
	// Final result object to return
	result := make(map[string]interface{})

	// Loop over each key in the mapping
	for outKey, path := range mapping {
		// If mapping value is empty (""), set result key to empty string
		if path == "" {
			result[outKey] = ""
			continue
		}

		// Extract the value from the input JSON based on the path
		value := ExtractValueFromPath(input, path)

		// Set the extracted value into the result using the new output key
		result[outKey] = value
	}

	// Return the fully mapped output JSON
	return result
}
