package models

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	token "golang_final_project/cmd/gym-here/utils"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Role     string `gorm:"size:255" json:"role"`
}

func GetUserByID(uid uint) (User, error) {
	var u User
	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}
	u.PrepareGive()
	return u, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {
	var err error
	u := User{}

	// Debug: log the username attempting to login
	log.Println("Login attempt for username:", username)

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		// Debug: log the error
		log.Println("Error retrieving user:", err)
		return "", err
	}

	// Debug: log the retrieved user details
	log.Printf("Retrieved user: %+v\n", u)

	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		// Debug: log the password mismatch
		log.Println("Password mismatch:", err)
		return "", err
	}

	token, err := token.GenerateToken(u.ID)
	if err != nil {
		// Debug: log the token generation error
		log.Println("Error generating token:", err)
		return "", err
	}

	return token, nil
}

func (u *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}
