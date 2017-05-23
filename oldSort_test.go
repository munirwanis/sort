package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{in: []int{5, 9, 1, 4, 0, 3}, want: []int{0, 1, 3, 4, 5, 9}},
		{in: []int{5, 0, 98, 133, 0, 456, 666, 55, 666}, want: []int{0, 0, 5, 55, 98, 133, 456, 666, 666}},
	}

	for _, c := range cases {
		t.Logf("BubbleSort(%v)\n", c.in)
		BubbleSort(c.in)
		if IntArrayEquals(c.want, c.in) == false {
			t.Errorf("Came %v, want %v", c.in, c.want)
		}
	}
}

func TestQuickSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{in: []int{5, 9, 1, 4, 0, 3}, want: []int{0, 1, 3, 4, 5, 9}},
		{in: []int{5, 0, 98, 133, 0, 456, 666, 55, 666}, want: []int{0, 0, 5, 55, 98, 133, 456, 666, 666}},
	}

	for _, c := range cases {
		t.Logf("QuickSort(%v)\n", c.in)
		QuickSort(c.in)
		if IntArrayEquals(c.want, c.in) == false {
			t.Errorf("Came %v, want %v", c.in, c.want)
		}
	}
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
