package main

import uuid "github.com/satori/go.uuid"

// Role model
type Role struct {
	ID    uuid.UUID `gorm:"primary_key;type:char(36);column:id"`
	Name  string    `gorm:"type:varchar(20);column:name;not null;index"`
	Level int       `gorm:"type:int;column:level;not null"`
}

// GetByName get role object by name
func (role *Role) GetByName(name string) {
	db.Where("name = ?", name).First(&role)
}
