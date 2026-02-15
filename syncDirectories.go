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

	dirPath1 := "/opt/learnGo"
	fileToHashPath := "/opt/learnGo/syncDirectories/orchid.txt"

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

func printTheContentsOfDir(dirPath string) {
	fmt.Println("placeholder")
}

func printTheContentsOfFile(filePath string) {
	fmt.Println("placeholder")
}

func copyFile() {
	fmt.Println("placeholder")
}

func copyDir() {
	fmt.Println("placeholder")
}

func copyDirWithFiles() {
	fmt.Println("placeholder")
}

func findLineInFile() {
	fmt.Println("placeholder")
}

func findPhraseInFile() {
	fmt.Println("placeholder")
}

func getFileMetadata() {
	fmt.Println("placeholder")
}

func modifyFileMetadata() {
	fmt.Println("placeholder")
}
