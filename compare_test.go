// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhash

import (
	"testing"
)

type compareTest struct {
	a        uint64
	b        uint64
	expected uint8
}

var compareTests = []compareTest{
	{4, 3, 3}, {3, 4, 3},
	{4, 2, 2}, {2, 4, 2},
	{4, 1, 2}, {1, 4, 2},
	{4, 0, 1}, {0, 4, 1},
	{4, 4, 0},
}

func TestCompare(t *testing.T) {
	for _, tt := range compareTests {
		actual := Compare(tt.a, tt.b)
		if actual != tt.expected {
			t.Errorf("Compare(%d, %d): expected %d, actual %d", tt.a, tt.b, tt.expected, actual)
		}
	}
}
