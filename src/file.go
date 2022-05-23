package src

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
)

type File struct {
	fileID string `json:"-"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Size   int64  `json:"size"`
}

type FileMap map[string]File
type FileList []File

var metaLock sync.Mutex

func ListAllFiles(password string) (FileList, error) {
	data, err := readFile(METADATA_PATH)
	if os.IsNotExist(err) {
		err = SaveMeta([]File{}, password)
		return []File{}, nil
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
	metaLock.Lock()
	defer metaLock.Unlock()
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
	filePath := filepath.Clean(filepath.Join(file.Path, file.Name))

	for _, f := range files {
		path := filepath.Clean(filepath.Join(f.Path, f.Name))
		if path == filePath {
			return fmt.Errorf("file already exists - %v", filePath)
		}
	}

	files = append(files, file)
	return SaveMeta(files, password)
}

func deleteFileFromMeta(path, password string, recursive bool) ([]File, error) {
	files, err := ListPath(path, password)
	if err != nil {
		return []File{}, err
	}

	if !recursive && len(files) > 1 {
		err = fmt.Errorf("more than one file exists with given path")
	}

	return files, err
}

func NewFile(data []byte, file File, password string) error {
	fileId := uuid.New().String()
	file.fileID = fileId
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

	return writeFile(filepath.Join(ENCRYPTED_FILES_DIR, file.fileID), data)
}

func NewFileWithPath(srcPath, destPath, password string) error {
	fileId := uuid.NewString()

	wd, err := getWorkDirPath()
	if err != nil {
		return err
	}

	size, err := CryptPath(srcPath, filepath.Join(wd, ENCRYPTED_FILES_DIR, fileId), password, false)
	if err != nil {
		return err
	}

	var file File
	file.Size = size
	file.Path = filepath.Dir(destPath)
	file.Name = filepath.Base(destPath) // TODO: Check whether destPath is dir or filename

	return addFileToMeta(file, password)
}

func ListPath(path, password string) ([]File, error) {
	path = filepath.Clean(path)
	files, err := ListAllFiles(password)
	if err != nil {
		return []File{}, err
	}

	var toReturn []File
	for _, v := range files {
		if filepath.Clean(filepath.Join(v.Path, v.Name)) == path || filepath.Clean(v.Path) == path {
			toReturn = append(toReturn, v)
		}
	}

	return toReturn, nil
}

func DeleteFile(path, password string, recursive bool) error {
	files, err := deleteFileFromMeta(path, password, recursive)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("no such file found")
	}

	wd, err := getWorkDirPath()
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.Remove(filepath.Join(wd, WORKING_PATH, ENCRYPTED_FILES_DIR, file.fileID))
		if err != nil {
			return err
		}
	}

	return nil
}
