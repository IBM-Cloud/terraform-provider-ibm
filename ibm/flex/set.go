// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package flex

// Set provides a simple implementation of a set which stores unique strings.
type Set map[string]struct{}

var voidValue struct{}

// NewSet returns a new (empty) Set.
func NewSet() Set {
	return make(Set)
}

// Size returns the number of values stored in s.
func (s Set) Size() int {
	return len(s)
}

// Add adds one or more strings to s.
func (s Set) Add(v ...string) {
	for _, elem := range v {
		s[elem] = voidValue
	}
}

// Values returns a string slice containing the strings found in s.
func (s Set) Values() []string {
	values := []string{}
	for key := range s {
		values = append(values, key)
	}
	return values
}

// Contains returns true iff v is a value stored in s.
func (s Set) Contains(v string) bool {
	for elem := range s {
		if v == elem {
			return true
		}
	}
	return false
}

// Union returns a Set containing all the values contained in either s or o.
func (s Set) Union(o Set) Set {
	union := NewSet()
	union.Add(s.Values()...)
	union.Add(o.Values()...)
	return union
}

// Intersection returns a Set containing the values that appear in both s and o.
func (s Set) Intersection(o Set) Set {
	intersection := NewSet()

	// Add s's elements that are also in o.
	for v := range s {
		if o.Contains(v) {
			intersection.Add(v)
		}
	}

	return intersection
}

// Difference returns a Set containing the values of s that are not contained in o.
func (s Set) Difference(o Set) Set {
	diff := NewSet()

	for v := range s {
		if !o.Contains(v) {
			diff.Add(v)
		}
	}

	return diff
}
