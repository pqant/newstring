package newstring

import "testing"

func Test(t *testing.T) {
	texts := []struct {
		key, value string
	}{
		{"Eralp", "plarE"},
		{"Ahmet", "temhA"},
		{"123", "321"},
	}
	for _, f := range texts {
		val := Reverse(f.key)
		if val != f.value {
			t.Errorf("Error : Value : %q ,Result :%q, Expected : %q", f.key, val, f.value)
		}
	}
}
