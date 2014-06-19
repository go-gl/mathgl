package mgl32

import (
	"sync"
)

var (
	slicePools []*sync.Pool
	listLock   sync.RWMutex
)

func getPool(i int) *sync.Pool {
	listLock.RLock()
	if i >= len(slicePools) {

		// Promote to a write lock because we now
		// need to mutate the pool
		listLock.RUnlock()
		listLock.Lock()
		defer listLock.Unlock()

		for n := i - len(slicePools); n >= 0; n-- {
			newFunc := genPoolNew(1 << uint(len(slicePools)))
			slicePools = append(slicePools, &sync.Pool{New: newFunc})
		}
	} else {
		defer listLock.RUnlock()
	}

	return slicePools[i]
}

func genPoolNew(i int) func() interface{} {
	return func() interface{} {
		return make([]float32, 0, i)
	}
}

func grabFromPool(size int) []float32 {
	pool, exact := binLog(size)

	// Tried to grab something of size
	// zero or less
	if pool == -1 {
		return nil
	}

	// If the log is not exact, we
	// need to "overallocate" so we have
	// log+1
	if !exact {
		pool++
	}

	slice := getPool(pool).Get().([]float32)
	slice = slice[:size]
	return slice
}

func returnToPool(slice []float32) {
	if cap(slice) == 0 {
		return
	}

	pool, exact := binLog(cap(slice))

	if !exact {
		panic("attempt to pool slice with non-exact cap")
	}

	getPool(pool).Put(slice)
}

// The integer base 2 log of the value
// and whether the log is exact or rounded down
//
// This is only for positive integers
//
// Faster ways to do this, open to suggestions
func binLog(val int) (int, bool) {
	if val <= 0 {
		return -1, false
	}

	exact := true
	l := 0
	for ; val > 1; val = val >> 1 {
		// If the current lsb is 1 and the number
		// is not equal to 1, this is not an exact
		// log, but rather a rounding of it
		if val&1 != 0 {
			exact = false
		}
		l++
	}

	return l, exact
}
