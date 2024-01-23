package handler

import (
	"encoding/json"
	"errors"
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"
	"event-planning-app/internal/response"
	"event-planning-app/internal/util"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	service  port.UserService
	response util.Response
	jwt      util.JWTManager
	validate *validator.Validate
}

func NewUserHandler(service port.UserService) port.UserHandler {
	return &UserHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var entity domain.UserAuth

	c, err := r.Cookie("token")
	if err == nil && c != nil {
		u.response.Error(w, http.StatusUnauthorized, "you are already logged in", errors.New("user is already logged in").Error())
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		u.response.Error(w, http.StatusBadRequest, "failed to read request body", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &entity)
	if err != nil {
		u.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err.Error())
		return
	}

	errValidate := util.Validate(u.validate, entity)
	if errValidate != nil {
		u.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := u.service.Login(entity)
	if err != nil {
		u.response.Error(w, http.StatusUnauthorized, "cannot login user", err.Error())
		return
	}

	_, err = u.jwt.GenerateToken(w, *result)
	if err != nil {
		u.response.Error(w, http.StatusInternalServerError, "failed to generate token", err.Error())
		return
	}

	u.response.Success(w, http.StatusOK, "login successful", response.User{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var entity domain.User

	c, err := r.Cookie("token")
	if err == nil && c != nil {
		uh.response.Error(w, http.StatusUnauthorized, "you are already logged in", errors.New("user is already logged in").Error())
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "failed to read request body", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &entity)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err.Error())
		return
	}

	errValidate := util.Validate(uh.validate, entity)
	if errValidate != nil {
		uh.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := uh.service.Create(entity)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "cannot register user", err.Error())
		return
	}

	uh.response.Success(w, http.StatusOK, "successful register user", response.User{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := uh.service.GetAll()
	if err != nil {
		uh.response.Error(w, http.StatusInternalServerError, "failed to retrieve user data", err.Error())
		return
	}

	var userList []response.User
	for _, user := range result {
		userList = append(userList, response.User{
			ID:        user.ID,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	uh.response.Success(w, http.StatusOK, "successfully retrieved user data", userList)
}

func (uh *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "invalid user id", err.Error())
		return
	}

	result, err := uh.service.GetByID(uint(id))
	if err != nil {
		uh.response.Error(w, http.StatusInternalServerError, "failed to get user by id", err.Error())
		return
	}

	uh.response.Success(w, http.StatusOK, "successfully retrieved user by id", response.User{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var entity domain.User

	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "invalid user id", err.Error())
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "failed to read request body", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &entity)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err.Error())
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)
	entity.ID = claims.ID

	errValidate := util.Validate(uh.validate, entity)
	if errValidate != nil {
		uh.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := uh.service.Update(entity, uint(id))
	if err != nil {
		uh.response.Error(w, http.StatusInternalServerError, "failed to update user", err.Error())
		return
	}

	uh.response.Success(w, http.StatusOK, "successfully update user data", response.User{
		ID:        uint(id),
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (uh *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var entity domain.User
	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "invalid user id", err.Error())
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)
	entity.ID = uint(id)
	entity.Name = claims.Name

	err = uh.service.Delete(entity)
	if err != nil {
		uh.response.Error(w, http.StatusInternalServerError, "failed to delete user", err.Error())
		return
	}

	uh.response.Success(w, http.StatusOK, "successfully deleted user", nil)
}
