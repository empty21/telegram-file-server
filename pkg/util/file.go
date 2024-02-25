package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"
	"telegram-file-server/pkg/model"
)

const joinSeparator = ":"

var encoding = base64.URLEncoding

type fileUtil struct {
	cipher cipher.Block
}

type FileUtil interface {
	EncryptFile(file *model.File) error
	DecryptFile(file *model.File) error
}

func (f fileUtil) fileJoin(file model.File) string {
	return strings.Join([]string{file.FileId, file.Name, file.MimeType}, joinSeparator)
}

func (f fileUtil) fileSplit(encryption string) (*model.File, error) {
	split := strings.Split(encryption, joinSeparator)
	if len(split) != 3 {
		return nil, fmt.Errorf("invalid encryption")
	}
	return &model.File{
		Id:       encryption,
		FileId:   split[0],
		Name:     split[1],
		MimeType: split[2],
	}, nil
}

func (f fileUtil) EncryptFile(file *model.File) error {
	b := []byte(f.fileJoin(*file))
	cipherText := make([]byte, aes.BlockSize+len(b))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(f.cipher, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], b)

	file.Id = encoding.EncodeToString(cipherText)

	return nil
}

func (f fileUtil) DecryptFile(file *model.File) error {
	cipherText, err := encoding.DecodeString(file.Id)
	if err != nil {
		return err
	}

	if len(cipherText) < aes.BlockSize {
		return errors.New("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(f.cipher, iv)
	stream.XORKeyStream(cipherText, cipherText)

	_file, err := f.fileSplit(string(cipherText))
	if err == nil {
		file.Name = _file.Name
		file.FileId = _file.FileId
		file.MimeType = _file.MimeType
	}

	return err
}

func NewFileUtil(secret string) FileUtil {
	if len(secret) > 32 {
		secret = secret[:32]
	} else if len(secret) > 24 {
		secret = secret[:24]
	} else if len(secret) > 16 {
		secret = secret[:16]
	}
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return nil
	}
	return fileUtil{
		cipher: block,
	}
}
