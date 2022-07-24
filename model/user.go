package model

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var (
	listData []User
)

func init() {
	listData = []User{}
}

type User struct {
	ID       int
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel struct{}

func GenerateID(last int) int {
	if last > 0 {
		last++
		return last
	}
	return 1
}

func (um *UserModel) Insert(newUser User) (User, error) {
	if len(balanceHistory) == 0 {
		newUser.ID = GenerateID(0)
	} else {
		newUser.ID = GenerateID(listData[len(listData)-1].ID)
	}
	encryptPass, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(encryptPass)
	listData = append(listData, newUser)

	return newUser, nil
}

func (um *UserModel) Update(updatedUser User) (User, error) {
	use := -1
	for idx, val := range listData {
		if val.ID == updatedUser.ID {
			use = idx
		}
	}

	if use == -1 {
		return User{}, errors.New("no data found")
	}

	if updatedUser.Nama != "" {
		listData[use].Nama = updatedUser.Nama
	}

	if updatedUser.Email != "" {
		listData[use].Email = updatedUser.Email
	}

	return listData[use], nil
}
func (um *UserModel) GetAll() ([]User, error) {
	if len(listData) == 0 {
		return nil, errors.New("no record on database")
	}
	return listData, nil
}

func (um *UserModel) Login(email string, password string) (User, error) {
	log.Println(email, password, listData)
	var res User
	for _, val := range listData {
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
