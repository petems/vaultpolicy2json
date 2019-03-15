package removesingletonarrays

import (
	"encoding/json"
)

// removeOneElementSlice is a function to loop through an array and
// check if it contains a certain string
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// removeOneElementSlice is a recursive function to loop through a map and remove
// any slices with one element, with the addition of ignoreing an array of keys given
func removeOneElementSlice(data map[string]interface{}, keysToIgnore []string) {
	for k, v := range data {
		if !stringInSlice(k, keysToIgnore) {
			switch t := v.(type) {
			case []interface{}:
				if len(t) == 1 {
					data[k] = t[0]
				}
				for _, tv := range t {
					if m, ok := tv.(map[string]interface{}); ok {
						removeOneElementSlice(m, keysToIgnore)
					}
				}
			case map[string]interface{}:
				removeOneElementSlice(t, keysToIgnore)
			}
		}
	}
}

// RemoveSingletonArrays iterates through a given JSON string
// and replaces any arrays with only one element with the just the
// one element of that array
//
// Arrays with more than one element are unaffected
//
//  RemoveSingletonArrays(`{"a":1,"b":["2"]}`)  // returns `{"a":1,"b":"2"}`
//  RemoveSingletonArrays(`{"a":1,"b":["2, 3"]}`) // returns `{"a":1,"b":["2","3"]}`
//
func RemoveSingletonArrays(jsonString string) (string, error) {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		return "", err
	}

	removeOneElementSlice(data, []string{})

	buf, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(buf), nil

}

// WithIgnores iterates through a given JSON string
// and replaces any arrays with only one element with the just the
// one element of that array, if they are not contained in keys to ignore
//
// Arrays with more than one element are unaffected
//
//  WithIgnores(`{"a":1,"b":["2"]}`, []string{"b"})  // returns `{"a":1,"b":["2"]}`
//  WithIgnores(`{"a":1,"b":["2"]}`, []string{"a"})  // returns `{"a":1,"b":"2"}`
//
func WithIgnores(jsonString string, keysToIgnore []string) (string, error) {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		return "", err
	}

	removeOneElementSlice(data, keysToIgnore)

	buf, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(buf), nil

}
