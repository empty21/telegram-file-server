package model

import (
	"github.com/gofiber/utils/v2"
)

type File struct {
	Uuid     string `json:"uuid" gorm:"primaryKey"`
	FileId   string `json:"-"`
	Name     string `json:"name"`
	MimeType string `json:"mimeType"`
}

func (f *File) TableName() string {
	return "files"
}

func NewFile(fileId string, name string, mimeType string) *File {
	return &File{
		Uuid:     utils.UUIDv4(),
		FileId:   fileId,
		Name:     name,
		MimeType: mimeType,
	}
}
