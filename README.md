# Magic-Query-Builder
This is a meta-language query builder implemented in Go. The meta-language consists of string-format text where . (dot) is the separator between keys (or different levels in a JSON) and [] is a special keyword that indicates the word that immediately precedes is a list.

## Applied implementation perspective

This is a Go program that constructs a JSON query from a list of filters and saves the result to a file named "generated_query.json".

The program starts with importing several packages, including "fmt", "os", "strings", and "encoding/json".

A "Query" type is defined as a map with string keys and interface{} values, which can hold any Go data type.

The main function of the program is "MagicBuildQuery". It takes a slice of strings as input, each string being a filter in the format of "key=value". The function constructs a Query type from these filters and returns it.

The function works by first creating an empty Query map, and adding a key "query" to it. Then, for each filter, the function splits the string by "=", trims the whitespaces of the two parts, and splits the key by ".". The value part is then assigned to the corresponding key in the map, using the keys as the path to the value.

If any key contains "[]", it means the value is an array, and the function adds a new key "match" to the current map, with its value being another map, which then has a key "eq" with the value of the filter.

Finally, the constructed Query is converted to a JSON string using the "json.MarshalIndent" function and saved to a file using the "os.Create" function.

In the "main" function, a slice of filters is defined, and the "MagicBuildQuery" is called with it as the input. The returned Query is then converted to a JSON string and saved to a file.

If any error occurs during the process, the program will print an error message and return. If everything goes well, the program will print a message indicating success.

## Issue with the Output

As i'm noob to this go lang so in this approch there is a problem in the function MagicBuildQuery() seems to be working as intended, as it splits the input string by "." and treats each segment as a key in a nested map. When it encounters a segment that ends with "[]", it treats the key as an array, which will contain maps that are matched against the conditions defined by the key-value pairs.

In the input "interests.[].sport.name": "football", the keys are split into the following segments: ["interests", "", "sport", "name"]. The second segment, "", is not a valid key. So in the output file it also generate it.
