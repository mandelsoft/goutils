// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package set

import (
	"cmp"
	"maps"
	"slices"

	"github.com/mandelsoft/goutils/maputils"
)

type Set[K comparable] map[K]struct{}

func New[K comparable](keys ...K) Set[K] {
	return Set[K]{}.Add(keys...)
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Both sets contain the same elements.
func (s1 Set[K]) Equal(s2 Set[K]) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

func (s Set[K]) Elements(yield func(e K) bool) {
	for e := range s {
		if !yield(e) {
			return
		}
	}
}

func (s Set[K]) Add(keys ...K) Set[K] {
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Set[K]) AddAll(set Set[K]) Set[K] {
	for k := range set {
		s[k] = struct{}{}
	}
	return s
}

func (s Set[K]) Delete(keys ...K) Set[K] {
	for _, k := range keys {
		delete(s, k)
	}
	return s
}

func (s Set[K]) DeleteAll(set Set[K]) Set[K] {
	if s != nil {
		for k := range set {
			delete(s, k)
		}
	}
	return s
}

func (s Set[K]) Contains(keys ...K) bool {
	if s == nil {
		return len(keys) == 0
	}
	for _, k := range keys {
		if _, ok := s[k]; !ok {
			return false
		}
	}
	return true
}

func (s Set[K]) ContainsAll(items ...K) bool {
	if s == nil {
		return false
	}
	for _, item := range items {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

func (s Set[K]) ContainsAny(items ...K) bool {
	for _, item := range items {
		if s.Contains(item) {
			return true
		}
	}
	return false
}

// GetAny fetches a single element from the set and removes it.
func (s Set[K]) GetAny() (K, bool) {
	for key := range s {
		s.Delete(key)
		return key, true
	}
	var zero K
	return zero, false
}

// Difference returns a set of objects that are not in s2.
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}
func (s1 Set[K]) Difference(s2 Set[K]) Set[K] {
	result := Set[K]{}
	for key := range s1 {
		if !s2.Contains(key) {
			result.Add(key)
		}
	}
	return result
}

// Union returns a new set which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}
func (s1 Set[K]) Union(s2 Set[K]) Set[K] {
	result := s1.Clone()
	for key := range s2 {
		result.Add(key)
	}
	return result
}

// SymmetricDifference returns a set of elements which are in either of the sets, but not in their intersection.
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.SymmetricDifference(s2) = {a3, a4, a5}
// s2.SymmetricDifference(s1) = {a3, a4, a5}
func (s1 Set[K]) SymmetricDifference(s2 Set[K]) Set[K] {
	result := Set[K]{}

	for key := range s2 {
		if !s1.Contains(key) {
			result.Add(key)
		}
	}
	for key := range s1 {
		if !s2.Contains(key) {
			result.Add(key)
		}
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (s1 Set[K]) Intersection(s2 Set[K]) Set[K] {
	var walk, other Set[K]
	result := Set[K]{}
	if len(s1) < len(s2) {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for key := range walk {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 Set[K]) IsSuperset(s2 Set[K]) bool {
	for item := range s2 {
		if !s1.Contains(item) {
			return false
		}
	}
	return true
}

// IsSubset returns true if and only if s1 is a subset of s2.
func (s1 Set[K]) IsSubset(s2 Set[K]) bool {
	return s2.IsSuperset(s1)
}

func (s Set[K]) Len() int {
	return len(s)
}

func (s Set[K]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[K]) Clone() Set[K] {
	return maps.Clone(s)
}

func (s Set[K]) WithAdded(elems ...K) Set[K] {
	return s.Clone().Add(elems...)
}

func (s Set[K]) Clear() bool {
	ok := len(s) != 0
	if ok {
		clear(s)
	}
	return ok
}

func (s Set[K]) AsArray() []K {
	keys := []K{}
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s Set[K]) Has(keys ...K) bool {
	return s.Contains(keys...)
}
func (s Set[K]) HasAny(keys ...K) bool {
	return s.ContainsAny(keys...)
}
func (s Set[K]) HasAll(keys ...K) bool {
	return s.ContainsAll(keys...)
}

// AsSortedArray returns the contents as a sorted K slice.
//
// This is a separate function and not a method because not all types supported
// by Generic are ordered and only those can be sorted.
func AsSortedArray[K cmp.Ordered](s Set[K]) []K {
	res := s.AsArray()
	slices.Sort(res)
	return res
}

func Keys[K comparable, V any](m map[K]V, cmp maputils.CompareFunc[K]) []K {
	return maputils.Keys(m, cmp)
}

func KeySet[K comparable, V any](m map[K]V) Set[K] {
	s := Set[K]{}
	for k := range m {
		s.Add(k)
	}
	return s
}
