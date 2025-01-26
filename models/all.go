package models

import (
	"gorm.io/gorm"
)

type Location struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type Project struct {
	gorm.Model
	Name  string
	Files []File `gorm:"foreignKey:ProjectID"`
}

type File struct {
	gorm.Model
	Name          string
	ProjectID     uint
	Project       Project        `gorm:"foreignKey:ProjectID"`
	Classes       []Class        `gorm:"foreignKey:FileID"`
	Methods       []Method       `gorm:"foreignKey:FileID"`
	Variables     []Variable     `gorm:"foreignKey:FileID"`
	Relationships []Relationship `gorm:"foreignKey:FileID"`
}

type Class struct {
	gorm.Model
	Name       string
	FileID     uint
	File       File `gorm:"foreignKey:FileID"`
	Extends    string
	Implements string   `gorm:"type:text"`
	Location   Location `gorm:"embedded;embeddedPrefix:location_"`
}

type Method struct {
	gorm.Model
	Name       string
	FileID     uint
	File       File `gorm:"foreignKey:FileID"`
	ClassID    uint
	Class      Class `gorm:"foreignKey:ClassID"`
	ReturnType string
	Parameters string   `gorm:"type:text"`
	Location   Location `gorm:"embedded;embeddedPrefix:location_"`
}

type Variable struct {
	gorm.Model
	Name     string
	Type     string
	FileID   uint
	File     File `gorm:"foreignKey:FileID"`
	ClassID  uint
	MethodID uint
	Location Location `gorm:"embedded;embeddedPrefix:location_"`
}

type Relationship struct {
	gorm.Model
	FileID           uint
	File             File `gorm:"foreignKey:FileID"`
	SourceClassID    uint
	TargetClassID    uint
	RelationshipType string
	SourceClassName  string
	TargetClassName  string
	Location         Location `gorm:"embedded;embeddedPrefix:location_"`
}
