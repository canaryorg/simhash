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
	"golang.org/x/text/unicode/norm"
	"hash/fnv"
)

// Feature consists of a 64-bit hash and a weight
type Feature interface {
	// Sum returns the 64-bit sum of this feature
	Sum() uint64

	// Weight returns the weight of this feature
	Weight() int
}

type feature struct {
	sum    uint64
	weight int
}

// Sum returns the 64-bit hash of this feature
func (f feature) Sum() uint64 {
	return f.sum
}

// Weight returns the weight of this feature
func (f feature) Weight() int {
	return f.weight
}

// Returns a new feature representing the given byte slice, using a weight of 1
func NewFeature(f []byte) feature {
	h := fnv.New64()
	h.Write(f)
	return feature{h.Sum64(), 1}
}

// Returns a new feature representing the given byte slice with the given weight
func NewFeatureWithWeight(f []byte, weight int) feature {
	fw := NewFeature(f)
	fw.weight = weight
	return fw
}
