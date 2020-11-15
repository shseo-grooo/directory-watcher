package helper

import (
	"log"
	"os"
)

func IsNotExist(path string) bool {
	return !IsExist(path)
}

func IsExist(path string) bool {
	_, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			log.Fatalln(err)
		}
	}
	return true
}

func IsDir(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	info, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	return info.IsDir()
}
