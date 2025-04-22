package main

import (
	"fmt"
	"hash/fnv"
	"hash/maphash"
	"math/rand/v2"
)

func main() {
	len := 40
	hashes := 5

	gloomFilter, _ := newGloomFilter(len, hashes)
	gloomFilter.addItem("NISHANT")
	gloomFilter.addItem("ARUSHI")

	fmt.Println(gloomFilter.checkMembership("NISHANT"))
	fmt.Println(gloomFilter.checkMembership("ARUSHI"))

}

func newGloomFilter(length int, hashes int) (GloomFilter, error) {
	var gloomFilter GloomFilter
	if length < 1 {
		return gloomFilter, fmt.Errorf("length cannot be less than 1")
	}

	gloomFilter.len = length
	gloomFilter.createGloomArr()
	gloomFilter.createSeed()

	gloomFilter.hashes = make([]func(string) uint64, hashes)

	for index := range gloomFilter.hashes {

		// generation should be outside invokation obviously
		n := rand.Uint64N(100)

		gloomFilter.hashes[index] = func(s string) uint64 {
			return gloomFilter.mapHash(s) * n
		}
	}

	gloomFilter.hashLen = len(gloomFilter.hashes)

	return gloomFilter, nil
}

type GloomFilter struct {
	gloomArr []int
	seed     maphash.Seed
	len      int
	hash     maphash.Hash
	hashes   []func(string) uint64
	hashLen  int
}

func (f *GloomFilter) createGloomArr() {
	f.gloomArr = make([]int, f.len)
}

func (f *GloomFilter) createSeed() {
	f.seed = maphash.MakeSeed()
}

func (f *GloomFilter) addItem(s string) {

	for _, hashFunc := range f.hashes {
		hashInd := f.modHash(hashFunc(s))

		f.gloomArr[hashInd] = 1
	}
}

func (f *GloomFilter) checkMembership(s string) bool {

	for _, hashFunc := range f.hashes {
		hashInd := f.modHash(hashFunc(s))

		if f.gloomArr[hashInd] != 1 {
			return false
		}
	}
	return true
}

func (f *GloomFilter) mapHash(s string) uint64 {

	f.hash.SetSeed(f.seed)
	f.hash.WriteString(s)
	str := f.hash.Sum64()

	return str
}

func (f *GloomFilter) basicHash(s string) uint64 {
	h := fnv.New32a()

	h.Write([]byte(s))

	return uint64(h.Sum32())
}

func (f *GloomFilter) modHash(hash uint64) uint64 {
	return hash % uint64(f.len)
}
