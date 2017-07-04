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
	"regexp"

	"golang.org/x/text/unicode/norm"
)

// FeatureSet represents a set of features in a given document
type FeatureSet interface {
	GetFeatures() []Feature
}

// Splits the given []byte using the given regexp, then returns a slice
// containing a Feature constructed from each piece matched by the regexp
func getFeatures(b []byte, r *regexp.Regexp) []Feature {
	words := r.FindAll(b, -1)
	features := make([]Feature, len(words))
	for i, w := range words {
		features[i] = NewFeature(w)
	}
	return features
}
