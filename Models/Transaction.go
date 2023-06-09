package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Hash      string `gorm:"uniqueIndex"`
	BlockHash string
	Type      int
	From      string
	To        string
	Value     string
	Input     string
	BlockID   uint
}
