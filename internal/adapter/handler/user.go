package handler

import (
	"encoding/json"
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"
	"event-planning-app/internal/response"
	"event-planning-app/internal/util"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
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

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.UserLogin

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "failed to read request body", err)
		return
	}

	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err)
		return
	}

	errValidate := util.Validate(uh.validate, req)
	if errValidate != nil {
		uh.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := uh.service.Login(req)
	if err != nil {
		uh.response.Error(w, http.StatusUnauthorized, "cannot login user", err.Error())
		return
	}

	_, err = uh.jwt.GenerateToken(w, *result)
	if err != nil {
		uh.response.Error(w, http.StatusInternalServerError, "failed to generate token", err.Error())
		return
	}

	uh.response.Success(w, http.StatusOK, "login successful", response.UserResponse{
		Name:     result.Name,
		Email:    result.Email,
		CreateAt: result.CreatedAt,
		UpdateAt: result.UpdatedAt,
	})
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.User

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "failed to read request body", err)
		return
	}

	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err)
		return
	}

	errValidate := util.Validate(uh.validate, req)
	if errValidate != nil {
		uh.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := uh.service.Create(req)
	if err != nil {
		uh.response.Error(w, http.StatusBadRequest, "cannot register user", err.Error())
		return
	}

	uh.response.Success(w, http.StatusOK, "successful register user", response.UserResponse{
		Name:     result.Name,
		Email:    result.Email,
		CreateAt: result.CreatedAt,
		UpdateAt: result.UpdatedAt,
	})
}

func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"message": "success create user",
		// "data":
	}

	uh.response.Success(w, http.StatusOK, "Successful register user", res)
}

func (uh *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"message": "success create user",
		// "data":
	}

	response, _ := json.Marshal(res)
	w.Write(response)
}

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"message": "success create user",
		// "data":
	}

	response, _ := json.Marshal(res)
	w.Write(response)
}

func (uh *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"message": "success create user",
		// "data":
	}

	response, _ := json.Marshal(res)
	w.Write(response)
}
