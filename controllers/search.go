package controllers

import (
	"cse/initializers"
	"cse/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "search.html", gin.H{
		"title": "Code Search Engine",
	})
}

func Search(c *gin.Context) {
	var searchRequest struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	}

	if err := c.ShouldBindJSON(&searchRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := searchRequest.Query
	searchType := searchRequest.Type

	var results []map[string]interface{}

	if !isValidSearchType(searchType) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search type"})
		return
	}

	switch searchType {
	case "method":
		var methods []models.Method
		initializers.DB.Where("name LIKE ?", "%"+query+"%").
			Preload("Class.File.Project").
			Find(&methods)
		for _, method := range methods {
			results = append(results, map[string]interface{}{
				"type":      "method",
				"name":      method.Name,
				"class":     method.Class.Name,
				"file":      method.Class.File.Name,
				"project":   method.Class.File.Project.Name,
				"projectID": method.Class.File.ProjectID,
				"line":      method.Location.Line,
				"column":    method.Location.Column,
			})
		}

	case "class":
		var classes []models.Class
		initializers.DB.Where("name LIKE ?", "%"+query+"%").
			Preload("File.Project").
			Find(&classes)
		for _, class := range classes {
			results = append(results, map[string]interface{}{
				"type":      "class",
				"name":      class.Name,
				"file":      class.File.Name,
				"project":   class.File.Project.Name,
				"projectID": class.File.ProjectID,
				"line":      class.Location.Line,
				"column":    class.Location.Column,
			})
		}

	case "file":
		var files []models.File
		initializers.DB.Where("name LIKE ?", "%"+query+"%").
			Preload("Project").
			Find(&files)
		for _, file := range files {
			results = append(results, map[string]interface{}{
				"type":      "file",
				"name":      file.Name,
				"project":   file.Project.Name,
				"projectID": file.ProjectID,
			})
		}

	case "variable":
		var variables []models.Variable
		initializers.DB.Where("name LIKE ? OR type LIKE ?", "%"+query+"%", "%"+query+"%").
			Preload("File.Project").
			Preload("Class").
			Find(&variables)
		for _, variable := range variables {
			result := map[string]interface{}{
				"type":      "variable",
				"name":      variable.Name,
				"varType":   variable.Type,
				"file":      variable.File.Name,
				"project":   variable.File.Project.Name,
				"projectID": variable.File.ProjectID,
				"line":      variable.Location.Line,
				"column":    variable.Location.Column,
			}
			if variable.ClassID != 0 {
				//result["class"] = variable.Class.Name
			}
			results = append(results, result)
		}

	case "relationship":
		var relationships []models.Relationship
		initializers.DB.Where(
			"source_class_name LIKE ? OR target_class_name LIKE ? OR relationship_type LIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%",
		).
			Preload("File.Project").
			Find(&relationships)
		for _, rel := range relationships {
			results = append(results, map[string]interface{}{
				"type":             "relationship",
				"sourceClass":      rel.SourceClassName,
				"targetClass":      rel.TargetClassName,
				"relationshipType": rel.RelationshipType,
				"file":             rel.File.Name,
				"project":          rel.File.Project.Name,
				"projectID":        rel.File.ProjectID,
				"line":             rel.Location.Line,
				"column":           rel.Location.Column,
			})
		}

	case "text":
		var methods []models.Method
		var classes []models.Class
		var files []models.File
		var variables []models.Variable
		var relationships []models.Relationship

		initializers.DB.Where("name LIKE ?", "%"+query+"%").
			Preload("Class.File.Project").
			Find(&methods)

		initializers.DB.Where("name LIKE ?", "%"+query+"%").
			Preload("File.Project").
			Find(&classes)

		initializers.DB.Where("name LIKE ?", "%"+query+"%").
			Preload("Project").
			Find(&files)

		initializers.DB.Where("name LIKE ? OR type LIKE ?", "%"+query+"%", "%"+query+"%").
			Preload("File.Project").
			Preload("Class").
			Find(&variables)

		initializers.DB.Where(
			"source_class_name LIKE ? OR target_class_name LIKE ? OR relationship_type LIKE ?",
			"%"+query+"%", "%"+query+"%", "%"+query+"%",
		).
			Preload("File.Project").
			Find(&relationships)

		appendResults(&results, methods, "method")
		appendResults(&results, classes, "class")
		appendResults(&results, files, "file")
		appendResults(&results, variables, "variable")
		appendResults(&results, relationships, "relationship")
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

func isValidSearchType(searchType string) bool {
	validTypes := []string{"method", "class", "file", "text", "variable", "relationship"}
	for _, validType := range validTypes {
		if searchType == validType {
			return true
		}
	}
	return false
}

func appendResults(results *[]map[string]interface{}, items interface{}, itemType string) {
	switch itemType {
	case "method":
		for _, method := range items.([]models.Method) {
			*results = append(*results, map[string]interface{}{
				"type":      "method",
				"name":      method.Name,
				"class":     method.Class.Name,
				"file":      method.Class.File.Name,
				"project":   method.Class.File.Project.Name,
				"projectID": method.Class.File.ProjectID,
				"line":      method.Location.Line,
				"column":    method.Location.Column,
			})
		}
	case "class":
		for _, class := range items.([]models.Class) {
			*results = append(*results, map[string]interface{}{
				"type":      "class",
				"name":      class.Name,
				"file":      class.File.Name,
				"project":   class.File.Project.Name,
				"projectID": class.File.ProjectID,
				"line":      class.Location.Line,
				"column":    class.Location.Column,
			})
		}
	case "file":
		for _, file := range items.([]models.File) {
			*results = append(*results, map[string]interface{}{
				"type":      "file",
				"name":      file.Name,
				"project":   file.Project.Name,
				"projectID": file.ProjectID,
			})
		}
	case "variable":
		for _, variable := range items.([]models.Variable) {
			result := map[string]interface{}{
				"type":      "variable",
				"name":      variable.Name,
				"varType":   variable.Type,
				"file":      variable.File.Name,
				"project":   variable.File.Project.Name,
				"projectID": variable.File.ProjectID,
				"line":      variable.Location.Line,
				"column":    variable.Location.Column,
			}
			if variable.ClassID != 0 {
				//result["class"] = variable.Class.Name
			}
			*results = append(*results, result)
		}
	case "relationship":
		for _, rel := range items.([]models.Relationship) {
			*results = append(*results, map[string]interface{}{
				"type":             "relationship",
				"sourceClass":      rel.SourceClassName,
				"targetClass":      rel.TargetClassName,
				"relationshipType": rel.RelationshipType,
				"file":             rel.File.Name,
				"project":          rel.File.Project.Name,
				"projectID":        rel.File.ProjectID,
				"line":             rel.Location.Line,
				"column":           rel.Location.Column,
			})
		}
	}
}
