package repository

import (
	"gorm.io/gorm"
	"telegram-file-server/pkg/model"
)

type FileRepository interface {
	table() *gorm.DB
	model() *model.File
	Save(file *model.File) error
	FindById(id string) (*model.File, error)
	Delete(id string) error
}

type fileRepository struct {
	db *gorm.DB
}

func (f fileRepository) table() *gorm.DB {
	return f.db.Table(f.model().TableName())
}

func (f fileRepository) model() *model.File {
	return &model.File{}
}

func (f fileRepository) Save(file *model.File) error {
	trx := f.table().Save(file)
	return trx.Error
}

func (f fileRepository) FindById(id string) (*model.File, error) {
	file := f.model()
	trx := f.table().Where("uuid = ?", id).First(file)
	return file, trx.Error
}

func (f fileRepository) Delete(id string) error {
	trx := f.table().Where("uuid = ?", id).Delete(f.model())
	return trx.Error
}

func NewFileRepository(db *gorm.DB) FileRepository {
	r := &fileRepository{db: db}
	_ = db.AutoMigrate(r.model())
	return r
}
