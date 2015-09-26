package cli

import "testing"

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

		res := flag.String(tt.flag, tt.value)

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
