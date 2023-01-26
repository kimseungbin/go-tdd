package numeral

import "testing"

func TestConvertToRoman(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("got %q, wawnt %q", got, want)
	}
}
