package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Query represents a query in key-value format
type Query map[string]interface{}

// MagicBuildQuery constructs a query from the input query filters
func MagicBuildQuery(filters []string) Query {
	// Initialize an empty Query object
	query := make(Query)
	// Create a map called "__query__" inside the Query object
	query["__query__"] = make(map[string]interface{})
	// Loop through each filter in the input filters slice
	for _, filter := range filters {
			// Split the filter into key and value
		keyValue := strings.Split(filter, "=")
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])
		// Split the key into keys
		keys := strings.Split(key, ".")
		// Create a pointer to the "__query__" map
		current := query["__query__"].(map[string]interface{})
		// Loop through each key in the keys slice
		for i, k := range keys {
				// Check if the current key contains "[]"
			if strings.Contains(k, "[]") {
				// Trim the brackets "[]" from the key
				arrayKey := strings.Trim(k, "[]")
					// If the current key is not present in the map, create it
				if _, done := current[arrayKey]; !done {
					
					current[arrayKey] = make(map[string]interface{})
					current[arrayKey].(map[string]interface{})["__match__"] = make(map[string]interface{})
				}
				// If this is the last key, set the value of the "__eq__" key to the value
				if i == len(keys)-1 {
					current[arrayKey].(map[string]interface{})["__match__"].(map[string]interface{})["__eq__"] = value
					break
				}
				// Update the current pointer to the "__match__" map
				current = current[arrayKey].(map[string]interface{})["__match__"].(map[string]interface{})
				continue
			}
				// If this is the last key, set the value of the key to the value
			if i == len(keys)-1 {
				current[k] = value
				break
			}
					// If the current key is not present in the map, create it
			if _, done := current[k]; !done {
				current[k] = make(map[string]interface{})
			}
			// Update the current pointer to the map of the current key
			current = current[k].(map[string]interface{})
		}
	}
	return query
}

func main() {
	filters := []string{
		"name.first_name =Sam",
		"address.country = United Kingdom",
		"interests.[].sport.name = football",
		"ingredients.[].milk.[].calcium = 10",
		"info.[].transport. [] =car",
	}
	query := MagicBuildQuery(filters)

	queryJSON, err := json.MarshalIndent(query, "", "  ")
	if err != nil {
		fmt.Println("Error converting query to JSON:", err)
		return
	}

	file, err := os.Create("generated_query.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(queryJSON)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Succesfully generated")
}
