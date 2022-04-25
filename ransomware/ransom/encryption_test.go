package ransom

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	TestData := []byte("Test Data")
	key := []byte("dO)CIKA)as90i1e2;zlckoasdjzix102")
	got := Encrypt(key, TestData)
	Decrypted := Decrypt(key, got)
	if !bytes.Equal(Decrypted, TestData) {
		t.Errorf("got %q, wanted %q", got, TestData)
	}
}

//faster if you have more extensions
func BenchmarkExtensionDict(b *testing.B) {
	test := make([]string, 100)
	listLen := len(extensions_list)
	for i := 0; i < 100; i++ {
		test[i] = extensions_list[rand.Intn(listLen)]
	}
	for i := 0; i < b.N; i++ {
		tmp := extension_dict[test[rand.Intn(listLen)]]
		_ = tmp
	}
}

//faster if you have under 100 extensions
func BenchmarkExtensionList(b *testing.B) {
	test := make([]string, 100)
	listLen := len(extensions_list) - 1
	for i := 0; i < 100; i++ {
		test[i] = extensions_list[rand.Intn(listLen)]
	}
	for i := 0; i < b.N; i++ {
		nr := rand.Intn(listLen) - 1
		for j := 0; j < nr; j++ {
			if j == len(test[j]) {
				tmp := extensions_list[nr]
				_ = tmp
			}
		}
	}
}

func TestEncryptDecryptFile(t *testing.T) {
	TestData := []byte("TestString")
	key := []byte("dO)CIKA)as90i1e2;zlckoasdjzix102")
	err := ioutil.WriteFile("test.txt", TestData, 0644)
	if err != nil {
		t.Error(err)
	}

	err = EncryptFile("test.txt", key)
	if err != nil {
		os.Remove("test.txt")
		t.Error(err)
	}

	err = DecryptFile("test.txt"+extension, key)
	if err != nil {
		os.Remove("test.txt" + extension)
		t.Error(err)
	}

	out, err := ioutil.ReadFile("test.txt")
	if err != nil {
		os.Remove("test.txt")
		t.Error(err)
	}

	if !bytes.Equal(out, TestData) {
		t.Error("Files do not match")
	}

	os.Remove("test.txt")
}

//TODO: Add tests for recursive encryption and decryption using randomly generated data
