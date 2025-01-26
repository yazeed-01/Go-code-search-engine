package processors

import (
	"cse/models"
	"fmt" // Import fmt for debugging
	"regexp"
	"strings"
)

type ClassProcessor struct{}

func (p *ClassProcessor) Process(content string) interface{} {
	classRegex := regexp.MustCompile(`(?m)^(public|private|protected)?\s*(class|interface|enum)\s+(\w+)(?:\s+extends\s+(\w+))?(?:\s+implements\s+([\w,\s]+))?`)

	var classes []models.Class
	lines := strings.Split(content, "\n")

	for lineNum, line := range lines {
		matches := classRegex.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			// Get the full match details
			fullMatch := classRegex.FindStringSubmatch(line)

			// Debugging output
			fmt.Printf("Line: %d, Match: %s, Column: %d\n", lineNum+1, fullMatch[0], match[0]+1)

			class := models.Class{
				Name:    fullMatch[3],
				Extends: fullMatch[4],
				Location: models.Location{
					Line:   lineNum + 1,
					Column: match[0] + 1,
				},
			}

			if fullMatch[5] != "" {
				implements := regexp.MustCompile(`\s*,\s*`).Split(fullMatch[5], -1)
				class.Implements = strings.Join(implements, ",")
			}

			classes = append(classes, class)
		}
	}
	return classes
}
