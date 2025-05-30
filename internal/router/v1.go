package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/usmaarn/yummeals_api/internal/handler"
	"github.com/usmaarn/yummeals_api/internal/service"
	"github.com/usmaarn/yummeals_api/internal/storage"
	"github.com/usmaarn/yummeals_api/pkg/utils"
)

type V1Router struct {
	db       *sqlx.DB
	validate *validator.Validate
}

func NewV1Router(db *sqlx.DB, validate *validator.Validate) *V1Router {
	return &V1Router{db, validate}
}

func (v1 *V1Router) Register() chi.Router {
	r := chi.NewRouter()
	storage := storage.NewStorage(v1.db)
	service := service.RegisterServices(storage)

	api := utils.Api{
		Validate: v1.validate,
		Service:  service,
		Storage:  storage,
	}

	//Auth routes
	auth := handler.NewAuthHandler(&api)
	r.Post("/auth/register", auth.RegisterHandler)
	r.Post("/auth/login", auth.LoginHandler)
	r.Post("/auth/logout", auth.LogoutHandler)
	r.Get("/auth/user", auth.GetAuthenticatedUserHandler)

	return r
}
