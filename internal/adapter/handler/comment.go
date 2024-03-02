package handler

import (
	"encoding/json"
	"ghepa/internal/core/domain"
	"ghepa/internal/core/port"
	"ghepa/internal/util"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
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
	var entity domain.CommentRequest

	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		c.response.Error(w, http.StatusBadRequest, "invalid user id", err.Error())
		return
	}

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

	claims := r.Context().Value("claims").(*domain.Claims)
	entity.UserID = claims.ID
	entity.EventID = uint(id)

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

	c.response.Success(w, http.StatusOK, "successfully to create comment", domain.CommentResponse{
		ID:        result.ID,
		Comment:   result.Comment,
		EventName: result.Event.Name,
		UserID:    result.User.ID,
		EventID:   result.EventID,
		UserName:  result.User.Name,
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

	var commentList []domain.CommentResponse

	for _, comment := range result {
		commentList = append(commentList, domain.CommentResponse{
			ID:        comment.ID,
			Comment:   comment.Comment,
			EventName: comment.Event.Name,
			UserID:    comment.User.ID,
			EventID:   comment.EventID,
			UserName:  comment.User.Name,
			CreateAt:  comment.CreatedAt,
			UpdateAt:  comment.UpdatedAt,
		})
	}

	c.response.Success(w, http.StatusOK, "successfully get all comments", commentList)
}
