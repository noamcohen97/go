// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package os

import (
	"internal/syscall/unix"
)

func init() {
	if unix.KernelVersionGE(5, 6) {
		rootOpenFileNologFast = rootOpenFileNologOpenat2
	}
}

func rootOpenFileNologOpenat2(root *Root, name string, flag int, perm FileMode) (int, error) {
	return unix.Openat2(root.root.fd, name, &unix.OpenHow{
		Flags: uint64(flag),
		Mode:  uint64(perm),
		// We use unix.RESOLVE_BENEATH, which currently includes unix.RESOLVE_NO_SYMLINKS.
		// To ensure future compatibility, we explicitly add unix.RESOLVE_NO_MAGICLINKS.
		Resolve: unix.RESOLVE_BENEATH | unix.RESOLVE_NO_MAGICLINKS,
	})
}
