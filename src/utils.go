package src

import (
	"os"
	"path/filepath"
)

func getWorkDirPath() (string, error) {
	binaryPath, err := os.Getwd()

	path := filepath.Join(binaryPath, "/encryption")
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
