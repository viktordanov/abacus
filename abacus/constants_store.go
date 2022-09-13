package abacus

import "github.com/viktordanov/golang-lru"

type Constant string

const (
	PI  Constant = "pi"
	PHI Constant = "phi"
	E   Constant = "e"
)

type ConstantsStore struct {
	Constants map[Constant]Number
	LogCache  *lru.Cache[string, Number] // TODO: migrate back to hashicorp/golang-lru when it supports generics
}

func (cs *ConstantsStore) Constant(constant Constant) (Number, bool) {
	value, ok := cs.Constants[constant]
	return value, ok
}

func NewConstantsStore() (*ConstantsStore, error) {
	cache, err := lru.New[string, Number](1000)
	if err != nil {
		return nil, err
	}
	return &ConstantsStore{
		Constants: map[Constant]Number{},
		LogCache:  cache,
	}, nil
}

func (cs *ConstantsStore) SetConstant(constant Constant, value Number) {
	cs.Constants[constant] = value
}

func (cs *ConstantsStore) Cache(key string, value Number) {
	cs.LogCache.Add(key, value)
}

func (cs *ConstantsStore) Cached(key string) (Number, bool) {
	value, ok := cs.LogCache.Get(key)
	return value, ok
}
