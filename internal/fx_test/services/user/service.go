package user

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

func (s *Service) GetUser() {
	s.logger.Log("service get user")
	s.repository.GetUser()
}

func (s *Service) CreateUser() {
	s.logger.Log("service create user")
	s.repository.CreateUser()

}

func (s *Service) UpdateUser() {
	s.logger.Log("service update user")
	s.repository.UpdateUser()

}

func (s *Service) DeleteUser() {
	s.logger.Log("service delete user")
	s.repository.DeleteUser()
}
