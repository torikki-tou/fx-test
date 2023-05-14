package task

type Repository struct {
	logger logger
	db     db
}

func NewRepository(db db, logger logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) GetTask() {
	r.logger.Log("repo get task")
	r.db.Select()
}

func (r *Repository) CreateTask() {
	r.logger.Log("repo create task")
	r.db.Insert()
}

func (r *Repository) UpdateTask() {
	r.logger.Log("repo update task")
	r.db.Update()
}

func (r *Repository) DeleteTask() {
	r.logger.Log("repo delete task")
	r.db.Delete()
}
