package user

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

func (r *Repository) GetUser() {
	r.logger.Log("repo get user")
	r.db.Select()
}

func (r *Repository) CreateUser() {
	r.logger.Log("repo create user")
	r.db.Insert()
}

func (r *Repository) UpdateUser() {
	r.logger.Log("repo update user")
	r.db.Update()
}

func (r *Repository) DeleteUser() {
	r.logger.Log("repo delete user")
	r.db.Delete()
}
