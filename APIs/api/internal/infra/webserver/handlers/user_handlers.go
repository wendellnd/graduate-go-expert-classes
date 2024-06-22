package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/wendellnd/graduate-go-expert-classes/api/internal/dto"
	"github.com/wendellnd/graduate-go-expert-classes/api/internal/entity"
	"github.com/wendellnd/graduate-go-expert-classes/api/internal/infra/database"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       userDB,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var jwt dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&jwt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(jwt.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(jwt.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, err := h.Jwt.Encode(map[string]interface{}{
		"sub":   u.ID.String(),
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * time.Duration(h.JwtExpiresIn)).Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
