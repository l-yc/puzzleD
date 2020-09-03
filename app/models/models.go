package models

import (
	"time"
  "github.com/jinzhu/gorm"
)

// Define GORM-backend models
type User struct {
  gorm.Model
	Name				string
	AuthoredPuzzles []Puzzle `gorm:"many2many:puzzle_authors"`
	SolvedPuzzles []Puzzle	`gorm:"many2many:puzzle_solves"`
}

type Puzzle struct {
  gorm.Model
  Name        string
  Description string
	//Attachment	oss.OSS `sql:"size:4294967295;"` //`sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{filename_with_hash}}"`
	//Attachment  media_library.MediaLibraryStorage `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}}"`
	ReleaseDate	time.Time
	Authors			[]User	`gorm:"many2many:puzzle_authors"`
	Solution		string
	SolvedUsers []User	`gorm:"many2many:puzzle_solves"`
}
