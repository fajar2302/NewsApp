package datausers

import (
	users "NEWSAPP/features/Users"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.DataUserInterface {
	return &userQuery{
		db: db,
	}
}

// CreateAccount implements users.DataUserInterface.
func (u *userQuery) CreateAccount(account users.User) error {
	userGorm := Users{
		ProfilePicture: account.ProfilePicture,
		FullName:       account.FullName,
		Email:          account.Email,
		Password:       account.Password,
		PhoneNumber:    account.PhoneNumber,
		Address:        account.Address,
	}
	tx := u.db.Create(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// AccountByEmail implements users.DataUserInterface.
func (u *userQuery) AccountByEmail(email string) (*users.User, error) {
	var userData Users
	tx := u.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var users = users.User{
		UserID:      userData.ID,
		FullName:    userData.FullName,
		Email:       userData.Email,
		Password:    userData.Password,
		Address:     userData.Address,
		PhoneNumber: userData.PhoneNumber,
	}

	return &users, nil
}

func (u *userQuery) AccountById(userid uint) (*users.User, error) {
	var userData Users
	tx := u.db.First(&userData, userid)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var user = users.User{
		ProfilePicture: userData.ProfilePicture,
		FullName:       userData.FullName,
		Email:          userData.Email,
		PhoneNumber:    userData.PhoneNumber,
		Address:        userData.Address,
	}

	return &user, nil
}

func (u *userQuery) UpdateAccount(userid uint, account users.User) error {
	var userGorm Users
	tx := u.db.First(&userGorm, userid)
	if tx.Error != nil {
		return tx.Error
	}
	userGorm.ProfilePicture = account.ProfilePicture
	userGorm.FullName = account.FullName
	userGorm.Email = account.Email
	userGorm.Password = account.Password
	userGorm.PhoneNumber = account.PhoneNumber
	userGorm.Address = account.Address

	tx = u.db.Save(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *userQuery) DeleteAccount(userid uint) error {
	tx := u.db.Delete(&Users{}, userid)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
