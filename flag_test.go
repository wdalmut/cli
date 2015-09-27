package cli

import "testing"

func TestBooleanExtractions(t *testing.T) {
	data := []struct {
		args []string
		flag string
		out  bool
	}{
		{[]string{"t"}, "t", true},
		{[]string{"t", "p", "s"}, "t", true},
		{[]string{"t"}, "w", false},
	}

	for _, tt := range data {
		flag := Parse(tt.args)

		res := flag.Bool(tt.flag)

		if res != tt.out {
			t.Errorf("Invalid flag, wants %s got %s", tt.out, res)
		}
	}
}

func TestAddressWithMultipleFlags(t *testing.T) {
	data := []struct {
		args []string
		out  bool
	}{
		{[]string{"-t", "-p", "-c"}, true},
		{[]string{"--testing", "-p", "-c"}, true},
		{[]string{"-p", "-c"}, false},
	}

	for _, tt := range data {
		flag := Parse(tt.args)

		res := flag.Bool("-t", "--testing")

		if res != tt.out {
			t.Errorf("Invalid flag, wants %s got %s", tt.out, res)
		}
	}
}

func TestSimpleStringExtractions(t *testing.T) {
	data := []struct {
		args  []string
		flag  string
		value string
		out   string
	}{
		{[]string{"w", "2"}, "w", "5", "2"},
		{[]string{}, "w", "5", "5"},
	}

	for _, tt := range data {
		flag := Parse(tt.args)

		res := flag.String(tt.value, tt.flag)

		if res != tt.out {
			t.Errorf("Invalid flag, wants %s got %s", tt.out, res)
		}
	}
}

func TestAddressStringWithDifferentKeys(t *testing.T) {
	data := []struct {
		args  []string
		value string
		out   string
	}{
		{[]string{"--testing", "new value"}, "default value", "new value"},
		{[]string{"--testing", "new value", "-p", "ok"}, "default value", "new value"},
		{[]string{}, "default value", "default value"},
	}

	for _, tt := range data {
		flag := Parse(tt.args)

		res := flag.String(tt.value, "-t", "--testing")

		if res != tt.out {
			t.Errorf("Invalid flag, wants %s got %s", tt.out, res)
		}
	}
}

func TestSimpleStringSliceExtractions(t *testing.T) {
	data := []struct {
		args  []string
		flag  string
		count int
		out   []string
	}{
		{[]string{}, "w", 0, []string{}},
		{[]string{"w", "2"}, "w", 1, []string{"2"}},
		{[]string{"w", "2", "w", "3"}, "w", 2, []string{"2", "3"}},
		{[]string{"w", "2", "w", "3", "w", "4"}, "w", 3, []string{"2", "3", "4"}},
	}

	for _, tt := range data {
		flag := Parse(tt.args)

		res := flag.StringSlice(tt.flag)

		if len(res) != tt.count {
			t.Errorf("Invalid extaction, wants %s elements got %s", tt.count, len(res))
		}

		for idx, e := range res {
			if e != tt.out[idx] {
				t.Errorf("Invalid element, wants %s got %s", tt.out[idx], e)
			}
		}
	}
}

func TestAddressStringSliceWithDifferentKeys(t *testing.T) {
	data := []struct {
		args  []string
		count int
		out   []string
	}{
		{[]string{}, 0, []string{}},
		{[]string{"--testing", "2"}, 1, []string{"2"}},
		{[]string{"--testing", "2", "--testing", "3"}, 2, []string{"2", "3"}},
		{[]string{"--testing", "2", "--testing", "3", "-t", "4"}, 3, []string{"2", "3", "4"}},
	}

	for _, tt := range data {
		flag := Parse(tt.args)

		res := flag.StringSlice("-t", "--testing")

		if len(res) != tt.count {
			t.Errorf("Invalid extaction, wants %s elements got %s", tt.count, len(res))
		}

		for idx, e := range res {
			if e != tt.out[idx] {
				t.Errorf("Invalid element, wants %s got %s", tt.out[idx], e)
			}
		}
	}
}
