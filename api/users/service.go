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
func (svc *UserService) AddUserEducation(education Education) (Education, error) {
	return svc.userRepository.AddUserEducation(education)
}

func (svc *UserService) GetEducationById(id uint) (*Education, error) {
	return svc.userRepository.GetEducationById(id)
}

func (svc *UserService) GetUserEducationByUserAndEducationId(userId, id uint) (*Education, error) {
	return svc.userRepository.GetUserEducationByUserAndEducationId(userId, id)
}

func (svc *UserService) UpdateEducation(education *Education) error {
	return svc.userRepository.UpdateEducation(education)
}

func (svc *UserService) GetUserEducationByUserId(userId uint) (*Education, error) {
	return svc.userRepository.GetUserEducationByUserId(userId)
}

func (svc *UserService) DeleteUserEducationByID(userId uint) error {
	return svc.userRepository.DeleteUserEducationByID(userId)
}

func (svc *UserService) GetAllUserEducation(userId uint, limit int, offset int, orderBy string) ([]Education, uint, error) {
	return svc.userRepository.GetAllUserEducation(userId, limit, offset, orderBy)
}
