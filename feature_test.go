// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhash

import (
	"testing"
)

func TestNewFeature(t *testing.T) {
	expected := uint64(8811532157352841348)
	f := NewFeature([]byte("test string"))

	if f.Weight() != 1 {
		t.Errorf("feature.Weight(): expected 1, actual %d", f.Weight())
	}

	if f.Sum() != expected {
		t.Errorf("feature.Sum(): expected %d, actual %d", expected, f.Sum())
	}
}

func TestNewFeatureWithWeight(t *testing.T) {
	weight := 10
	expected := uint64(8811532157352841348)
	f := NewFeatureWithWeight([]byte("test string"), weight)

	if f.Weight() != weight {
		t.Errorf("feature.Weight(): expected %d, actual %d", weight, f.Weight())
	}

	if f.Sum() != expected {
		t.Errorf("feature.Sum(): expected %d, actual %d", expected, f.Sum())
	}
}
