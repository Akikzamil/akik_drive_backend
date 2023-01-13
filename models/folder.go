package models

import "gorm.io/gorm"

type Folder struct {
	gorm.Model
	Name string `json:"name"`
	ParentFolderID int `json:"parent_folder_id"`
	Folders []Folder `json:"folders" gorm:"foreignKey:ParentFolderID"`
	Files []File `json:"files" gorm:"foreignKey:FolderId"`
}