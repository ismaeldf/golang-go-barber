package models

import "mime/multipart"

type IStorageProvider interface {
	SaveFile(file multipart.File) string
	DeleteFile(filename string)
}
