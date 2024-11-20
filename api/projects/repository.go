package projects

// ProjectRepository Used to store and retrieve projects
type ProjectRepository interface {
	AddUserProject(project AddProjectsInBulk) (int, error)
	GetProjectById(id uint) (*Project, error)
	GetUserProjectByUserAndProjectId(userId, id uint) (*UserProject, error)
	UpdateProject(project *Project) error
	GetUserProjectByUserId(userId, projId uint) (*Project, error)
	DeleteUserProjectByID(userId uint) error
	GetAllUserProject(userId uint, limit int, offset int, orderBy string) ([]Project, uint, error)
	AddBulkUserProject(userId uint, addUserProjectRequest AddUserProjectRequest) (int, error)
	//createCategories(jsonData []Category) error
}
