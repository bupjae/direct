//go:build windows

package dwrite

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go factory.go
