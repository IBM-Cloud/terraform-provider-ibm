// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package flex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetEmpty(t *testing.T) {

	// Test empty set.
	s := NewSet()
	assert.NotNil(t, s)
	assert.Equal(t, []string{}, s.Values())
	assert.Equal(t, 0, s.Size())
	assert.False(t, s.Contains("foo"))

	// Test union of two empty sets
	r := NewSet()
	u := s.Union(r)
	assert.NotNil(t, u)
	assert.Equal(t, []string{}, u.Values())
	assert.Equal(t, 0, u.Size())
	// Test intersection of sets.
	s = NewSet()
	r = NewSet()
	i := s.Intersection(r)
	assert.NotNil(t, i)
	assert.Equal(t, 0, i.Size())

}

func TestSetAdd(t *testing.T) {
	s := NewSet()

	s.Add("foo")
	assert.Equal(t, []string{"foo"}, s.Values())
	assert.Equal(t, 1, s.Size())
	assert.True(t, s.Contains("foo"))
	assert.False(t, s.Contains("not there"))

	s.Add("bar")
	assert.Equal(t, 2, s.Size())
	assert.True(t, s.Contains("bar"))

	// Add duplicate (should be ignored).
	s.Add("foo")
	assert.Equal(t, 2, s.Size())
}

func TestSetUnion(t *testing.T) {
	s := NewSet()
	s.Add("foo")

	r := NewSet()
	r.Add("bar", "bar", "bar", "baz")
	assert.Equal(t, 2, r.Size())
	assert.True(t, r.Contains("bar"))
	assert.True(t, r.Contains("baz"))
	assert.False(t, r.Contains("foo"))

	u := s.Union(r)
	assert.NotNil(t, u)
	assert.Equal(t, 3, u.Size())
	assert.True(t, u.Contains("foo"))
	assert.True(t, u.Contains("bar"))
	assert.True(t, u.Contains("baz"))
}

func TestSetIntersection(t *testing.T) {
	s := NewSet()
	s.Add("n1", "n2", "n3")

	r := NewSet()
	r.Add("n2", "n4")

	// intersection of s and r.
	u := s.Intersection(r)
	assert.NotNil(t, u)
	assert.Equal(t, 1, u.Size())
	assert.True(t, u.Contains("n2"))

	// intersection of r and s should be the same.
	u = r.Intersection(s)
	assert.Equal(t, 1, u.Size())
	assert.True(t, u.Contains("n2"))
}

func TestSetDifference(t *testing.T) {
	s := NewSet()
	s.Add("n1", "n2", "n3", "n4")

	r := NewSet()
	r.Add("n1", "n4")

	u := s.Difference(r)
	assert.NotNil(t, u)
	assert.Equal(t, 2, u.Size())
	assert.True(t, u.Contains("n2"))
	assert.True(t, u.Contains("n3"))
}
