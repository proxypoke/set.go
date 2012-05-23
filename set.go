// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// This program is Free Software under the terms of the
// Do What The Fuck You Want To Public License, 
// found in the COPYING file or at http://sam.zoy.org/wtfpl/COPYING

// Package set provides set operations for any type implementing the
// set.Interface. A set is an unordered collection of non-repeating values
// which usually has constant time for adding, removing and querying
// membership of element. They also support powerful set-theoretical operations
// such as union and intersection.
//
// This interface doesn't guarantee these times, though, as it depends on the
// underlying type implementing it. A good candidate would be, for example, a
// map[T]bool, where T are the elements in the set and the boolean value
// indicating whether an element is part of the set.
//
// Sets can be used anywhere duplicate values would not make sense and fast
// membership tests are desired, for example a tagging mechanism or a friend
// list. 
package set

// If a type implements set.Interface, it can be used as a set. It's preferable
// for these methods to have constant time complexity, but this isn't
// necessary.
// 
// NOTE: All of these methods are meant for this package to call them. Use the
// appropriate function in this package instead (eg. don't use
// type.Add(x), use set.Add(type, x))
type Interface interface {
	Add(x interface{}) bool
	Discard(x interface{}) bool
	Contains(x interface{}) bool
	Members() []interface{}
	Copy() Interface
}

// Add adds an element to the set. Its time complexity is O(1). It returns true
// if that element wasn't in the set before, and false if it was.
func Add(set Interface, x interface{}) bool {
	return set.Add(x)
}

// Discard removes an from the set. Its time complexity is O(1). It returns
// true when the element was in the set, and false if it wasn't.
func Discard(set Interface, x interface{}) bool {
	return set.Discard(x)
}

// Contains returns a boolean value indicating whether x is in the set (true)
// or not (false). Its time complexity is O(1).
func Contains(set Interface, x interface{}) bool {
	return set.Contains(x)
}

// Members returns a slice containing all elements in the set. Its time
// complexity is O(n), where n = Len(set).
func Members(set Interface) []interface{} {
	return set.Members()
}

// Copy returns a new instance of the implementing type containing all valyes
// of the given instance.
func Copy(set Interface) Interface {
	return set.Copy()
}

// Len gives the size (or cardinality) or a set. Its time complexity is O(n),
// where n is the number of elements in the set.
func Len(set Interface) int {
	return len(set.Members())
}

// IsSubset indicates whether a t is a subset of s (true) or not (false). Its
// time complexity is O(n) where n = Len(t).
func IsSubset(s, t Interface) bool {
	for _, x := range t.Members() {
		if !s.Contains(x) {
			return false
		}
	}
	return true
}

// Union does a set union over s and t, returning a set that contains all
// elements which are in either set. Its time complexity is // O(n+m), where 
// n = Len(s) and m = Len(t).
func Union(s, t Interface) (u Interface) {
	u = s.Copy()
	for _, x := range t.Members() {
		u.Add(x)
	}
	return
}

// Intersection does a set intersection over s and t, returning a set that
// contains all elements which are in both sets. Its time complexity is O(n),
// where n = Len(s).
func Intersection(s, t Interface) (u Interface) {
	u = s.Copy()
	for _, x := range s.Members() {
		if !t.Contains(x) {
			u.Discard(x)
		}
	}
	return
}

// Difference does a set difference over s and t, returning a set that contains
// the elements in s which are not in t. Its time complexity is O(n), where 
// n = Len(t).
func Difference(s, t Interface) (u Interface) {
	u = s.Copy()
	for _, x := range t.Members() {
		if u.Contains(x) {
			u.Discard(x)
		}
	}
	return
}
