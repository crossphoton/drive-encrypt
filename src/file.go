package src

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type File struct {
	Name string
	Type string
	Size int64
}

type FileMap map[string]File
type FileList []File

func ListAllFiles(password string) (FileList, error) {
	data, err := readFile(METADATA_PATH)
	if os.IsNotExist(err) {
		err = SaveMeta([]File{}, password)
		return ListAllFiles(password)
	}

	if err != nil {
		return []File{}, err
	}

	metadata, err := decrypt(data, password)
	if err != nil {
		return []File{}, err
	}

	var files []File

	err = json.Unmarshal(metadata, &files)
	if err != nil {
		return []File{}, err
	}
	return files, nil
}

func SaveMeta(files []File, password string) error {
	jsondata, err := json.Marshal(files)
	if err != nil {
		return err
	}

	metadata, err := encrypt(jsondata, password)
	if err != nil {
		return err
	}

	return writeFile(METADATA_PATH, metadata)
}

func addFileToMeta(file File, password string) error {
	files, err := ListAllFiles(password)
	if err != nil {
		return err
	}

	files = append(files, file)
	return SaveMeta(files, password)
}

func NewFile(data []byte, file File, password string) error {
	fileId := uuid.New().String()
	data, err := encrypt(data, password)
	if err != nil {
		return err
	}

	err = addFileToMeta(file, password)
	if err != nil {
		return err
	}

	wd, err := getWorkDirPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(wd, ENCRYPTED_FILES_DIR), os.ModePerm)
	if err != nil {
		return err
	}

	return writeFile(filepath.Join(ENCRYPTED_FILES_DIR, fileId), data)
}
