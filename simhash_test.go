// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhash

import (
	"testing"
)

func TestSimhash(t *testing.T) {
	var fp = []uint64{
		Simhash(WordFeatureSet{[]byte("this is a test phrase")}),
		Simhash(WordFeatureSet{[]byte("this is a test phrass")}),
		Simhash(WordFeatureSet{[]byte("foo bar")}),
	}

	if Compare(fp[0], fp[1]) != 2 {
		t.Errorf("Comparison failed")
	}

	if Compare(fp[0], fp[2]) != 29 {
		t.Errorf("Comparison failed")
	}
}
