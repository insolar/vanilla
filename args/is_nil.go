package args

import "github.com/insolar/vanilla/reflectkit"

func IsNil(v interface{}) bool {
	return reflectkit.IsNil(v)
}

type ShuffleFunc func(n int, swap func(i, j int))
