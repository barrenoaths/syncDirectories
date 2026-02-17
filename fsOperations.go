package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("--- Sync stuff ---")

	dirPath1 := "/opt/learnGo"
	fileToHashPath := "/opt/learnGo/fsOperations/orchid.txt"

	sourceFile := "/opt/learnGo/fsOperations/orchid.txt"
	destinationFile := "/opt/learnGo/fsOperations/orchid2.txt"

	dirToCopy := "/opt/learnGo/fsOperations/orchid"
	targetDirToCopy := "/opt/learnGo/fsOperations/copyOfOrchid/"

	if dirExists(dirPath1) {
		fmt.Printf("Directory %s exists\n", dirPath1)
	}

	if fileExist(fileToHashPath) {
		fmt.Printf("File %s exists\n", fileToHashPath)
	}

	wantedHash, _ := getFileHash(fileToHashPath)
	fmt.Printf("Hash of the %s is %s\n", fileToHashPath, wantedHash)

	fmt.Println("---------------------------")
	printTheContentsOfDir(dirPath1)
	fmt.Println("---------------------------")
	printTheContentsOfFile(fileToHashPath)
	fmt.Println("---------------------------")
	copyFile(sourceFile, destinationFile)

	copyDirWithFiles(dirToCopy, targetDirToCopy)
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
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory.", err)
		return
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

func printTheContentsOfFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file.", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
	}
}

func copyFile(srcPath, destPath string) error {
	sourceFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}

func copyDir() {
	fmt.Println("placeholder")
}

func copyDirWithFiles(srcDir, destDir string) error {
	return filepath.WalkDir(srcDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Construct the destination path
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(destDir, relPath)

		if d.IsDir() {
			// Create the directory in the destination
			return os.MkdirAll(dstPath, 0755)
		} else {
			// Copy the file
			return copyFile(path, dstPath)
		}
	})
}

func findLineInFile(filePath, phrase string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), phrase) {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
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
