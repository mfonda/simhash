// Copyright 2013 Matthew Fonda. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simhash

import (
	"testing"
)

func TestFingerprint(t *testing.T) {
	var v [5]Vector
	for i := 0; i < 64; i++ {
		v[0][i] = 0
		v[1][i] = 1
		v[2][i] = -1
		v[3][i] = -1
		if i%2 == 0 {
			v[4][i] = 1
		} else {
			v[4][i] = -1
		}
	}
	v[3][0] = 1
	expected := []uint64{0xffffffffffffffff, 0xffffffffffffffff, 0, 1, 6148914691236517205}
	for i := 0; i < len(v); i++ {
		actual := Fingerprint(v[i])
		if actual != expected[i] {
			t.Errorf("Fingerprint(%v): expected %d, actual %d", v[i], expected[i], actual)
		}
	}
}
