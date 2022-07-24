package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel struct {
	Data []User
}

func GenerateID(last int) int {
	if last > 0 {
		last++
		return last
	}
	return 1
}

func (um *UserModel) Insert(newUser User) (User, error) {
	newUser.ID = GenerateID(len(um.Data))
	encryptPass, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(encryptPass)
	um.Data = append(um.Data, newUser)

	return newUser, nil
}

// func (um *UserModel) Update(updatedUser User) (User, error) {
// 	if err := um.db.Save(&updatedUser).Error; err != nil {
// 		return User{}, err
// 	}

// 	return updatedUser, nil
// }
func (um *UserModel) GetAll() ([]User, error) {
	if len(um.Data) == 0 {
		return nil, errors.New("no record on database")
	}
	return um.Data, nil
}

func (um *UserModel) Login(email string, password string) (User, error) {
	var res User
	for _, val := range um.Data {
		if val.Email == email {
			res = val
		}
	}

	if res.ID == 0 {
		return User{}, errors.New("no data found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
	if err != nil {
		return User{}, errors.New("wrong password")
	}

	return res, nil

}
