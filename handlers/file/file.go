package file

import (
	"akik_drive/config/database"
	"akik_drive/models"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)


func UploadAFile(c echo.Context) error {
	folderId := c.FormValue("folder_id");
	val ,_ := generateOTP(6);
	folderIdInt, _ := strconv.Atoi(folderId)
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(fmt.Sprintf("assets/%s%s",val, file.Filename))
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	

	fileData := models.File{Name: file.Filename,FolderId: folderIdInt,Path: fmt.Sprintf("/%s%s",val,file.Filename)}
	database.DB.Create(&fileData)
	return c.JSON(201,fileData)
}

func GetFilesById(c echo.Context) error {
	var files []models.File
	id := c.Param("id")
	database.DB.Where("folder_id = ?", id).Find(&files)
	return c.JSON(200,files)
}

func DeleteAFile(c echo.Context) error {
	id := c.Param("id")
	var file models.File
	database.DB.Where("id=?",id).First(&file)
	err := os.Remove(fmt.Sprintf("assets/%s",file.Path))
	if(err != nil){
		// return c.String(405,err.Error())
		fmt.Println(err)
	}
	result := database.DB.Delete(&file, id)
	if result.RowsAffected == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.String(http.StatusOK, "Successfully Deleted")
}

func DeleteFilesByFolderId(id int){
	var files []models.File
	database.DB.Where("folder_id=?",id).Find(&files)
	for _,file := range files{
		os.Remove(fmt.Sprintf("assets/%s",file.Path))
	}
	database.DB.Where("folder_id=?",id).Delete(&files)
}

func GetFile(c echo.Context) error {
	id := c.Param("id");
	var fileModel models.File
	database.DB.First(&fileModel,id)
	file,err := os.Open(fmt.Sprintf("assets/%s",fileModel.Path))
	if err != nil{
		return c.JSON(405,"invalid file Data")
	}
	defer file.Close()

	return c.Stream(http.StatusOK,"video/mp4",file)
}


const otpChars = "1234567890"

func generateOTP(length int) (string, error) {
    buffer := make([]byte, length)
    _, err := rand.Read(buffer)
    if err != nil {
        return "", err
    }

    otpCharsLength := len(otpChars)
    for i := 0; i < length; i++ {
        buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
    }

    return string(buffer), nil
}


func UploadFiles(c echo.Context) error {
	folderId := c.FormValue("folder_id");
	val ,_ := generateOTP(6);
	folderIdInt, _ := strconv.Atoi(folderId)
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	 filesJson := []models.File{} 

	for _,file := range files{
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
	
		// Destination
		dst, err := os.Create(fmt.Sprintf("assets/%s%s",val, file.Filename))
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	
		fileData := models.File{Name: file.Filename,FolderId: folderIdInt,Path: fmt.Sprintf("/%s%s",val,file.Filename)}
		filesJson = append(filesJson, fileData)
	}
	database.DB.Create(&filesJson)
	return c.JSON(201,filesJson)
}