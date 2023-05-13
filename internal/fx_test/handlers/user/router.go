package user

type Router struct {
	service service
	logger  logger
}

func NewRouter(service service, logger logger) *Router {
	return &Router{
		service: service,
		logger:  logger,
	}
}

func (r *Router) GetUser() {
	r.logger.Log("handler get user")
	r.service.GetUser()
}

func (r *Router) CreateUser() {
	r.logger.Log("handler create user")
	r.service.CreateUser()
}

func (r *Router) UpdateUser() {
	r.logger.Log("handler update user")
	r.service.UpdateUser()
}

func (r *Router) DeleteUser() {
	r.logger.Log("handler delete user")
	r.service.DeleteUser()
}
