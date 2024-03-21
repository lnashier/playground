package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testDef struct {
	Input        []any
	AccumulateFn func(any, any) any
	IDFn         func(any) any
	Output       []any
}

var tests = []testDef{
	{
		Input:        []any{1},
		AccumulateFn: func(t1, t2 any) any { return t1.(int) + t2.(int) },
		IDFn:         func(any) any { return nil },
		Output:       []any{1},
	},
	{
		Input:        []any{1, 2, 3, 4},
		AccumulateFn: func(t1, t2 any) any { return t1.(int) + t2.(int) },
		IDFn:         func(any) any { return nil },
		Output:       []any{10},
	},
	{
		Input:        []any{"a"},
		AccumulateFn: func(t1, t2 any) any { return t2.(string) + t1.(string) },
		IDFn:         func(any) any { return nil },
		Output:       []any{"a"},
	},
	{
		Input:        []any{"a"},
		AccumulateFn: func(t1, t2 any) any { return t2.(string) + t1.(string) },
		IDFn:         func(t any) any { return t },
		Output:       []any{"a"},
	},
	{
		Input:        []any{"a", "b", "c", "d"},
		AccumulateFn: func(t1, t2 any) any { return t2.(string) + t1.(string) },
		IDFn:         func(any) any { return nil },
		Output:       []any{"abcd"},
	},
	{
		Input:        []any{"a", "b", "c", "d"},
		AccumulateFn: func(t1, t2 any) any { return t2.(string) + t1.(string) },
		IDFn:         func(t any) any { return t },
		Output:       []any{"a", "b", "c", "d"},
	},
}

func TestAccumulatePrimitive(t *testing.T) {
	for _, theTest := range tests {
		result := Accumulate(theTest.Input, theTest.AccumulateFn, theTest.IDFn)
		assert.Truef(t, deepEqual(theTest.Output, result, theTest.IDFn), "Input: %v", theTest.Input)
	}
}

func deepEqual(v1, v2 []any, idFn func(any) any) bool {
	v1Map := make(map[any]any)
	for _, v := range v1 {
		v1Map[idFn(v)] = v
	}
	for _, v := range v2 {
		if v1Map[idFn(v)] != v {
			return false
		}
	}
	return true
}

var testObjs = []testDef{
	{
		Input: []any{newObj(1, 3)},
		AccumulateFn: func(t1, t2 any) any {
			t1.(*Obj).Val = t1.(*Obj).Val + t2.(*Obj).Val
			return t1
		},
		IDFn:   func(any) any { return nil },
		Output: []any{newObj(1, 3)},
	},
	{
		Input: []any{newObj(1, 3)},
		AccumulateFn: func(t1, t2 any) any {
			t1.(*Obj).Val = t1.(*Obj).Val + t2.(*Obj).Val
			return t1
		},
		IDFn:   func(t any) any { return t.(*Obj).ID },
		Output: []any{newObj(1, 3)},
	},
	{
		Input: []any{newObj(1, 3), newObj(2, 4), newObj(1, 5), newObj(3, 10), newObj(2, 6)},
		AccumulateFn: func(t1, t2 any) any {
			t1.(*Obj).Val = t1.(*Obj).Val + t2.(*Obj).Val
			return t1
		},
		IDFn:   func(t any) any { return t.(*Obj).ID },
		Output: []any{newObj(1, 8), newObj(2, 10), newObj(3, 10)},
	},
	{
		Input: []any{newObj("a", 3), newObj("b", 4), newObj("a", 5), newObj("c", 10), newObj("b", 6)},
		AccumulateFn: func(t1, t2 any) any {
			t1.(*Obj).Val = t1.(*Obj).Val + t2.(*Obj).Val
			return t1
		},
		IDFn:   func(t any) any { return t.(*Obj).ID },
		Output: []any{newObj("a", 8), newObj("b", 10), newObj("c", 10)},
	},
}

func TestAccumulateObj(t *testing.T) {
	for _, theTest := range testObjs {
		result := Accumulate(theTest.Input, theTest.AccumulateFn, theTest.IDFn)
		assert.Truef(t, deepEqualObj(theTest.Output, result, theTest.IDFn), "Input: %v", theTest.Input)
	}
}

func deepEqualObj(v1, v2 []any, idFn func(any) any) bool {
	v1Map := make(map[any]any)
	for _, v := range v1 {
		v1Map[idFn(v)] = v
	}
	for _, v := range v2 {
		if v1Map[idFn(v)].(*Obj).Val != v.(*Obj).Val {
			return false
		}
	}
	return true
}

type Obj struct {
	ID  any
	Val int
}

func newObj(id any, v int) *Obj {
	return &Obj{
		ID:  id,
		Val: v,
	}
}
