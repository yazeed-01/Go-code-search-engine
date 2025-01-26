package processors

import (
	"cse/models"
	"encoding/json"
	"regexp"
	"strings"
)

type MethodProcessor struct{}

type Parameter struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func (p *MethodProcessor) Process(content string) interface{} {
	methodRegex := regexp.MustCompile(`(?m)^(public|private|protected)?\s*(?:static\s+)?(\w+(?:<.*?>)?)\s+(\w+)\s*$$(.*?)$$`)

	var methods []models.Method
	lines := strings.Split(content, "\n")

	for lineNum, line := range lines {
		matches := methodRegex.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			// Get the full match details
			fullMatch := methodRegex.FindStringSubmatch(line)

			// Parse parameters
			var params []Parameter
			if fullMatch[4] != "" {
				paramStrings := strings.Split(fullMatch[4], ",")
				for _, param := range paramStrings {
					parts := strings.Fields(strings.TrimSpace(param))
					if len(parts) >= 2 {
						params = append(params, Parameter{
							Type: parts[0],
							Name: parts[1],
						})
					}
				}
			}

			// Convert parameters to JSON string
			paramsJSON, _ := json.Marshal(params)

			method := models.Method{
				Name:       fullMatch[3],
				ReturnType: fullMatch[2],
				Parameters: string(paramsJSON),
				Location: models.Location{
					Line:   lineNum + 1,
					Column: match[0] + 1,
				},
			}
			methods = append(methods, method)
		}
	}
	return methods
}
