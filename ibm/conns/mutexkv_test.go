// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package conns

import (
	"testing"
	"time"
)

func TestMutexKVLock(t *testing.T) {
	mkv := NewMutexKV()

	mkv.Lock("foo")

	doneCh := make(chan struct{})

	go func() {
		mkv.Lock("foo")
		close(doneCh)
	}()

	select {
	case <-doneCh:
		t.Fatal("Second lock was able to be taken. This shouldn't happen.")
	case <-time.After(50 * time.Millisecond):
		// pass
	}
}

func TestMutexKVUnlock(t *testing.T) {
	mkv := NewMutexKV()

	mkv.Lock("foo")
	mkv.Unlock("foo")

	doneCh := make(chan struct{})

	go func() {
		mkv.Lock("foo")
		close(doneCh)
	}()

	select {
	case <-doneCh:
		// pass
	case <-time.After(50 * time.Millisecond):
		t.Fatal("Second lock blocked after unlock. This shouldn't happen.")
	}
}

func TestMutexKVDifferentKeys(t *testing.T) {
	mkv := NewMutexKV()

	mkv.Lock("foo")

	doneCh := make(chan struct{})

	go func() {
		mkv.Lock("bar")
		close(doneCh)
	}()

	select {
	case <-doneCh:
		// pass
	case <-time.After(50 * time.Millisecond):
		t.Fatal("Second lock on a different key blocked. This shouldn't happen.")
	}
}
