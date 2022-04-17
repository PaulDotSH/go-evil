package ransom

import (
	"crypto/rand"
	"strings"
)

func GenerateKey() string {
	const size = byte(len(charset) - 1)
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
