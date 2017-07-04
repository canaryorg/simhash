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
	"hash/fnv"
)

type Vector [64]int

// Vectorize generates 64 dimension vectors given a set of features.
// Vectors are initialized to zero. The i-th element of the vector is then
// incremented by weight of the i-th feature if the i-th bit of the feature
// is set, and decremented by the weight of the i-th feature otherwise.
func Vectorize(features []Feature) Vector {
	var v Vector
	for _, feature := range features {
		sum := feature.Sum()
		weight := feature.Weight()
		for i := uint8(0); i < 64; i++ {
			bit := ((sum >> i) & 1)
			if bit == 1 {
				v[i] += weight
			} else {
				v[i] -= weight
			}
		}
	}
	return v
}

// VectorizeBytes generates 64 dimension vectors given a set of [][]byte,
// where each []byte is a feature with even weight.
//
// Vectors are initialized to zero. The i-th element of the vector is then
// incremented by weight of the i-th feature if the i-th bit of the feature
// is set, and decremented by the weight of the i-th feature otherwise.
func VectorizeBytes(features [][]byte) Vector {
	var v Vector
	h := fnv.New64()
	for _, feature := range features {
		h.Reset()
		h.Write(feature)
		sum := h.Sum64()
		for i := uint8(0); i < 64; i++ {
			bit := ((sum >> i) & 1)
			if bit == 1 {
				v[i]++
			} else {
				v[i]--
			}
		}
	}
	return v
}
