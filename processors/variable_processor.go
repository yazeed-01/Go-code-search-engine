package processors

import (
	"cse/models"
	"fmt"
	"regexp"
	"strings"
)

type VariableProcessor struct{}

func (p *VariableProcessor) Process(content string) interface{} {
	variableRegex := regexp.MustCompile(`(?m)^(public|private|protected)?\s*(?:static\s+)?(\w+(?:<.*?>)?)\s+(\w+)\s*(?:=.*?)?;`)

	var variables []models.Variable
	lines := strings.Split(content, "\n")

	for lineNum, line := range lines {
		matches := variableRegex.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			// Get the full match details
			fullMatch := variableRegex.FindStringSubmatch(line)

			// Debugging output
			fmt.Printf("Line: %d, Match: %s, Column: %d\n", lineNum+1, fullMatch[0], match[0]+1)

			variable := models.Variable{
				Type: fullMatch[2],
				Name: fullMatch[3],
				Location: models.Location{
					Line:   lineNum + 1,
					Column: match[0] + 1,
				},
			}

			variables = append(variables, variable)
		}
	}
	return variables
}

func (p *VariableProcessor) AssociateWithClass(variables []models.Variable, classID uint) {
	for i := range variables {
		variables[i].ClassID = classID
	}
}

func (p *VariableProcessor) AssociateWithMethod(variables []models.Variable, methodID uint) {
	for i := range variables {
		variables[i].MethodID = methodID
	}
}
