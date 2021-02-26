package implementations

import (
	"fmt"
	"io/ioutil"
	"ismaeldf/golang-gobarber/config"
	"mime/multipart"
	"os"
)

type DiskStorageProvider struct{}

func (d *DiskStorageProvider) SaveFile(file multipart.File) string {
	tempFile, err := ioutil.TempFile(config.FileDirectory, "avatar-*.png")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(err.Error())
	}
	tempFile.Write(fileBytes)

	return tempFile.Name()
}

func (d *DiskStorageProvider) DeleteFile(filename string){
	_ = os.Remove(config.FileDirectory + filename)
}
