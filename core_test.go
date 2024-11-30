package maple_test

import (
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
