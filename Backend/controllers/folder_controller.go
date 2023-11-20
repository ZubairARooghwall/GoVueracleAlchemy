package controllers

import (
	"net/http"
	"strconv"

	"github.com/ZubairARooghwall/GoVueracleAlchemy/models"
	"github.com/ZubairARooghwall/GoVueracleAlchemy/repository"
	"github.com/gin-gonic/gin"
)

type FolderController struct {
	FolderRepo *repository.FolderRepository
}

func NewFolderController(folderRepo *repository.FolderRepository) *FolderController {
	return &FolderController{FolderRepo: folderRepo}
}

func (fc *FolderController) CreateFolder(c *gin.Context) {
	var folder models.Folder
	if err := c.ShouldBindJSON(&folder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := fc.FolderRepo.CreateFolder(folder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, folder)
}

func (fc *FolderController) GetFolderByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	folders, err := fc.FolderRepo.GetFoldersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, folders)
}

func (fc *FolderController) DeleteFolder(c *gin.Context) {
	folderID, err := strconv.Atoi(c.Param("folderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid folder ID"})
		return
	}

	if err := fc.FolderRepo.DeleteFolder(folderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Folder deleted successfully"})
}
