package cli

import (
	"fmt"
	"os"
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

func TestNoCommandButOptions(t *testing.T) {
	line := "program -option1 -option2"
	c := NewFromString(line)
	command := c.GetCommand()
	if command != "program" {
		t.Errorf("The line is '%v', the command is '%v' (it should be 'program')", line, command)
	}
}

func TestNoCommand(t *testing.T) {
	line := "program"
	c := NewFromString(line)
	command := c.GetCommand()
	if command != "program" {
		t.Errorf("The line is '%v', the command is '%v' (it should be 'program')", line, command)
	}
}

func TestCommandAndOptions(t *testing.T) {
	line := "mycommand -option1 -option2 value1"
	c := NewFromString(line)
	command := c.GetCommand()
	if command != "mycommand" {
		t.Errorf("The commmand should be mycommand, it is '%v'", command)
	}
}

func TestIndexOf1(t *testing.T) {
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

func TestGetStringFromSetOrDefault(t *testing.T) {
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

func TestGetStringFromSetOrDie(t *testing.T) {
	line := "program mycommand -foo a"
	c := NewFromString(line)
	c.GetStringFromSetOrDie("-foo", []string{"a", "b", "c"})
}

func TestGetIntOrEnvOrDefault(t *testing.T) {
	line := "-mykey 10"
	c := NewFromString(line)
	value := c.GetIntOrDefault("-mykey", 6)
	if value != 10 {
		t.Errorf("Should be 10")
	}

	os.Setenv("MY_KEY", "")
	value2 := c.GetIntOrEnvOrDefault("-myotherkey", "MY_KEY", 10)
	if value2 != 10 {
		t.Errorf("Should be 10")
	}

	os.Setenv("MY_KEY", "50")
	value3 := c.GetIntOrEnvOrDefault("-myotherkey", "MY_KEY", 10)
	if value3 != 50 {
		t.Errorf("Should be 50, is %v\n", value3)
	}

}

func TestShift(t *testing.T) {
	line := "10 9 8 7 6 5 4 3 2 1"
	c := NewFromString(line)
	if c.GetCommand() != "10" {
		t.Errorf("Command should be ten.")
	}

	if c.Flatten() != "10 9 8 7 6 5 4 3 2 1" {
		t.Errorf("Flatten after 0 shift should be '10 9 8 7 6 5 4 3 2 1' but was '%v'", c.Flatten())
	}

	c.Shift()
	if c.Flatten() != "9 8 7 6 5 4 3 2 1" {
		t.Errorf("Flatten after 1 shift should be '9 8 7 6 5 4 3 2 1' but was '%v'", c.Flatten())
	}

	c.Shift()
	c.Shift()
	c.Shift()
	c.Shift()
	c.Shift()
	c.Shift()
	if c.Flatten() != "3 2 1" {
		t.Errorf("Flatten after 7 shifts should be '3 2 1' but was '%v'", c.Flatten())
	}

	c.Shift()
	c.Shift()
	c.Shift()
	if c.Flatten() != "" {
		t.Errorf("Flatten after 10 shifts should be '' but was '%v'", c.Flatten())
	}

	c.Shift()
	if c.Flatten() != "" {
		t.Errorf("Flatten after 10 shifts should be '' but was '%v'", c.Flatten())
	}

}

func TestVerbosityNone(t *testing.T) {
	line := "mycommand "
	c := NewFromString(line)
	if c.IS_VERBOSE {
		t.Error("not verbose")
	}
	if c.IS_VERBOSE2 {
		t.Error("not verbose")
	}
	if c.IS_VERBOSE3 {
		t.Error("not verbose")
	}
}

func TestVerbosityV1(t *testing.T) {
	line := "mycommand -v "
	c := NewFromString(line)
	if !c.IS_VERBOSE {
		t.Error("is verbose")
	}
	if c.IS_VERBOSE2 {
		t.Error("not verbose")
	}
	if c.IS_VERBOSE3 {
		t.Error("not verbose")
	}
}

func TestVerbosityV2(t *testing.T) {
	line := "mycommand -vv "
	c := NewFromString(line)
	if !c.IS_VERBOSE {
		t.Error("is verbose")
	}
	if !c.IS_VERBOSE2 {
		t.Error("is verbose")
	}
	if c.IS_VERBOSE3 {
		t.Error("not verbose")
	}
}

func TestVerbosityV3(t *testing.T) {
	line := "mycommand -vvv "
	c := NewFromString(line)
	if !c.IS_VERBOSE {
		t.Error("is verbose")
	}
	if !c.IS_VERBOSE2 {
		t.Error("is verbose")
	}
	if !c.IS_VERBOSE3 {
		t.Error("is verbose")
	}
}
