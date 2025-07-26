package main

import (
	"fmt"
	"hash/fnv"
)

func NewGloomFilter() *GloomFilter {
	return new(GloomFilter)
}

func main() {
	len := 40
	hashes := 5

	gloomFilter := NewGloomFilter()
	gloomFilter.InstantiateGloomFilter(len, hashes, MapHash)

	gloomFilter.AddItem("NISHANT")
	gloomFilter.AddItem("ARUSHI")

	fmt.Println(gloomFilter.CheckMembership("NISHANT"))
	fmt.Println(gloomFilter.CheckMembership("ARUSH1"))

}

func MapHash(f *GloomFilter, s string) uint64 {

	f.hash.SetSeed(f.seed)
	f.hash.WriteString(s)
	str := f.hash.Sum64()

	return str
}

func BasicHash(f *GloomFilter, s string) uint64 {

	h := fnv.New32a()
	h.Write([]byte(s))

	return uint64(h.Sum32())
}
