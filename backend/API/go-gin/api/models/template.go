package models

import "gorm.io/gorm"

// This is a template model, and it is migrated in "migrations.go" file.

type Template struct {
	gorm.Model
	Text string
}

