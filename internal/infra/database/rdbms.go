package database

import "gorm.io/gorm"

type RDBMS struct {
	DB *gorm.DB
}
