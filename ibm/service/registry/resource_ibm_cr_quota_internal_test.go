// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package registry

import "testing"

func TestBytesToMegabytes(t *testing.T) {
	cases := map[int64]int64{-1: -1, 0: 0, 524288000: 500, 5368709120: 5120}
	for in, want := range cases {
		if got := bytesToMegabytes(in); got != want {
			t.Errorf("bytesToMegabytes(%d) = %d, want %d", in, got, want)
		}
	}
}
