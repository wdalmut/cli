package cli

type Flag struct {
	args []string
}

func (f Flag) StringSlice(key string) []string {
	slice := []string{}

	for idx, elem := range f.args {
		if elem == key {
			slice = append(slice, f.args[idx+1])
		}
	}

	return slice
}

func (f Flag) String(key, value string) string {
	for idx, elem := range f.args {
		if elem == key {
			return f.args[idx+1]
		}
	}

	return value
}

func Parse(args []string) Flag {
	return Flag{args}
}
