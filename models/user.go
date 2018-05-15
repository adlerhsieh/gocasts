package models

import (
	"database/sql"

	"github.com/adlerhsieh/gocasts/db"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    sql.NullString
	Username sql.NullString
	Password sql.NullString
	Role     sql.NullString
}

func (this *User) IsAdmin() bool {
	return this.Role.String == "admin"
}

func AuthenticateUser(email string, password string) User {
	user := User{}
	db.DB.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return user
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password.String),
		[]byte(password),
	)
	if err != nil {
		return User{}
	}

	return user
}
