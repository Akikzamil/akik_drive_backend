package routes

import (
	"akik_drive/handlers/file"
	"akik_drive/handlers/folder"

	"github.com/labstack/echo/v4"
)

func SetupAuthorizedRoutes(e *echo.Group) {
	e.GET("/folders/:id",folder.GetFolderByParentId)
	e.GET("/folders",folder.GetFolderByParentId)
	e.GET("/folder/:id",folder.GetFolder)
	e.GET("/folder",folder.GetFolder)
	e.POST("/folder",folder.CreateFolder)
	e.DELETE("/folder/:id",folder.DeleteFolder)
	e.POST("/file",file.UploadAFile)
	e.POST("/files",file.UploadFiles)
	e.GET("/file/:id",file.GetFilesById)
	e.DELETE("/file/:id",file.DeleteAFile)
	e.GET("/",Welcome)
	e.Static("/static", "assets")
}