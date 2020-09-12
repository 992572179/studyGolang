package te

import "testing"

func TestAdd(t *testing.T) {
	test := []struct {
		a, b   int
		result int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{3, 4, 7},
	}
	for _, ite := range test {
		actual := Add(ite.a, ite.b)
		if actual != ite.result {
			t.Errorf("excepted %d but got %d", ite.result, actual)
		}
	}
}
