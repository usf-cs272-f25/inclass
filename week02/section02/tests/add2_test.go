package main

import "testing"

func TestAdd2(t *testing.T) {
	tests := []struct {
		a, b, want int
		name       string
	}{
		{1, 1, 2, "1+1=2"},
		{1, -1, 0, "1+-1=0"},
		{0, 7, 7, "0+7=7"},
	}

	for _, test := range tests {
		got := Add2(test.a, test.b)
		if got != test.want {
			t.Errorf("Add2(%d, %d) = %d but wanted: %d\n", test.a, test.b, got, test.want)
		}
	}
}
