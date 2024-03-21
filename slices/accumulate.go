package slices

import (
	"golang.org/x/exp/maps"
)

func Accumulate[T any](ts []T, accumulateFn func(T, T) T, idFn func(T) any) []T {
	if len(ts) < 2 {
		return ts
	}
	accumulated := make(map[any]T)
	for _, t := range ts {
		id := idFn(t)
		if ot, ok := accumulated[id]; ok {
			t = accumulateFn(t, ot)
		}
		accumulated[id] = t
	}
	return maps.Values(accumulated)
}
