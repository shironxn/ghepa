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

type CommentHandler struct {
	service  port.CommentService
	response util.Response
	validate *validator.Validate
}

func NewCommentHandler(service port.CommentService) port.CommentHandler {
	return &CommentHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (c *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var entity domain.Comment

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		c.response.Error(w, http.StatusBadRequest, "failed to read request body", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &entity)
	if err != nil {
		c.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err.Error())
		return
	}

	errValidate := util.Validate(c.validate, entity)
	if errValidate != nil {
		c.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := c.service.Create(entity)
	if err != nil {
		c.response.Error(w, http.StatusBadRequest, "failed to create comment", err.Error())
		return
	}

	c.response.Success(w, http.StatusOK, "successfully to create comment", response.Comment{
		UserName:  result.User.Name,
		EventName: result.Event.Title,
		Comment:   result.Comment,
		CreateAt:  result.CreatedAt,
		UpdateAt:  result.UpdatedAt,
	})
}

func (c *CommentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetAll()
	if err != nil {
		c.response.Error(w, http.StatusInternalServerError, "failed to get all comment", err.Error())
		return
	}

	c.response.Success(w, http.StatusOK, "successfully get all comments", result)
}
