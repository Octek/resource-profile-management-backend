package questions

type QuestionService struct {
	questionRepository QuestionRepository
}

func NewService(r QuestionRepository) QuestionService {
	return QuestionService{questionRepository: r}
}
