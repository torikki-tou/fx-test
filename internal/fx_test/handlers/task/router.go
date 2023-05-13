package task

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

func (r *Router) GetTask() {
	r.logger.Log("handler get task")
	r.service.GetTask()
}

func (r *Router) CreateTask() {
	r.logger.Log("handler create task")
	r.service.CreateTask()
}

func (r *Router) UpdateTask() {
	r.logger.Log("handler update task")
	r.service.UpdateTask()
}

func (r *Router) DeleteTask() {
	r.logger.Log("handler delete task")
	r.service.DeleteTask()
}
