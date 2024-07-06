package service

import (
	"NEWSAPP/app/middlewares"
	users "NEWSAPP/features/Users"
	"NEWSAPP/utils/encrypts"
	"errors"
)

type userService struct {
	userData    users.DataUserInterface
	hashService encrypts.HashInterface
}

func New(ud users.DataUserInterface, hash encrypts.HashInterface) users.ServiceUserInterface {
	return &userService{
		userData:    ud,
		hashService: hash,
	}

}

// LoginAccount implements users.ServiceUserInterface.
func (u *userService) LoginAccount(email string, password string) (data *users.User, token string, err error) {
	data, err = u.userData.AccountByEmail(email)
	if err != nil {
		return nil, "", err
	}

	isLoginValid := u.hashService.CheckPasswordHash(data.Password, password)
	// ketika isloginvalid = true, maka login berhasil
	if !isLoginValid {
		return nil, "", errors.New("[validation] email atau password tidak sesuai")
	}

	token, errJWT := middlewares.CreateToken(int(data.UserID))
	if errJWT != nil {
		return nil, "", errJWT
	}
	return data, token, nil

}

// RegistrasiAccount implements users.ServiceUserInterface.
func (u *userService) RegistrasiAccount(accounts users.User) error {
	if accounts.FullName == "" || accounts.Email == "" || accounts.Password == "" || accounts.PhoneNumber == "" || accounts.Address == "" {
		return errors.New("[validation] nama/email/password/phone/address tidak boleh kosong")
	}

	// proses hash password
	var errHash error
	if accounts.Password, errHash = u.hashService.HashPassword(accounts.Password); errHash != nil {
		return errHash
	}

	return u.userData.CreateAccount(accounts)
}

// UpdateProfile implements users.ServiceUserInterface.
func (u *userService) UpdateProfile(userid uint, accounts users.User) error {
	if accounts.FullName == "" || accounts.Email == "" || accounts.Password == "" || accounts.PhoneNumber == "" || accounts.Address == "" {
		return errors.New("[validation] nama/email/password/phone/address tidak boleh kosong")
	}

	// proses hash password
	var errHash error
	if accounts.Password, errHash = u.hashService.HashPassword(accounts.Password); errHash != nil {
		return errHash
	}

	return u.userData.UpdateAccount(userid, accounts)
}

func (us *userService) DeleteAccount(userid uint) error {
	err := us.userData.DeleteAccount(userid)
	if err != nil {
		return err
	}
	return nil
}

func (us *userService) GetProfile(userid uint) (*users.User, error) {
	profile, err := us.userData.AccountById(userid)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
