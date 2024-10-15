package projects

type ProjectService struct {
	projectRepository ProjectRepository
}

func NewService(r ProjectRepository) ProjectService {
	return ProjectService{projectRepository: r}
}
