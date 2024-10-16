package user

// UserRepository Used to store and retrieve user details
type UserRepository interface {
	createCategories(jsonData []UserCategory) error
	createRoles(jsonData []Role) error
	CreateUser(user *User) (*User, error)
	GetAllUser(keyword string, limit int, offset int, orderBy string) ([]User, uint, error)
	GetUserDetailsByUserId(userId uint) (*User, error)
	DeleteUserByUserID(userId uint) error
	UpdateUserByUserID(userId uint, user *User) (*User, error)
}