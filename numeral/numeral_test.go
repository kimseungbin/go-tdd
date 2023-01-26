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
		{
			Description: "5 gets converted to V",
			Arabic:      5,
			Want:        "V",
		},
		{
			Description: "9 gets converted to IX",
			Arabic:      9,
			Want:        "IX",
		},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
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
