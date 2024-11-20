package projects

type ProjectService struct {
	projectRepository ProjectRepository
}

func NewService(r ProjectRepository) ProjectService {
	return ProjectService{projectRepository: r}
}
func (svc *ProjectService) AddUserProject(project AddProjectsInBulk) (int, error) {
	return svc.projectRepository.AddUserProject(project)
}

func (svc *ProjectService) GetProjectById(id uint) (*Project, error) {
	return svc.projectRepository.GetProjectById(id)
}

func (svc *ProjectService) GetUserProjectByUserAndProjectId(userId, id uint) (*UserProject, error) {
	return svc.projectRepository.GetUserProjectByUserAndProjectId(userId, id)
}

func (svc *ProjectService) UpdateProject(project *Project) error {
	return svc.projectRepository.UpdateProject(project)
}

func (svc *ProjectService) GetUserProjectByUserId(userId, projId uint) (*Project, error) {
	return svc.projectRepository.GetUserProjectByUserId(userId, projId)
}

func (svc *ProjectService) DeleteUserProjectByID(userId uint) error {
	return svc.projectRepository.DeleteUserProjectByID(userId)
}

func (svc *ProjectService) GetAllUserProject(userId uint, limit int, offset int, orderBy string) ([]Project, uint, error) {
	return svc.projectRepository.GetAllUserProject(userId, limit, offset, orderBy)
}
func (svc *ProjectService) AddBulkUserProject(userId uint, addUserProjectRequest AddUserProjectRequest) (int, error) {
	return svc.projectRepository.AddBulkUserProject(userId, addUserProjectRequest)
}
