package untyped

import (
	"fmt"
	"strconv"
	"strings"
)

func parsePathElement(str string) (int, bool) {
	if len(str) < 3 { // Must have at least [0]
		return 0, false
	}

	if str[0:1] == "[" && str[len(str)-1:len(str)] == "]" {
		index := str[1 : len(str)-1]
		typed, err := strconv.Atoi(index)
		if err != nil {
			return 0, false
		}
		return typed, true
	}

	return 0, false
}

func newError(traversedPathElements []string, msg string) error {
	if len(traversedPathElements) == 0 {
		return fmt.Errorf(msg)
	}

	return fmt.Errorf("%s: %s", strings.Join(traversedPathElements, "."), msg)
}

// GetValue returns a value from inside an untyped (interface{}) structure consisting of maps and lists
//
// Examples:
//  Get the value of the key 'level2' in a nested map:
//
//    value, err := GetValue(testData, "level1.level2")
//
//  Get the value of key 'level2' in a map which itself is the second element of a list:
//
//    value, err := GetValue(testData, "[1].level2")
//
func GetValue(unstructured interface{}, path string) (interface{}, error) {
	pathElements := strings.Split(path, ".")
	traversedPathElements := []string{}

	var value interface{}
	for _, pathElement := range pathElements {
		index, isIndex := parsePathElement(pathElement)

		if isIndex {
			typed, ok := unstructured.([]interface{})
			if !ok {
				return nil, newError(traversedPathElements, "not a list")
			}
			if index > len(typed)-1 {
				return nil, newError(traversedPathElements, fmt.Sprintf("index %d is out of bounds", index))
			}

			value = typed[index]
		} else {
			key := pathElement

			typed, ok := unstructured.(map[string]interface{})
			if !ok {
				return nil, newError(traversedPathElements, "not a map")
			}

			value, ok = typed[key]
			if !ok {
				return nil, fmt.Errorf("cannot find key '%s'", key)
			}
		}

		unstructured = value
		traversedPathElements = append(traversedPathElements, pathElement)
	}

	return value, nil
}

// GetFloat64 is the same as GetValue but also converts the result into a float64
func GetFloat64(unstructured interface{}, path string) (float64, error) {
	result, err := GetValue(unstructured, path)
	if err != nil {
		return -1, err
	}

	resultFloat, ok := result.(float64)
	if !ok {
		return -1, fmt.Errorf("%s: not a number", path)
	}

	return resultFloat, nil
}

// GetString is the same as GetValue but also converts the result into a string
func GetString(unstructured interface{}, path string) (string, error) {
	result, err := GetValue(unstructured, path)
	if err != nil {
		return "", err
	}

	resultString, ok := result.(string)
	if !ok {
		return "", fmt.Errorf("%s: not a string", path)
	}

	return resultString, nil
}

// SetValue sets the value of an element inside an untyped (interface{}) structure consisting of maps and lists
//
// Examples:
//  Set the value of the key 'level2' in a nested map:
//
//    err := SetValue(testData, "level1.level2", "new")
//
//  Set the value of key 'level2' in a map which itself is the second element of a list:
//
//    err := SetValue(testData, "[1].level2", "new")
//
//  Set the value of the 3rd element in a list nested in a map:
//
//    err := SetValue(testData, "level1.[2]", "new")
//
func SetValue(unstructured interface{}, path string, value interface{}) error {
	// This function works by getting the parent element first and then trying
	// to set the value on that parent element.
	pathElements := strings.Split(path, ".")
	lastPathElement := pathElements[len(pathElements)-1]
	parentPathElements := pathElements[0 : len(pathElements)-1]
	parentPath := strings.Join(parentPathElements, ".")

	var parent interface{}
	if len(parentPath) == 0 {
		parent = unstructured
	} else {
		var err error
		parent, err = GetValue(unstructured, parentPath)
		if err != nil {
			return fmt.Errorf("cannot set value: %s", err)
		}
	}

	index, isIndex := parsePathElement(lastPathElement)
	if isIndex {
		typedParent, ok := parent.([]interface{})
		if !ok {
			return newError(parentPathElements, "not a list")
		}

		typedParent[index] = value
	} else {
		key := lastPathElement

		typedParent, ok := parent.(map[string]interface{})
		if !ok {
			return newError(parentPathElements, "not a map")
		}

		typedParent[key] = value
	}

	return nil
}
