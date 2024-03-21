package slices

import (
	"golang.org/x/exp/maps"
)

// Accumulate accumulates elements in the slice based on a specified ID function and accumulation function.
// It combines elements with the same ID using the accumulation function.
func Accumulate[T any](ts []T, accumulateFn func(T, T) T, idFn func(T) any) []T {
	if len(ts) < 2 {
		return ts
	}
	accumulated := make(map[any]T)
	for _, t := range ts {
		id := idFn(t)
		// If an element with the same ID exists in the map,
		// apply the accumulation function to combine them.
		if ot, ok := accumulated[id]; ok {
			t = accumulateFn(t, ot)
		}
		accumulated[id] = t
	}
	return maps.Values(accumulated)
}
