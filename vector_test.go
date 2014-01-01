// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhash

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	var empty Vector
	v := Vectorize([]Feature{})
	if v != empty {
		t.Errorf("Expected empty vector for no features")
	}
}

func TestVectorize(t *testing.T) {
	features := []Feature{
		NewFeature([]byte("test string")),
		NewFeature([]byte("test thing")),
	}
	v := Vectorize(features)
	expected := Vector{0, -2, 0, 0, -2, -2, -2, 0, 0, 0, -2, 0, 0, 0, -2, 0, -2, 0, 0, 0, 0, 0, 2, 2, 2, -2, 0, -2, 0, -2, 2, 0, 0, 2, 0, 2, -2, 0, 2, 0, 2, -2, 0, 0, 2, 0, 0, 2, 0, 0, 0, 2, 0, 0, 2, -2, -2, 2, -2, 0, 2, 0, 2, -2}

	if v != expected {
		t.Errorf("Vectorize failed to return expected result")
	}
}

func TestVectorizeBytes(t *testing.T) {
	features := [][]byte{
		[]byte("test string"),
		[]byte("test thing"),
	}
	v := VectorizeBytes(features)
	expected := Vector{0, -2, 0, 0, -2, -2, -2, 0, 0, 0, -2, 0, 0, 0, -2, 0, -2, 0, 0, 0, 0, 0, 2, 2, 2, -2, 0, -2, 0, -2, 2, 0, 0, 2, 0, 2, -2, 0, 2, 0, 2, -2, 0, 0, 2, 0, 0, 2, 0, 0, 0, 2, 0, 0, 2, -2, -2, 2, -2, 0, 2, 0, 2, -2}

	if v != expected {
		t.Errorf("Vectorize failed to return expected result")
	}
}
