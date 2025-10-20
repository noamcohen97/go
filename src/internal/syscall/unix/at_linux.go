// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package unix

import (
	"syscall"
	"unsafe"
)

type OpenHow struct {
	Flags   uint64
	Mode    uint64
	Resolve uint64
}

func Openat2(dirfd int, path string, how *OpenHow) (int, error) {
	p, err := syscall.BytePtrFromString(path)
	if err != nil {
		return 0, err
	}

	fd, _, errno := syscall.Syscall6(openat2Trap, uintptr(dirfd), uintptr(unsafe.Pointer(p)), uintptr(unsafe.Pointer(how)), uintptr(unsafe.Sizeof(OpenHow{})), 0, 0)
	if errno != 0 {
		return 0, errno
	}

	return int(fd), nil
}
