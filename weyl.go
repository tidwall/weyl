package weyl

import (
	"math"
	"math/rand"
	"sync"
)

// Generator represents an random number generator
type Generator struct {
	x, w, s uint64
}

// Uint64 returns a random uint64
func (w *Generator) Uint64() uint64 {
	w.w += w.s
	w.x = ((w.x*w.x + w.w) >> 32) | ((w.x*w.x + w.w) << 32)
	return w.x
}

// Int63 returns a random Int63 returns a random int64 that is always a
// positive number.
func (w *Generator) Int63() int64 {
	return int64(w.Uint64() << 1 >> 1)
}

// Float64 returns a random float64
func (w *Generator) Float64() float64 {
	f := float64(w.Uint64()) / float64(math.MaxUint64)
again:
	if f == 1 {
		goto again
	}
	return f
}

// Seed or re-seed the generator
func (w *Generator) Seed(seed int64) {
	*w = Generator{s: (uint64(seed) << 1) + 0xb5ad4eceda1ce2a9}
}

// NewSource returns a new rand.Source for used with rand.New()
func NewSource(seed int64) rand.Source {
	var w Generator
	w.Seed(seed)
	return &w
}

// NewSource64 returns a new rand.Source64
func NewSource64(seed int64) rand.Source64 {
	var w Generator
	w.Seed(seed)
	return &w
}

var mu sync.Mutex
var w Generator

// Int63 returns a random int64
func Int63() int64 {
	mu.Lock()
	v := w.Int63()
	mu.Unlock()
	return v
}
