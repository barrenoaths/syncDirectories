package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("--- Sync stuff ---")

	dirPath1 := "/home/pawel/learnGo"
	fileToHashPath := "/home/pawel/learnGo/syncDirectories/orchid.txt"

	if dirExists(dirPath1) {
		fmt.Println("Directory exists")
	}

	fmt.Println(getFileHash(fileToHashPath))
	fmt.Println(fileExist(fileToHashPath))
}

func dirExists(absPath string) bool {
	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func fileExist(absPath string) bool {
	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}
