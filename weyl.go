package weyl

import (
	"math/rand"
	"sync"
)

type Generator struct {
	x, w, s uint64
}

func (w *Generator) Uint64() uint64 {
	w.w += w.s
	w.x = ((w.x*w.x + w.w) >> 32) | ((w.x*w.x + w.w) << 32)
	return w.x
}

func (w *Generator) Int63() int64 {
	return int64(w.Uint64() << 1 >> 1)
}

func (w *Generator) Seed(seed int64) {
	*w = Generator{s: (uint64(seed) << 1) + 0xb5ad4eceda1ce2a9}
}

func NewSource(seed int64) rand.Source {
	var w Generator
	w.Seed(seed)
	return &w
}

func NewSource64(seed int64) rand.Source64 {
	var w Generator
	w.Seed(seed)
	return &w
}

var mu sync.Mutex
var w Generator

func Int63() int64 {
	mu.Lock()
	v := w.Int63()
	mu.Unlock()
	return v
}
