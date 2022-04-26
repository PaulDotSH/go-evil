# GO EVIL

Go Evil is a collection of malware-like tools that work with each other, using the "evil" backend.

# Features

- ## General
- Cross Platform
- "Cheap" to host
- Easy to setup

- ## Ransomware
- Encrypts all files it can from the computer
- Uses a random salt
- Uses a cryptographically secure random generator (files CANNOT be decrypted without the key)
- Customizable extension and list of extensions to encrypt
- Files are decrpytable with the same binary (./file decrypt $key)
- Generates a routine for all drives (encrypts multiple files at the same time)
- Much more

- ## Stealer

## License
All of these tools which are written by me are licensed under AGPL 3.0
However, any tools I integrate from other projects maintain their original license unless they are modified and the license permits so.

# DISCLAIMER
## YOU SHOULD NOT USE ANY OF THESE TOOLS FOR ILLEGAL PURPOSES

# SETUP
Tutorial Comming soon

# HOW TO BUILD
Run the command below in any of the folders to compile the software
```go build .```

# TODO
- Add auto startup for apps (separate package)
- Complete the project
- Check user idle time (separate package)
- Add backend
- Add package to change user's wallpaper
- Add persistence for windows (bsod on process exit)

# SCANS
The malware should start to get detected as soon as it gets more popular...
Until then, have fun pentesting :D

### Ransomware
![FUD Ransomware Windows 0/26](https://antiscan.me/images/result/DUBMelFlJ6ql.png)
