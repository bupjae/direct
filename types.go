//go:build windows

package direct

import "syscall"

type (
	HResult = syscall.Errno
	IID     = syscall.GUID
)

type VTable [1024]uintptr
