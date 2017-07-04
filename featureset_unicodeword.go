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

	"golang.org/x/text/unicode/norm"
)

var unicodeBoundaries = regexp.MustCompile(`[\pL-_']+`)

// UnicodeWordFeatureSet is a feature set in which each word is a feature,
// all equal weight.
//
// See: http://blog.golang.org/normalization
// See: https://groups.google.com/forum/#!topic/golang-nuts/YyH1f_qCZVc
type UnicodeWordFeatureSet struct {
	b []byte
	f norm.Form
}

func NewUnicodeWordFeatureSet(b []byte, f norm.Form) *UnicodeWordFeatureSet {
	fs := &UnicodeWordFeatureSet{b, f}
	fs.normalize()
	return fs
}

func (w *UnicodeWordFeatureSet) normalize() {
	b := bytes.ToLower(w.f.Append(nil, w.b...))
	w.b = b
}

// Returns a []Feature representing each word in the byte slice
func (w *UnicodeWordFeatureSet) GetFeatures() []Feature {
	return getFeatures(w.b, unicodeBoundaries)
}
