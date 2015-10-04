package cli

import "strconv"

type Flag struct {
	args []string
}

// Check if a flag exists in the argument list
func (f Flag) Bool(keys ...string) bool {
	for _, elem := range f.args {
		for _, key := range keys {
			if elem == key {
				return true
			}
		}
	}
	return false
}

// Extract a list of parameters with the same key
func (f Flag) StringSlice(keys ...string) []string {
	slice := []string{}

	for idx, elem := range f.args {
		for _, key := range keys {
			if elem == key {
				slice = append(slice, f.args[idx+1])
			}
		}
	}

	return slice
}

// Return an integer flag value
// If an error occurs, the function will return the default value
// Example:
//   - flag.Int(5, "-x", "--extract")
func (f Flag) Int(value int, keys ...string) int {
	for idx, elem := range f.args {
		for _, key := range keys {
			if elem == key {
				i, err := strconv.Atoi(f.args[idx+1])
				if err != nil {
					return value
				}

				return i
			}
		}
	}

	return value
}

func (f Flag) String(value string, keys ...string) string {
	for idx, elem := range f.args {
		for _, key := range keys {
			if elem == key {
				return f.args[idx+1]
			}
		}
	}

	return value
}

func Parse(args []string) Flag {
	return Flag{args}
}
