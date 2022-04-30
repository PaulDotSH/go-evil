package ransomware

import (
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkGetExtension(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filepath.Walk("/run/media/", func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			_ = GetFileExtension(path)
			//fmt.Println(GetFileExtension(path))
			return nil
		})
	}
}

//the difference becomes apparent as the paths are longer and you test with more files
func BenchmarkGetExtensionFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filepath.Walk("/run/media/", func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			_ = GetFileExtensionFast(path)
			//fmt.Println(GetFileExtension(path))
			return nil
		})
	}
}

func BenchmarkGetExtensionFastest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filepath.Walk("/run/media/", func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			_ = GetFileExtensionFast(path)
			//fmt.Println(GetFileExtension(path))
			return nil
		})
	}
}
