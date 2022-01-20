package cache_test

import (
	"testing"
	"time"

	"github.com/kevinstrong/cache"

	"github.com/google/go-cmp/cmp"
)

func TestCacheExpiration(t *testing.T) {
	expiration := 1000 * time.Millisecond
	c := cache.New[string](expiration)

	testString := "/v1/users/1"
	c.Add(testString)

	if !cmp.Equal(c.Members(), []string{testString}) {
		t.Fatal(cmp.Diff(c.Members(), []string{testString}))
	}

	time.Sleep(expiration)

	if !cmp.Equal(c.Members(), []string{}) {
		t.Fatal(cmp.Diff(c.Members(), []string{}))
	}

}
