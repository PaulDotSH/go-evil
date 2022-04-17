package ransom

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	TestData := []byte("Test Data")
	Key := []byte("dO)CIKA)as90i1e2;zlckoasdjzix102")
	got := Encrypt(Key, TestData)
	Decrypted := Decrypt(Key, got)
	if !bytes.Equal(Decrypted, TestData) {
		t.Errorf("got %q, wanted %q", got, TestData)
	}
}

func TestEncryptDecryptFile(t *testing.T) {
	TestData := []byte("TestString")
	Key := []byte("dO)CIKA)as90i1e2;zlckoasdjzix102")
	err := ioutil.WriteFile("test.txt", TestData, 0644)
	if err != nil {
		t.Error(err)
	}

	err = EncryptFile("test.txt", Key)
	if err != nil {
		os.Remove("test.txt")
		t.Error(err)
	}

	err = DecryptFile("test.txt"+extension, Key)
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
