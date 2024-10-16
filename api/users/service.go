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
