package cli

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

func Test_NoCommandButOptions(t *testing.T) {
	line := "program -option1 -option2"
	c := NewFromString(line)
	command := c.GetCommand()
	if command != "-option1" {
		t.Errorf("There should be no command when options are specified.  The line is '%v', the command is '%v' (it should be '')", line, command)
	}
}

func Test_NoCommand(t *testing.T) {
	line := "program"
	c := NewFromString(line)
	command := c.GetCommand()
	if command != "" {
		t.Errorf("There should be no command when nothign but the program is specified.  The line is '%v', the command is '%v' (it should be '')", line, command)
	}
}

func Test_CommandAndOptions(t *testing.T) {
	line := "program mycommand -option1 -option2 value1"
	c := NewFromString(line)
	command := c.GetCommand()
	if command != "mycommand" {
		t.Errorf("The commmand should be mycommand, it is '%v'", command)
	}
}

func Test_IndexOf1(t *testing.T) {
	line := "program -option1 -option2 value2"
	c := NewFromString(line)
	if c.IndexOf("-option1") != 1 {
		t.Errorf("line is '%v', indexOf(-option1) should be 0, was %v\n", line, c.IndexOf("-option1"))
	}
	if c.IndexOf("-option2") != 2 {
		t.Errorf("line is '%v', indexOf(-option2) shoudl be 0, was %v\n", line, c.IndexOf("-option2"))
	}

	option1value := c.GetStringOrDefault("-option1", "default")
	option2value := c.GetStringOrDefault("-option2", "default")

	if option1value != "default" {
		t.Errorf("option1 should be empty but instead it is '%v'\n", option1value)
	}

	if option2value != "value2" {
		t.Errorf("option2 should be 'value2' but instead it is '%v'\n", option2value)
	}

}

func Test_GetStringFromSetOrDefault(t *testing.T) {
	line := "program mycommand -foo bar"
	c := NewFromString(line)
	value := c.GetStringOrDefault("-foo", "default")
	if value != "bar" {
		t.Errorf("The value should be 'bar', it is '%v'", value)
	}

	value = c.GetStringFromSetOrDefault("-foo", "default", []string{"a", "b", "c"})
	if value != "default" {
		t.Errorf("The value should be 'bar', it is '%v'", value)
	}

}

func Test_GetStringFromSetOrDie(t *testing.T) {
	line := "program mycommand -foo a"
	c := NewFromString(line)
	c.GetStringFromSetOrDie("-foo", []string{"a", "b", "c"})
	// if value != "bar" {
	// 	t.Errorf("The value should be 'bar', it is '%v'", value)
	// }

}
