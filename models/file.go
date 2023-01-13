package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Path string `json:"path"`
	FolderId int `json:"folder_id"`
}