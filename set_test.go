package set_test

import (
	"set"
	"testing"
)

type Set map[interface{}]bool

func (s Set) Add(x interface{}) (ret bool) {
	ret = !s.Contains(x)
	s[x] = true
	return
}

func (s Set) Discard(x interface{}) (ret bool) {
	ret = s.Contains(x)
	s[x] = false
	return
}

func (s Set) Contains(x interface{}) bool {
	return s[x]
}

func (s Set) Members() (mem []interface{}) {
	for m := range s {
		mem = append(mem, m)
	}
	return
}

func (s Set) Copy() (t set.Interface) {
	t = make(Set)
	for _, x := range s.Members() {
		t.Add(x)
	}
	return
}

func TestInterface(test *testing.T) {
	s := make(Set)
	for i := 0; i < 20; i++ {
		if !set.Add(s, i) {
			test.Errorf("Add() returned false for new element %v", i)
		}
		if !set.Contains(s, i) {
			test.Errorf("Contains() returned false for existing element %v", i)
		}
		t := set.Copy(s)
		if !set.Discard(s, i) {
			test.Errorf("Discard() returned false for existing element %v", i)
		}
		if set.Contains(s, i) {
			test.Errorf(
				"Contains() returned true for nonexisting element %v", i)
		}
	}
}

func TestSetOps(test *testing.T) {
	var u set.Interface
	s := make(Set)
	t := make(Set)
	for i := 0; i < 20; i += 2 {
		set.Add(s, i)
	}
	for i := 0; i < 20; i += 3 {
		set.Add(t, i)
	}

	u = set.Union(s, t)
	if !set.IsSubset(u, s) || !set.IsSubset(u, t) {
		test.Error(
			"Union() produced a set which doesn't contain all elements from " +
				"the source sets")
	}

	u = set.Intersection(s, t)
	for i := 0; i < 20; i++ {
		if i%6 == 0 && !set.Contains(u, i) {
			test.Errorf(
				"Intersection() produced a set lacking an element "+
					"that was in both sets: %v", i)
		} else if i%6 != 0 && set.Contains(u, i) {
			test.Errorf(
				"Intersection() produced a set containing an element that "+
					"was not in both sets: %v", i)
		}
	}

	u = set.Difference(s, t)
	for i := 0; i < 20; i += 2 {
		if i%6 == 0 && set.Contains(u, i) {
			test.Errorf(
				"Difference() produced a set that contains elements from "+
					"second set: %v", i)
		} else if i%6 != 0 && !set.Contains(u, i) {
			test.Errorf(
				"Difference() produced a set that lacks an element from the " +
					"first set which was not in the second set: %v", i)
		}
	}

}
