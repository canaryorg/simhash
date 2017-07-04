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
	"regexp"
)

var boundaries = regexp.MustCompile(`[\w']+(?:\://[\w\./]+){0,1}`)

// WordFeatureSet is a feature set in which each word is a feature,
// all equal weight.
type WordFeatureSet struct {
	b []byte
}

func NewWordFeatureSet(b []byte) *WordFeatureSet {
	fs := &WordFeatureSet{b}
	fs.normalize()
	return fs
}

func (w *WordFeatureSet) normalize() {
	w.b = bytes.ToLower(w.b)
}

// Returns a []Feature representing each word in the byte slice
func (w *WordFeatureSet) GetFeatures() []Feature {
	return getFeatures(w.b, boundaries)
}
