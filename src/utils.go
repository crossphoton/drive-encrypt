package src

import (
	"log"
	"os"
	"path/filepath"
)

func getWorkDirPath() (string, error) {
	binaryPath, err := os.Getwd()

	path := filepath.Join(binaryPath, WORKING_PATH)
	if err != nil {
		return "", err
	}

	return path, nil
}

func writeFile(path string, data []byte) error {
	workingPath, err := getWorkDirPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(workingPath, os.ModePerm)
	if err != nil {
		return err
	}

	workingPath = filepath.Join(workingPath, path)
	os.Remove(workingPath)
	return os.WriteFile(workingPath, data, 0444)
}

func readFile(path string) ([]byte, error) {
	workingPath, err := getWorkDirPath()
	if err != nil {
		return nil, err
	}

	return os.ReadFile(filepath.Join(workingPath, path))
}

func init() {
	workingPath, err := getWorkDirPath()
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(workingPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(filepath.Join(workingPath, ENCRYPTED_FILES_DIR), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
