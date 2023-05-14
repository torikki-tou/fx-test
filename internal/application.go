package internal

import (
	"context"
	taskHandler "github.com/torikki-tou/fx-test/internal/api/handlers/task"
	userHandler "github.com/torikki-tou/fx-test/internal/api/handlers/user"
	"github.com/torikki-tou/fx-test/internal/components/config"
	"github.com/torikki-tou/fx-test/internal/components/db"
	"github.com/torikki-tou/fx-test/internal/components/logger"
	"github.com/torikki-tou/fx-test/internal/components/server"
	taskRepository "github.com/torikki-tou/fx-test/internal/repositories/task"
	userRepository "github.com/torikki-tou/fx-test/internal/repositories/user"
	taskService "github.com/torikki-tou/fx-test/internal/services/task"
	userService "github.com/torikki-tou/fx-test/internal/services/user"
	"go.uber.org/fx"
)

func Run() {
	fx.New(
		fx.Provide(
			func() *config.Config {
				return config.New()
			},
			func(config *config.Config) *logger.Logger {
				return logger.New(config)
			},
			func(config *config.Config, logger *logger.Logger) *db.Connection {
				return db.New(config, logger)
			},
			func(config *config.Config, logger *logger.Logger) *server.Server {
				return server.New(config, logger)
			},
			func(db *db.Connection, logger *logger.Logger) *taskRepository.Repository {
				return taskRepository.NewRepository(db, logger)
			},
			func(db *db.Connection, logger *logger.Logger) *userRepository.Repository {
				return userRepository.NewRepository(db, logger)
			},
			func(repository *taskRepository.Repository, logger *logger.Logger) *taskService.Service {
				return taskService.NewService(repository, logger)
			},
			func(repository *userRepository.Repository, logger *logger.Logger) *userService.Service {
				return userService.NewService(repository, logger)
			},
			func(service *taskService.Service, logger *logger.Logger) *taskHandler.Router {
				return taskHandler.NewRouter(service, logger)
			},
			func(service *userService.Service, logger *logger.Logger) *userHandler.Router {
				return userHandler.NewRouter(service, logger)
			},
		),
		fx.Invoke(
			runServer,
		),
	).Run()
}

func runServer(server *server.Server, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.Serve()
			return nil
		},
	})
}
