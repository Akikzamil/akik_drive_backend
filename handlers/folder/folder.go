package folder

import (
	"akik_drive/config/database"
	"akik_drive/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateFolder(c echo.Context) error {
	folder := new(models.Folder)
	if err := c.Bind(&folder); err != nil {
		return err
	}
	result := database.DB.Where(&models.Folder{Name: folder.Name, ParentFolderID: folder.ParentFolderID}).First(&folder)
	if result.RowsAffected != 0 {
		return c.String(http.StatusFound,"This Name of already exists")
	}
	database.DB.Create(&folder)
	return c.JSON(http.StatusCreated, &folder)
}

func GetFolderByParentId(c echo.Context) error {
	var folders []models.Folder
	parentId := c.Param("id");if parentId == ""{
		database.DB.Where("parent_folder_id IS NULL").Find(&folders)
	}else{
		database.DB.Where("parent_folder_id=?", parentId).Find(&folders)
	}
	return c.JSON(http.StatusOK, folders)
}

func GetFolder(c echo.Context) error {
	var folder models.Folder
	id := c.Param("id");if id == ""{
		database.DB.Preload("Folders").Preload("Files").Where("parent_folder_id IS NULL").Find(&folder)
	}else{
		database.DB.Preload("Folders").Preload("Files").Where("ID=?", id).Find(&folder)
	}
	return c.JSON(http.StatusOK, folder)
}

func DeleteFolder(c echo.Context) error {
	id := c.Param("id")
	var folder models.Folder
	result := database.DB.Delete(&folder, id)
	if result.RowsAffected == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.String(http.StatusOK, "Successfully Deleted")
}

func ImportDefaultFolder() {
	var folders []models.Folder
	result := database.DB.Where("parent_folder_id IS NULL").First(&folders)
	if result.RowsAffected == 0 {
		// database.DB.Create(&models.Folder{Name: "Main", ParentFolderID: 0})
		database.DB.Exec("INSERT INTO folders (created_at,updated_at,deleted_at,name,parent_folder_id) VALUES ('2023-01-21 23:55:45.348','2023-01-21 23:55:45.348',NULL,'Main',NULL) RETURNING id")
	}
}
