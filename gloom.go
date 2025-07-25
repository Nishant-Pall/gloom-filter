package main

import (
	"fmt"
	"hash/maphash"
	"math/rand/v2"
)

type GloomFilterHashFunc func(*GloomFilter, string) uint64

type GloomFilter struct {
	gloomArr []int
	seed     maphash.Seed
	len      int
	hash     maphash.Hash
	hashes   []func(string) uint64
	hashLen  int
}

func (gloomFilter *GloomFilter) InstantiateGloomFilter(length int, hashes int, hashFunc GloomFilterHashFunc) error {
	if length < 1 {
		return fmt.Errorf("length cannot be less than 1")
	}

	gloomFilter.len = length
	gloomFilter.CreateGloomArr()
	gloomFilter.CreateSeed()
	gloomFilter.GenerateHashFunctions(hashes, hashFunc)
	gloomFilter.hashLen = len(gloomFilter.hashes)

	return nil
}

func (f *GloomFilter) CreateGloomArr() {
	f.gloomArr = make([]int, f.len)
}

func (f *GloomFilter) CreateSeed() {
	f.seed = maphash.MakeSeed()
}

func (f *GloomFilter) GenerateHashFunctions(hashes int, hasFunc GloomFilterHashFunc) {
	f.hashes = make([]func(string) uint64, hashes)

	for index := range f.hashes {

		// generation should be outside invokation obviously
		n := rand.Uint64N(100)
		f.hashes[index] = func(s string) uint64 {
			return MapHash(f, s) * n
		}
	}

}

func (f *GloomFilter) AddItem(s string) {

	for _, hashFunc := range f.hashes {
		hashInd := f.ModHash(hashFunc(s))

		f.gloomArr[hashInd] = 1
	}
}

func (f *GloomFilter) CheckMembership(s string) bool {

	for _, hashFunc := range f.hashes {
		hashInd := f.ModHash(hashFunc(s))

		if f.gloomArr[hashInd] != 1 {
			return false
		}
	}
	return true
}

func (f *GloomFilter) ModHash(hash uint64) uint64 {
	return hash % uint64(f.len)
}
