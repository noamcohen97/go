// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

package os

import (
	"internal/syscall/unix"
	"syscall"
)

func init() {
	rootOpenFileNologFast = rootOpenFileNologNoFollowAny
}

func rootOpenFileNologNoFollowAny(root *Root, name string, flag int, perm FileMode) (int, error) {
	const O_NOFOLLOW_ANY = 0x20000000
	fd, err := unix.Openat(root.root.fd, name, syscall.O_NOFOLLOW|syscall.O_CLOEXEC|O_NOFOLLOW_ANY|flag, uint32(perm))
	if isNoFollowErr(err) {
		// Fall back to the slow path.
		return 0, syscall.EAGAIN
	}
	return fd, err
}
