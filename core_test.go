package maple_test

import (
	"slices"
	"testing"

	"github.com/briandamaged/maple"
	"github.com/google/go-cmp/cmp"
)

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
