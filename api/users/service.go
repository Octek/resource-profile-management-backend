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
