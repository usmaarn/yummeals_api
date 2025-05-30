package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/usmaarn/yummeals_api/internal/dto"
	"github.com/usmaarn/yummeals_api/internal/repository"
	"github.com/usmaarn/yummeals_api/internal/service"
)

type Api struct {
	Validate   *validator.Validate
	Service    *service.Service
	Repository *repository.Repository
}

func (a *Api) BindRequestBody(r *http.Request, entity any) error {
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		return err
	}
	defer r.Body.Close()

	if err := a.Validate.Struct(entity); err != nil {
		return err
	}

	return nil
}

func (a *Api) ErrorResponse(w http.ResponseWriter, statusCode int, message string, errors any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := dto.ApiResponse{
		Success:    false,
		Message:    message,
		StatusCode: statusCode,
		Errors:     errors,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}

func (a *Api) SuccessResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := dto.ApiResponse{
		Success:    true,
		Data:       data,
		StatusCode: statusCode,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode success response", http.StatusInternalServerError)
	}
}
