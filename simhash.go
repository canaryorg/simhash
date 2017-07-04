// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// simhash package implements Charikar's simhash algorithm to generate a 64-bit
// fingerprint of a given document.
//
// simhash fingerprints have the property that similar documents will have a similar
// fingerprint. Therefore, the hamming distance between two fingerprints will be small
// if the documents are similar
package simhash

import (
	"bytes"

	"golang.org/x/text/unicode/norm"
)

// Returns a 64-bit simhash of the given feature set
func Simhash(fs FeatureSet) uint64 {
	return Fingerprint(Vectorize(fs.GetFeatures()))
}

// Returns a 64-bit simhash of the given bytes
func SimhashBytes(b [][]byte) uint64 {
	return Fingerprint(VectorizeBytes(b))
}

// Fingerprint returns a 64-bit fingerprint of the given vector.
// The fingerprint f of a given 64-dimension vector v is defined as follows:
//   f[i] = 1 if v[i] >= 0
//   f[i] = 0 if v[i] < 0
func Fingerprint(v Vector) uint64 {
	var f uint64
	for i := uint8(0); i < 64; i++ {
		if v[i] >= 0 {
			f |= (1 << i)
		}
	}
	return f
}

// Compare calculates the Hamming distance between two 64-bit integers
//
// Currently, this is calculated using the Kernighan method [1]. Other methods
// exist which may be more efficient and are worth exploring at some point
//
// [1] http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan
func Compare(a uint64, b uint64) uint8 {
	v := a ^ b
	var c uint8
	for c = 0; v != 0; c++ {
		v &= v - 1
	}
	return c
}

// Shingle returns the w-shingling of the given set of bytes. For example, if the given
// input was {"this", "is", "a", "test"}, this returns {"this is", "is a", "a test"}
func Shingle(w int, b [][]byte) [][]byte {
	if w < 1 {
		// TODO: use error here instead of panic?
		panic("simhash.Shingle(): k must be a positive integer")
	}

	if w == 1 {
		return b
	}

	if w > len(b) {
		w = len(b)
	}

	count := len(b) - w + 1
	shingles := make([][]byte, count)
	for i := 0; i < count; i++ {
		shingles[i] = bytes.Join(b[i:i+w], []byte(" "))
	}
	return shingles
}
