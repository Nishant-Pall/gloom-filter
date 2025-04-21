package main

import (
	"fmt"
	"hash/fnv"
	"hash/maphash"
)

func main() {
	len := 40

	gloomFilter, _ := newGloomFilter(len)
	gloomFilter.addItem("NISHANT")
	gloomFilter.addItem("ARUSHI")

	fmt.Println(gloomFilter.checkMembership("NISHANT"))
	fmt.Println(gloomFilter.checkMembership("ARUSHI"))

	// arr := createGloomArr(len)
	// seed := maphash.MakeSeed()

	// addItem("Nishant", len, arr, seed)
	// addItem("Arsuhi", len, arr, seed)
	// fmt.Println(checkMembership("Arsuhi", len, arr, seed))

}

func newGloomFilter(len int) (GloomFilter, error) {
	var gloomFilter GloomFilter
	if len < 1 {
		return gloomFilter, fmt.Errorf("length cannot be less than 1")
	}

	gloomFilter.len = len
	gloomFilter.createGloomArr()
	gloomFilter.createSeed()

	return gloomFilter, nil
}

type GloomFilter struct {
	gloomArr []int
	seed     maphash.Seed
	len      int
	hash     maphash.Hash
}

func (f *GloomFilter) createGloomArr() {
	f.gloomArr = make([]int, f.len)
}

func (f *GloomFilter) createSeed() {
	f.seed = maphash.MakeSeed()
}

func (f *GloomFilter) addItem(s string) {

	fmt.Println(f.seed)
	h1 := f.basicHash(s)
	h2 := f.mapHash(s)

	modHash1 := f.modHash(h1)
	modHash2 := f.modHash(h2)

	f.gloomArr[modHash1] = 1
	f.gloomArr[modHash2] = 1

}

func (f *GloomFilter) checkMembership(s string) bool {

	fmt.Println(f.seed)
	h1 := f.basicHash(s)
	h2 := f.mapHash(s)

	modHash1 := f.modHash(h1)
	modHash2 := f.modHash(h2)

	return bool(f.gloomArr[modHash1] == 1 && f.gloomArr[modHash2] == 1)
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

func createGloomArr(len int) []int {
	arr := make([]int, len)
	return arr
}

func hash(s string) uint32 {

	h := fnv.New32a()

	h.Write([]byte(s))

	return h.Sum32()
}

func mapHash(s string, seed maphash.Seed) uint64 {

	var h maphash.Hash
	h.SetSeed(seed)
	h.WriteString(s)
	str := h.Sum64()

	return str

}

func addItem(item string, len int, gloomArr []int, seed maphash.Seed) {

	// item is hashed through k hash functions
	h1 := mapHash(item, seed)
	h2 := hash(item)

	// modulo of n (length of bit array) is executed to identify the k array positions(buckets)
	mod1 := uint32(h1) % uint32(len)
	mod2 := uint32(h2) % uint32(len)

	// the bits at all identified buckets are set to one
	gloomArr[mod1] = 1
	gloomArr[mod2] = 1

}

func checkMembership(item string, len int, gloomArr []int, seed maphash.Seed) bool {

	// item is hashed through k hash functions
	h1 := mapHash(item, seed)
	h2 := hash(item)

	// modulo of n (length of bit array) is executed to identify the k array positions(buckets)
	mod1 := uint32(h1) % uint32(len)
	mod2 := uint32(h2) % uint32(len)

	// check if all the bits are set to 1
	return bool(gloomArr[mod1] == 1 && gloomArr[mod2] == 1)
}
