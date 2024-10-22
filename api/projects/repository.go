package projects

// ProjectRepository Used to store and retrieve projects
type ProjectRepository interface {
	AddUserProject(userId uint, project *Project) (*Project, error)
	GetProjectById(id uint) (*Project, error)
	GetUserProjectByUserAndProjectId(userId, id uint) (*UserProject, error)
	UpdateProject(project *Project) error
	GetUserProjectByUserId(userId, projId uint) (*Project, error)
	DeleteUserProjectByID(userId uint) error
	GetAllUserProject(userId uint, limit int, offset int, orderBy string) ([]Project, uint, error)
	//createCategories(jsonData []Category) error
}
