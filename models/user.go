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

func (this *User) Save() *User {
	if this.ID == 0 {
		db.DB.Create(this)
	} else {
		db.DB.Save(this)
	}

	return this
}

func (this *User) FindByEmail(email string) {
	db.DB.Where("email = ?", email).First(this)
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
