// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import "syscall"

const (
	unlinkatTrap uintptr = syscall.SYS_UNLINKAT
	openatTrap   uintptr = syscall.SYS_OPENAT
	// openat2Trap    uintptr = syscall.SYS_OPENAT2
	readlinkatTrap uintptr = syscall.SYS_READLINKAT
	mkdiratTrap    uintptr = syscall.SYS_MKDIRAT
	fchmodatTrap   uintptr = syscall.SYS_FCHMODAT
	fchownatTrap   uintptr = syscall.SYS_FCHOWNAT
	linkatTrap     uintptr = syscall.SYS_LINKAT
	symlinkatTrap  uintptr = syscall.SYS_SYMLINKAT
)

const (
	AT_EACCESS          = 0x200
	AT_FDCWD            = -0x64
	AT_REMOVEDIR        = 0x200
	AT_SYMLINK_NOFOLLOW = 0x100

	RESOLVE_BENEATH       = 0x8
	RESOLVE_IN_ROOT       = 0x10
	RESOLVE_NO_MAGICLINKS = 0x2
	RESOLVE_NO_SYMLINKS   = 0x4
	RESOLVE_NO_XDEV       = 0x1

	UTIME_OMIT = 0x3ffffffe
)
