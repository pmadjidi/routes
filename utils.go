package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func uniqueDsts(dsts []string) []latlong {
	dict := make(map[string]bool)
	uDsts := make([]latlong, 0)
	for _, dst := range dsts {
		_, ok := dict[dst]
		if !ok {
			uDsts = append(uDsts, latlong(dst))
			dict[dst] = true
		}
	}
	return uDsts
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (s *system) hash(key string) int {
	h := sha1.New()
	h.Write([]byte(key))
	sha1_hash := hex.EncodeToString(h.Sum(nil)) // for the sake of sperading keys across processors...
	intValue := 0
	for _, char := range sha1_hash {
		intValue += int(char)
	}
	bucket := intValue % s.procNumber
	fmt.Println(key, bucket, sha1_hash)
	return bucket
}

func getKeyFromLatLong(src, dst latlong) string {
	return string(src) + ":" + string(dst)
}
