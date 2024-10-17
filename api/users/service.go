package user

type UserService struct {
	userRepository UserRepository
}

func NewService(r UserRepository) UserService {
	return UserService{userRepository: r}
}

func (svc *UserService) CreateCategories(jsonData []UserCategory) error {
	return svc.userRepository.createCategories(jsonData)
}
func (svc *UserService) CreateRoles(jsonData []Role) error {
	return svc.userRepository.createRoles(jsonData)
}
func (svc *UserService) CreateUser(user *User) (*User, error) {
	return svc.userRepository.CreateUser(user)
}
func (svc *UserService) GetAllUser(keyword string, limit int, offset int, orderBy string) ([]User, uint, error) {
	return svc.userRepository.GetAllUser(keyword, limit, offset, orderBy)
}
func (svc *UserService) GetUserDetailsByUserId(userId uint) (*User, error) {
	return svc.userRepository.GetUserDetailsByUserId(userId)
}
func (svc *UserService) DeleteUserByUserID(userId uint) error {
	return svc.userRepository.DeleteUserByUserID(userId)
}
func (svc *UserService) UpdateUserByUserID(user *User) (*User, error) {
	return svc.userRepository.UpdateUserByUserID(user)
}
func (svc *UserService) GetAllUserCategories(keyword string, limit int, offset int, orderBy string) ([]UserCategory, int64, error) {
	return svc.userRepository.GetAllUserCategories(keyword, limit, offset, orderBy)
}
