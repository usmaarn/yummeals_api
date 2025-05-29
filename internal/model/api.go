package model

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type API struct {
	Router   chi.Router
	DB       *sqlx.DB
	Validate *validator.Validate
}
