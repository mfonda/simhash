// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhash

import (
	"fmt"
	"testing"
)

func TestSimhash(t *testing.T) {
	var fp = []uint64{
		Simhash(&WordFeatureSet{[]byte("this is a test phrase")}),
		Simhash(&WordFeatureSet{[]byte("this is a test phrass")}),
		Simhash(&WordFeatureSet{[]byte("foo bar")}),
	}

	if Compare(fp[0], fp[1]) != 2 {
		t.Errorf("Comparison failed")
	}

	if Compare(fp[0], fp[2]) != 29 {
		t.Errorf("Comparison failed")
	}
}

var words = [][]byte{
	[]byte("one"),
	[]byte("two"),
	[]byte("three"),
	[]byte("four"),
	[]byte("five"),
}

var shingleTests = []struct {
	words    [][]byte
	w        int
	expected [][]byte
}{
	{words, 1, words},
	{words, 2, [][]byte{[]byte("one two"), []byte("two three"), []byte("three four"), []byte("four five")}},
	{words, 3, [][]byte{[]byte("one two three"), []byte("two three four"), []byte("three four five")}},
	{words, 4, [][]byte{[]byte("one two three four"), []byte("two three four five")}},
	{words, 5, [][]byte{[]byte("one two three four five")}},
	{words, 6, [][]byte{[]byte("one two three four five")}},
}

func TestShingle(t *testing.T) {
	for _, tt := range shingleTests {
		actual := Shingle(tt.w, tt.words)
		if !equal(actual, tt.expected) {
			t.Errorf("Shingle(%d, %v): expected %v, got %v", tt.w, tt.words, tt.expected, actual)
		}
	}
}

// Checks of two given [][]byte are equal
// TODO: is there a better way to do this?
func equal(a, b [][]byte) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}
