package main

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	// gorm.Model
	ID        uuid.UUID `gorm:"primary_key;type:char(36);column:id"`
	Email     string    `gorm:"type:varchar(100);unique_index;column:email;not null;"`
	Password  string    `gorm:"type:varchar(100);column:password;not null;"`
	APIKey    uuid.UUID `gorm:"type:char(36);column:api_key"`
	Name      string    `gorm:"type:varchar(100);column:name;not null;"`
	RoleID    uuid.UUID `gorm:"type:char(36);column:role_id;index"`
	Role      Role      `gorm:"ForeignKey:RoleId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// GetAPIKey get user api key
func (user *User) GetAPIKey() string {
	if user.APIKey == uuid.Nil {
		user.APIKey = uuid.NewV1()
		db.Save(&user)
	}
	return user.APIKey.String()
}

// GenAccessToken generates JWT
func (user *User) GenAccessToken() (ss string, err error) {
	signingKey := []byte(user.GetAPIKey())
	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Duration(24*30) * time.Hour), // 1 month
			Issuer:    "user",
		},
		*user,
		user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err = token.SignedString(signingKey)
	return
}

// SetPassword hashes string
func (user *User) SetPassword(password string) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(bytes)
	return
}

// CheckPassword compare hash with password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
