package main

import (
	"testing"
)

func TestMaxMin0(t *testing.T) {
	k := int32(3)
	arr := []int32{10, 100, 300, 200, 1000, 20, 30}

	unfairness := maxMin(k, arr)

	if unfairness != 20 {
		t.Errorf("got %d instead of 20", unfairness)
	}

}

func TestMaxMin1(t *testing.T) {
	k := int32(4)
	arr := []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200}

	unfairness := maxMin(k, arr)

	if unfairness != 3 {
		t.Errorf("got %d instead of 3", unfairness)
	}

}

func TestMaxMin2(t *testing.T) {
	k := int32(2)
	arr := []int32{1, 2, 1, 2, 1}

	unfairness := maxMin(k, arr)

	if unfairness != 0 {
		t.Errorf("got %d instead of 0", unfairness)
	}

}

func TestMaxMin3(t *testing.T) {
	k := int32(5)
	arr := []int32{4504, 1520, 5857, 4094, 4157, 3902, 822, 6643, 2422, 7288, 8245, 9948, 2822, 1784, 7802, 3142, 9739, 5629, 5413, 7232}

	unfairness := maxMin(k, arr)

	if unfairness != 1335 {
		t.Errorf("got %d instead of 1335", unfairness)
	}

}
