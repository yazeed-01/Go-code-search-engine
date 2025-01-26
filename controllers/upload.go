package controllers

import (
	"archive/zip"
	"cse/initializers"
	"cse/models"
	"cse/processors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleFileUpload(c *gin.Context) {
	// Receive the uploaded file
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// Create a new project
	project := models.Project{
		Name: header.Filename}

	if err := initializers.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project in database"})
		return
	}

	// Create a permanent directory to store the uploaded zip
	uploadDir := fmt.Sprintf("./uploads/project_%d", project.ID)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// Save the uploaded file with its original name
	zipPath := filepath.Join(uploadDir, header.Filename)
	out, err := os.Create(zipPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
		return
	}

	// Process the zip file
	invertedIndex, err := processZipFile(zipPath, uploadDir, project.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "File processed successfully",
		"project_id":     project.ID,
		"inverted_index": invertedIndex,
	})
}

func processZipFile(zipPath, extractDir string, projectID uint) (map[string]interface{}, error) {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open zip file: %v", err)
	}
	defer reader.Close()

	processorFactory := &processors.ProcessorFactory{}
	invertedIndex := make(map[string]interface{})

	for _, file := range reader.File {
		if filepath.Ext(file.Name) != ".java" {
			continue
		}

		filePath := filepath.Join(extractDir, file.Name)

		// Ensure the file path is inside the extract directory
		if !strings.HasPrefix(filePath, filepath.Clean(extractDir)+string(os.PathSeparator)) {
			return nil, fmt.Errorf("invalid file path: %s", filePath)
		}

		if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create directory: %v", err)
		}

		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %v", err)
		}

		rc, err := file.Open()
		if err != nil {
			outFile.Close()
			return nil, fmt.Errorf("failed to open file in zip: %v", err)
		}

		content, err := ioutil.ReadAll(rc)
		if err != nil {
			outFile.Close()
			rc.Close()
			return nil, fmt.Errorf("failed to read file content: %v", err)
		}

		_, err = outFile.Write(content)
		outFile.Close()
		rc.Close()

		if err != nil {
			return nil, fmt.Errorf("failed to write file: %v", err)
		}

		// Process file content
		fileIndex := processFileContent(string(content), processorFactory)

		dbFile := models.File{
			Name:          file.Name,
			ProjectID:     projectID,
			Classes:       fileIndex["classes"].([]models.Class),
			Methods:       fileIndex["methods"].([]models.Method),
			Variables:     fileIndex["variables"].([]models.Variable),
			Relationships: fileIndex["relationships"].([]models.Relationship),
		}

		if err := initializers.DB.Create(&dbFile).Error; err != nil {
			return nil, fmt.Errorf("failed to save file info to database: %v", err)
		}

		invertedIndex[file.Name] = fileIndex
	}

	return invertedIndex, nil
}

func processFileContent(content string, factory *processors.ProcessorFactory) map[string]interface{} {
	fileIndex := make(map[string]interface{})

	classProcessor := factory.CreateProcessor("class")
	methodProcessor := factory.CreateProcessor("method")
	variableProcessor := factory.CreateProcessor("variable")
	relationshipProcessor := factory.CreateProcessor("relationship")

	fileIndex["classes"] = classProcessor.Process(content)
	fileIndex["methods"] = methodProcessor.Process(content)
	fileIndex["variables"] = variableProcessor.Process(content)
	fileIndex["relationships"] = relationshipProcessor.Process(content)

	return fileIndex
}
