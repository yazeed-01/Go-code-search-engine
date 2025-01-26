package processors

import (
	"cse/models"
	"regexp"
)

type RelationshipProcessor struct{}

func (p *RelationshipProcessor) Process(content string) interface{} {
	inheritanceRegex := regexp.MustCompile(`(?m)^(public|private|protected)?\s*class\s+(\w+)\s+extends\s+(\w+)`)
	compositionRegex := regexp.MustCompile(`(?m)^(public|private|protected)?\s*(\w+(?:<.*?>)?)\s+(\w+)\s*;`)

	var relationships []models.Relationship

	// Process inheritance relationships
	inheritanceMatches := inheritanceRegex.FindAllStringSubmatch(content, -1)
	for _, match := range inheritanceMatches {
		relationships = append(relationships, models.Relationship{
			SourceClassName:  match[2],
			TargetClassName:  match[3],
			RelationshipType: "Inheritance",
		})
	}

	// Process composition relationships
	compositionMatches := compositionRegex.FindAllStringSubmatch(content, -1)
	for _, match := range compositionMatches {
		if match[2] != "String" && match[2] != "int" && match[2] != "boolean" && match[2] != "double" && match[2] != "float" {
			relationships = append(relationships, models.Relationship{
				SourceClassName:  "", // Will be set later when we know the containing class
				TargetClassName:  match[2],
				RelationshipType: "Composition",
			})
		}
	}

	return relationships
}
