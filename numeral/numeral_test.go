package numeral

import "testing"

func TestConvertToRoman(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{
			Description: "1 gets converted to I",
			Arabic:      1,
			Want:        "I",
		},
		{
			Description: "2 gets converted to II",
			Arabic:      2,
			Want:        "II",
		},
		{
			Description: "1 gets converted to III",
			Arabic:      3,
			Want:        "III",
		},
		{
			Description: "4 gets converted to IV (can't repeat more than 3 times)",
			Arabic:      4,
			Want:        "IV",
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
