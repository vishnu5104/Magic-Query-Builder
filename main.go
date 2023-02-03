package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Query represents a query in key-value format
type Query map[string]interface{}

// BuildQuery constructs a query from the input query filters
func MagicBuildQuery(filters []string) Query {
	query := make(Query)
	query["__query__"] = make(map[string]interface{})
	for _, filter := range filters {
		keyValue := strings.Split(filter, "=")
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])
		keys := strings.Split(key, ".")
		current := query["__query__"].(map[string]interface{})
		for i, k := range keys {
			if strings.Contains(k, "[]") {
				arrayKey := strings.Trim(k, "[]")
				if _, done := current[arrayKey]; !done {
					current[arrayKey] = make(map[string]interface{})
					current[arrayKey].(map[string]interface{})["__match__"] = make(map[string]interface{})
				}
				if i == len(keys)-1 {
					current[arrayKey].(map[string]interface{})["__match__"].(map[string]interface{})["__eq__"] = value
					break
				}
				current = current[arrayKey].(map[string]interface{})["__match__"].(map[string]interface{})
				continue
			}
			if i == len(keys)-1 {
				current[k] = value
				break
			}
			if _, done := current[k]; !done {
				current[k] = make(map[string]interface{})
			}
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
