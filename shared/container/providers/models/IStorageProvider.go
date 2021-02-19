package models

import "mime/multipart"

type IStorageProvider interface {
	SaveFile(file multipart.File) (string, error)
	DeleteFile(filename string)
}
