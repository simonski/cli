package goutils

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := NewFromString("run -file fred.txt c d -e 'hello'")
	fmt.Printf("CLI contains %d\n", c.GetIntOrDefault("-fooo", 4))

	// t.Errorf("Something %d %q", len(c.Args), c.Args)
}

// func TestIndexOf(t *testing.T) {
// 	for index := 0; index < len(c.Args); index++ {
// 		if c.Args[index] == key {
// 			return index
// 		}
// 	}
// 	return -1
// }

// func TestSplitStringToInts(t *testing.T) {
// 	columns := strings.Split(cols, delim)
// 	result := make([]int, len(columns))
// 	for index := 0; index < len(columns); index++ {
// 		str_value := columns[index]
// 		int_value, _ := strconv.Atoi(str_value)
// 		result[index] = int_value
// 	}
// 	return result
// }

// func TestGetStringOrDie(t *testing.T) {
// 	index := c.IndexOf(key)
// 	if index == -1 {
// 		fmt.Printf("Fatal: '%s' is required.\n", key)
// 		os.Exit(1)
// 		return ""
// 	} else {
// 		if index+1 < len(c.Args) {
// 			testValue := c.Args[index+1]
// 			if testValue[0:1] == "-" {
// 				// then there is no value
// 				return ""
// 			} else {
// 				return testValue
// 			}
// 		} else {
// 			return ""
// 		}
// 	}
// }

// func TestGetUIntOrDie(t *testing.T) {
// 	value := c.GetStringOrDie(key)
// 	v, err := strconv.Atoi(value)
// 	if err != nil {
// 		fmt.Printf("Fatal: '%s' should be an integer.\n", key)
// 		os.Exit(1)
// 		return -1
// 	}
// 	return v
// }

// func TestGetFileExistsOrDie(t *testing.T) {
// 	filename := c.GetStringOrDie(key)
// 	if filename == "" {
// 		fmt.Printf("Fatal: '%s' does not have a value.\n", key)
// 		os.Exit(1)
// 		return ""
// 	}

// 	if c.FileExists(filename) {
// 		return filename
// 	} else {
// 		fmt.Printf("Fatal: '%s' does not exist.\n", filename)
// 		os.Exit(1)
// 		return ""
// 	}
// }

// func TestFileExists(t *testing.T) {
// 	result, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		return false
// 	}
// 	return !result.IsDir()
// }
