package implementations

import (
	"errors"
	"io/ioutil"
	"ismaeldf/golang-gobarber/config"
	"mime/multipart"
	"os"
)

type DiskStorageProvider struct{}

func (d *DiskStorageProvider) SaveFile(file multipart.File) (string, error) {
	tempFile, err := ioutil.TempFile(config.FileDirectory, "avatar-*.png")
	if err != nil {
		return "", errors.New(err.Error())
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.New(err.Error())
	}
	tempFile.Write(fileBytes)

	return tempFile.Name(), nil
}

func (d *DiskStorageProvider) DeleteFile(filename string){
	_ = os.Remove(config.FileDirectory + filename)
}
