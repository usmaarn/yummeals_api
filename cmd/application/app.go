package application

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/usmaarn/yummeals_api/internal/router"
	"github.com/usmaarn/yummeals_api/pkg/config"
)

type Application struct {
	Router chi.Router
}

func New(r chi.Router) *Application {
	return &Application{
		Router: r,
	}
}

func (a *Application) RegisterMiddleware() {
	a.Router.Use(middleware.Logger)
	a.Router.Use(middleware.Recoverer)
	a.Router.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func (a *Application) RegisterRoutes() {
	//Initialize database and validator instances here
	config.LoadEnv()
	db := config.InitlializeDatabase()
	validate := config.InitializeValidator()

	v1 := router.NewV1Router(db, validate)
	a.Router.Mount("/v1", v1.Register())
}

func (a *Application) Run(addr string) error {
	fmt.Printf("Starting server on %s\n", addr)
	return http.ListenAndServe(addr, a.Router)
}
