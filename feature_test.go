// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhash

import (
	"golang.org/x/text/unicode/norm"
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

func TestFeatureSet(t *testing.T) {
	text := []byte("here's a test string.")
	fs := NewWordFeatureSet(text)
	expected := []Feature{
		NewFeature([]byte("here's")),
		NewFeature([]byte("a")),
		NewFeature([]byte("test")),
		NewFeature([]byte("string")),
	}
	actual := fs.GetFeatures()

	for i := 0; i < len(actual); i++ {
		if actual[i].Sum() != expected[i].Sum() {
			t.Errorf("feature.Sum(): expected %d, actual %d", expected[i].Sum(), actual[i].Sum())
		}
		if actual[i].Weight() != expected[i].Weight() {
			t.Errorf("feature.Weight(): expected %d, actual %d", expected[i].Weight(), actual[i].Weight())
		}
	}
}

func TestUnicodeWordFeatureSet(t *testing.T) {
	text := []byte("la fin d'un bel après-midi d'été")
	fs := NewUnicodeWordFeatureSet(text, norm.NFKC)
	expected := []Feature{
		NewFeature([]byte("la")),
		NewFeature([]byte("fin")),
		NewFeature([]byte("d'un")),
		NewFeature([]byte("bel")),
		NewFeature([]byte("après-midi")),
		NewFeature([]byte("d'été")),
	}

	actual := fs.GetFeatures()

	for i := 0; i < len(actual); i++ {
		if actual[i].Sum() != expected[i].Sum() {
			t.Errorf("feature.Sum(): expected %d, actual %d", expected[i].Sum(), actual[i].Sum())
		}
		if actual[i].Weight() != expected[i].Weight() {
			t.Errorf("feature.Weight(): expected %d, actual %d", expected[i].Weight(), actual[i].Weight())
		}
	}
}

func TestGetFeatures(t *testing.T) {
	actual := getFeatures([]byte("test string"), boundaries)
	expected := []Feature{NewFeature([]byte("test")), NewFeature([]byte("string"))}

	if len(actual) != len(expected) {
		t.Errorf("getFeatures returned wrong number of features")
	}

	for i := 0; i < len(actual); i++ {
		if actual[i].Sum() != expected[i].Sum() {
			t.Errorf("feature.Sum(): expected %d, actual %d", expected[i].Sum(), actual[i].Sum())
		}
		if actual[i].Weight() != expected[i].Weight() {
			t.Errorf("feature.Weight(): expected %d, actual %d", expected[i].Weight(), actual[i].Weight())
		}
	}
}
