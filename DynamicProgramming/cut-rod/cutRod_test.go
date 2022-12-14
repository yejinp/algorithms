package cutrod

import "testing"

func TestCutRod(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{18, 90}, 90},
		{[]int{1, 5}, 5},
		{[]int{1, 5, 8}, 8},
		{[]int{1, 5, 8, 9}, 10},
		{[]int{1, 5, 8, 9, 17, 17, 20, 24, 30}, 30},
	}

	for _, test := range tests {
		get := cutRod(test.input)
		if test.want != get {
			t.Errorf("cutRod(%v), got:%v, want:%v", test.input, get, test.want)
		}
	}
}
