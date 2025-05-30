package handler

import (
	"net/http"

	"github.com/usmaarn/yummeals_api/internal/dto"
	"github.com/usmaarn/yummeals_api/pkg/utils"
)

type AuthHandler struct {
	api *utils.Api
}

func NewAuthHandler(api *utils.Api) *AuthHandler {
	return &AuthHandler{api: api}
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var dto dto.LoginUserRequest
	if err := h.api.BindRequestBody(r, &dto); err != nil {
		h.api.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	token, err := h.api.Service.AuthService.Login(dto)
	if err != nil {
		h.api.ErrorResponse(w, http.StatusBadRequest, "Login failed", map[string]string{"message": err.Error()})
		return
	}
	h.api.SuccessResponse(w, http.StatusOK, token)
}

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic for user logout
	// This might involve invalidating a session or token, etc.
}

func (h *AuthHandler) GetAuthenticatedUserHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to retrieve the authenticated user's information
	// This might involve checking the session or token and returning user details
}
