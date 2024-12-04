package questions

type QuestionService struct {
	questionRepository QuestionRepository
}

func NewService(r QuestionRepository) QuestionService {
	return QuestionService{questionRepository: r}
}

func (svc *QuestionService) AddQuestion(question AddQuestionRequestBulk) error {
	return svc.questionRepository.AddQuestion(question)
}

func (svc *QuestionService) GetQuestionById(id uint) (*Question, error) {
	return svc.questionRepository.GetQuestionById(id)
}
func (svc *QuestionService) UpdateQuestionByID(question *Question) error {
	return svc.questionRepository.UpdateQuestionByID(question)
}

func (svc *QuestionService) GetAllQuestions(limitInt, offsetInt int, orderBy string) ([]Question, int64, error) {
	return svc.questionRepository.GetAllQuestions(limitInt, offsetInt, orderBy)
}

func (svc *QuestionService) DeleteQuestionByID(questionId uint) error {
	return svc.questionRepository.DeleteQuestionByID(questionId)
}
