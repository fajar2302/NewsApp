package users

type User struct {
	UserID      uint
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
}

type DataUserInterface interface {
	CreateAccount(account User) error
	AccountByEmail(email string) (*User, error)
	AccountById(userid uint) (*User, error)
	UpdateAccount(userid uint, account User) error
	DeleteAccount(userid uint) error
}

type ServiceUserInterface interface {
	RegistrasiAccount(accounts User) error
	LoginAccount(email string, password string) (data *User, token string, err error)
	GetProfile(userid uint) (data *User, err error)
	UpdateProfile(userid uint, accounts User) error
	DeleteAccount(userid uint) error
}
