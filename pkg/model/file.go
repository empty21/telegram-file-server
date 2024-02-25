package model

type File struct {
	Id       string `json:"id"`
	FileId   string `json:"-"`
	Name     string `json:"name"`
	MimeType string `json:"mimeType"`
}

func NewFile(fileId string, name string, mimeType string) *File {
	return &File{
		FileId:   fileId,
		Name:     name,
		MimeType: mimeType,
	}
}

func NewEncryptedFile(id string) *File {
	return &File{
		Id: id,
	}
}
