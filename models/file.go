package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name string `json:"name"`
	Path string `json:"path"`
	FolderId int `json:"folder_id"`
}