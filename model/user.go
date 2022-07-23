package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Nama     string
	Email    string
	Password string
}

type UserModel struct {
	db *gorm.DB
}

func (um *UserModel) Insert(newUser User) (User, error) {
	if err := um.db.Create(&newUser).Error; err != nil {
		return User{}, err
	}

	return newUser, nil
}
func (um *UserModel) Update(updatedUser User) (User, error) {
	if err := um.db.Save(&updatedUser).Error; err != nil {
		return User{}, err
	}

	return updatedUser, nil
}
func (um *UserModel) GetAll() ([]User, error) {
	var res []User
	if err := um.db.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func (um *UserModel) Login(email string, password string) (User, error) {
	var res User
	if err := um.db.Where("email = ?", email).First(&res).Error; err != nil {
		return User{}, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
	if err != nil {
		return User{}, errors.New("password salah")
	}

	return res, nil

}
