package service

import (
	"os"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	Check(err)
	return file

}

func ReadFile(file os.File) []byte {
	data := make([]byte, 200)
	file.Read(data)
	return data
}
