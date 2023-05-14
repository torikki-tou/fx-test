package task

type Service struct {
	repository repository
	logger     logger
}

func NewService(repository repository, logger logger) *Service {
	return &Service{
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) GetTask() {
	s.logger.Log("service get task")
	s.repository.GetTask()
}

func (s *Service) CreateTask() {
	s.logger.Log("service create task")
	s.repository.CreateTask()

}

func (s *Service) UpdateTask() {
	s.logger.Log("service update task")
	s.repository.UpdateTask()

}

func (s *Service) DeleteTask() {
	s.logger.Log("service delete task")
	s.repository.DeleteTask()
}
