package fakes

import (
	"math/rand"
	"mime/multipart"
	"strconv"
)

type FakeStorageProvider struct{
	files[] string
}

func (d *FakeStorageProvider) SaveFile(file multipart.File) string {
	var filename = "avatar-file-" + strconv.Itoa(rand.Intn(1000))
	d.files = append(d.files, filename)
	return filename
}

func (d *FakeStorageProvider) DeleteFile(filename string){
	var index = -1
	for i, f := range d.files {
		if f == filename {
			index = i
		}
	}
	d.files = append(d.files[:index], d.files[index+1:]...)
}
