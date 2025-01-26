package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UUID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string
	Files []File
}

type File struct {
	gorm.Model
	ProjectID uint
	Name      string
	Content   string
	Classes   []Class
	Methods   []Method
	Variables []Variable
	Loops     []Loop
}

type Class struct {
	gorm.Model
	FileID        uint
	Name          string
	ParentClassID *uint
	ParentClass   *Class
	Methods       []Method
	Variables     []Variable
}

type Method struct {
	gorm.Model
	ClassID    uint
	FileID     uint
	Name       string
	ReturnType string
	Content    string
}

type Variable struct {
	gorm.Model
	ClassID  uint
	FileID   uint
	Name     string
	DataType string
}

type Loop struct {
	gorm.Model
	FileID    uint
	Type      string
	StartLine int
	EndLine   int
}

type Composition struct {
	gorm.Model
	ClassID         uint
	ComposedClassID uint
}

type Token struct {
	gorm.Model
	Value string `gorm:"uniqueIndex"`
}

type FileToken struct {
	gorm.Model
	TokenID   uint
	FileID    uint
	Frequency int
	Positions string
}

type ClassToken struct {
	gorm.Model
	TokenID   uint
	ClassID   uint
	Frequency int
	Positions string
}

type MethodToken struct {
	gorm.Model
	TokenID   uint
	MethodID  uint
	Frequency int
	Positions string
}

type VariableToken struct {
	gorm.Model
	TokenID    uint
	VariableID uint
	Frequency  int
	Positions  string
}
