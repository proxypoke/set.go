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
}

func Add(set Interface, x interface{}) bool {
	return set.Add(x)
}

func Discard(set Interface, x interface{}) bool {
	return set.Discard(x)
}

func Contains(set Interface, x interface{}) bool {
	return set.Contains(x)
}

func Members(set Interface, x interface{}) []interface{} {
	return set.Members()
}

func Len(set Interface) int {
	return len(set.Members())
}

func IsSubset(s, t Interface) bool {
	for _, x := range t.Members() {
		if !s.Contains(x) {
			return false
		}
	}
	return true
}

func Union(s, t Interface) (u Interface) {
	for _, x := range s.Members() {
		u.Add(x)
	}
	for _, x := range t.Members() {
		u.Add(x)
	}
	return
}

func Intersection(s, t Interface) (u Interface) {
	for _, x := range s.Members() {
		if t.Contains(x) {
			u.Add(x)
		}
	}
	return
}

func Difference(s, t Interface) (u Interface) {
	for _, x := range s.Members() {
		if !t.Contains(x) {
			u.Add(x)
		}
	}
	return
}
