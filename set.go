// Copywrong 2012 slowpoke <proxypoke at lavabit dot com>
// This program is Free Software under the terms of the
// Do What The Fuck You Want To Public License, 
// found in the COPYING file or at http://sam.zoy.org/wtfpl/COPYING

// TODO: documentation for package and functions

package set

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
