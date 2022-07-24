package model

import "errors"

type User struct {
	ID       int
	Nama     string
	Email    string
	Password string
}

type UserModel struct {
	data []User
}

func GenerateID(last int) int {
	if last > 0 {
		last++
		return last
	}
	return 0
}

func (um *UserModel) Insert(newUser User) (User, error) {
	newUser.ID = GenerateID(len(um.data))
	um.data = append(um.data, newUser)

	return newUser, nil
}

// func (um *UserModel) Update(updatedUser User) (User, error) {
// 	if err := um.db.Save(&updatedUser).Error; err != nil {
// 		return User{}, err
// 	}

// 	return updatedUser, nil
// }
func (um *UserModel) GetAll() ([]User, error) {
	if len(um.data) == 0 {
		return nil, errors.New("no record on database")
	}
	return um.data, nil
}

// func (um *UserModel) Login(email string, password string) (User, error) {
// 	var res User
// 	if err := um.db.Where("email = ?", email).First(&res).Error; err != nil {
// 		return User{}, err
// 	}

// 	err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
// 	if err != nil {
// 		return User{}, errors.New("password salah")
// 	}

// 	return res, nil

// }
