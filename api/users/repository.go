package user

// UserRepository Used to store and retrieve user details
type UserRepository interface {
	createCategories(jsonData []UserCategory) error
	createRoles(jsonData []Role) error
	CreateUser(user *User) (*User, error)
	GetAllUser(keyword string, limit int, offset int, orderBy string) ([]User, uint, error)
	GetUserDetailsByUserId(userId uint) (*User, error)
	DeleteUserByUserID(userId uint) error
	UpdateUserByUserID(user *User) (*User, error)
	AddUserEducation(education Education) (Education, error)
	GetEducationById(id uint) (*Education, error)
	GetUserEducationByUserAndEducationId(userId, id uint) (*Education, error)
	UpdateEducation(education *Education) error
	GetUserEducationByUserId(userId uint) (*Education, error)
	DeleteUserEducationByID(userId uint) error
	GetAllUserEducation(userId uint, limit int, offset int, orderBy string) ([]Education, uint, error)
}
