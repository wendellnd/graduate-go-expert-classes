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

type Error struct {
	Message string `json:"message"`
}

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

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
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

	accessToken := dto.GetJWTOutput{
		AccessToken: token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
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
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
