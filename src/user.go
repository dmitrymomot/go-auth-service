package main

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User model
type User struct {
	// gorm.Model
	ID        uuid.UUID `gorm:"primary_key;type:char(36);column:id"`
	Email     string    `gorm:"type:varchar(100);unique_index;column:email;not null;"`
	Password  string    `gorm:"type:varchar(100);column:password;not null;"`
	APIToken  uuid.UUID `gorm:"type:char(36);column:api_token"`
	Name      string    `gorm:"type:varchar(100);column:name;not null;"`
	RoleID    uuid.UUID `gorm:"type:char(36);column:role_id;index"`
	Role      Role      `gorm:"ForeignKey:RoleId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Role model
type Role struct {
	// gorm.Model
	ID    uuid.UUID `gorm:"primary_key;type:char(36);column:id"`
	Name  string    `gorm:"type:varchar(20);column:name;not null;index"`
	Level int       `gorm:"type:int;column:level;not null"`
}
