package cmd

import (
	"os"
)

func CreateDirectoryIfNotExists(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
	}

	if err != nil {
		panic(err)
	}
}
