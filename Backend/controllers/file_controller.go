package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ZubairARooghwall/GoVueracleAlchemy/repository"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	FileRepo repository.FileRepository
}

func NewFileController(fileRepo repository.FileRepository) *FileController {
	return &FileController{
		FileRepo: fileRepo,
	}
}

func (fc *FileController) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve the file from the request"})
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open the file"})
		return
	}
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
		return
	}

	userID := getUserIDFromContext(c)

	err = fc.FileRepo.SaveFile(fileBytes, file.Filename, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func (fc *FileController) GetAllFilesByUserID(c *gin.Context) {
	userID := getUserIDFromContext(c)

	files, err := fc.FileRepo.GetAllFilesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch files"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}

func (fc *FileController) GetFileByID(c *gin.Context) {
	fileID := c.Param("id")

	fileIDInt, err := strconv.Atoi(fileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID, Not Int"})
		return
	}

	file, err := fc.FileRepo.GetFileByID(fileIDInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file": file})
}

func (fc *FileController) DeleteFileByID(c *gin.Context) {
	fileID := c.Param("id")

	err := fc.FileRepo.DeleteFile(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Filed to delete file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})

}

func (fc *FileController) SaveFile(c *gin.Context) {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	c.Request.ParseMultipartForm(8 << 20)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	userID := getUserIDFromContext(c)

	if file.Size > (8 << 20) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the limit"})
		return
	}

	fileBytes, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	fileContent, err := ioutil.ReadAll(fileBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
		return
	}

	err = fc.FileRepo.SaveFile(fileContent, file.Filename, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File saved successfully"})
}

func getUserIDFromContext(c *gin.Context) int {
	userID, exists := c.Get("userID")
	if !exists {
		fmt.Errorf("username does not exist")
		return 0
	}

	return userID.(int)
}
