package maple_test

import (
	c "cmp"
	"slices"
	"testing"

	"github.com/briandamaged/maple"
	"github.com/google/go-cmp/cmp"
)

func TestNewDefaulter(t *testing.T) {
	m := map[int]int{
		3: 10,
	}

	d := maple.DefaulterFor(&m, func(k int) int {
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

func TestInversePairings(t *testing.T) {
	m := map[string]int{}
	for _, w := range []string{"foo", "bar", "quuz", "blorp"} {
		m[w] = len(w)
	}

	ips := maple.InversePairings(m)
	if len(ips) != 3 {
		t.Error("expected 3 InversePairings")
	}

	slices.SortFunc(ips, func(x, y maple.InversePairing[string, int]) int {
		return c.Compare(x.Value, y.Value)
	})

	expected := []maple.InversePairing[string, int]{
		{
			Value: 3,
			Keys:  []string{"foo", "bar"},
		},
		{
			Value: 4,
			Keys:  []string{"quuz"},
		},
		{
			Value: 5,
			Keys:  []string{"blorp"},
		},
	}

	if diff := cmp.Diff(ips, expected); diff != "" {
		t.Error(diff)
	}
}
