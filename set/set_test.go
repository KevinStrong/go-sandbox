package set_test

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kevinstrong/set"
)

func TestCachePutThingsIn(t *testing.T) {
	setOne := set.New[string]()
	setOne.Add("b")
	setOne.Add("d")

	setTwo := set.New[string]()
	setTwo.Add("c")
	setTwo.Add("e")

	unionSet := setOne.Union(*setTwo)
	got := unionSet.Members()
	sort.Slice(got, func(i, j int) bool {
		return got[i] < got[j]
	})
	want := []string{"b", "c", "d", "e"}
	if !cmp.Equal(got, want) {
		t.Fatalf(cmp.Diff(want, got))
	}
}

func TestConcurrency(t *testing.T) {
	s := set.New[int]()

	go func() {
		for i := 0; i < 1000; i++ {
			s.Members()
		}
	}()

	for i := 0; i < 1000; i++ {
		s.Add(i)
	}

}
