package maple_test

import (
	"slices"
	"testing"

	"github.com/briandamaged/maple"
	"github.com/google/go-cmp/cmp"
)

func TestNewDefaulter(t *testing.T) {
	m := map[int]int{
		3: 10,
	}

	d := maple.DefaulterFor(m, func(k int) int {
		return 2 * k
	})

	expected := map[int]int{
		3: 10,
		4: 8,
		1: 2,
	}

	for k, v := range expected {
		if d.Get(k) != v {
			t.Errorf("expected d.Get(%d) to return %d", k, v)
		}
	}

	if diff := cmp.Diff(m, expected); diff != "" {
		t.Error(diff)
	}
}

func TestInvert(t *testing.T) {
	m := map[string]int{
		"foo":   1,
		"bar":   2,
		"quuz":  1,
		"blarg": 3,
		"fwoop": 1,
	}

	im := maple.Invert(m)

	expected := map[int][]string{
		1: {"foo", "fwoop", "quuz"},
		2: {"bar"},
		3: {"blarg"},
	}

	for v, expectedKeys := range expected {
		if keys, exists := im[v]; exists {
			slices.Sort(keys)
			if diff := cmp.Diff(keys, expectedKeys); diff != "" {
				t.Error(diff)
			}
		} else {
			t.Errorf("expected `im[%d]` to exist", v)
		}
	}
}
