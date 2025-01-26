package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleFileUpload(c *gin.Context) {
	// Get the file from the form input
	file, err := c.FormFile("file")
	if err != nil || file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No file uploaded"})
		return
	}

	// Save the file locally (you can change the path to where you want to save)
	dst := fmt.Sprintf("./uploads/%s", file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to save file"})
		return
	}

	// Send a success response
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "File uploaded successfully"})
}

