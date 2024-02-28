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

type EventHandler struct {
	service  port.EventService
	validate *validator.Validate
	response util.Response
}

func NewEventHandler(service port.EventService) port.EventHandler {
	return &EventHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (e *EventHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.EventRequest

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to read request body", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err.Error())
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)
	req.User = domain.UserResponse{
		ID:   claims.ID,
		Name: claims.Name,
	}

	errValidate := util.Validate(e.validate, req)
	if errValidate != nil {
		e.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := e.service.Create(req)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed create event", err.Error())
		return
	}

	var participantList []domain.ParticipantList
	for _, participant := range result.Participants {
		participantList = append(participantList, domain.ParticipantList{
			Name:  participant.User.Name,
			Email: participant.User.Email,
		})
	}

	e.response.Success(w, http.StatusOK, "success create event", domain.EventResponse{
		ID:           result.ID,
		Name:         result.Name,
		Description:  result.Description,
		UserID:       result.User.ID,
		UserName:     result.User.Name,
		EndDate:      result.EndDate,
		Participants: participantList,
		UpdatedAt:    result.UpdatedAt,
		CreatedAt:    result.CreatedAt,
	})
}

func (e *EventHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := e.service.GetAll()
	if err != nil {
		e.response.Error(w, http.StatusInternalServerError, "failed to retrieve event", err.Error())
		return
	}

	var eventList []domain.EventResponse
	for _, event := range result {
		var commentList []domain.CommentList
		for _, comment := range event.Comments {
			commentList = append(commentList, domain.CommentList{
				Name:    comment.User.Name,
				Comment: comment.Comment,
			})
		}

		var participantList []domain.ParticipantList
		for _, participant := range event.Participants {
			participantList = append(participantList, domain.ParticipantList{
				Name:  participant.User.Name,
				Email: participant.User.Email,
			})
		}

		eventList = append(eventList, domain.EventResponse{
			ID:           event.ID,
			Name:         event.Name,
			Description:  event.Description,
			UserID:       event.User.ID,
			UserName:     event.User.Name,
			EndDate:      event.EndDate,
			Comments:     commentList,
			Participants: participantList,
			UpdatedAt:    event.UpdatedAt,
			CreatedAt:    event.CreatedAt,
		})
	}

	e.response.Success(w, http.StatusOK, "successfully retrived event", eventList)
}

func (e *EventHandler) GetAllByUser(w http.ResponseWriter, r *http.Request) {
	var req domain.EventRequest

	claims := r.Context().Value("claims").(*domain.Claims)
	req.User.ID = claims.ID

	result, err := e.service.GetAllByUser(req)
	if err != nil {
		e.response.Error(w, http.StatusInternalServerError, "failed to retrieve event", err.Error())
		return
	}

	var eventList []domain.EventResponse
	for _, event := range result {
		var commentList []domain.CommentList
		for _, comment := range event.Comments {
			commentList = append(commentList, domain.CommentList{
				Name:    comment.User.Name,
				Comment: comment.Comment,
			})
		}

		var participantList []domain.ParticipantList
		for _, participant := range event.Participants {
			participantList = append(participantList, domain.ParticipantList{
				Name:  participant.User.Name,
				Email: participant.User.Email,
			})
		}

		eventList = append(eventList, domain.EventResponse{
			ID:           event.ID,
			Name:         event.Name,
			Description:  event.Description,
			UserID:       event.User.ID,
			UserName:     event.User.Name,
			EndDate:      event.EndDate,
			Participants: participantList,
			UpdatedAt:    event.UpdatedAt,
			CreatedAt:    event.CreatedAt,
			Comments:     commentList,
		})
	}

	e.response.Success(w, http.StatusOK, "successfully retrivied event", eventList)
}

func (e *EventHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	var req domain.EventRequest

	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "invalid event id", err.Error())
		return
	}

	req.User.ID = uint(id)

	result, err := e.service.GetByID(req)
	if err != nil {
		e.response.Error(w, http.StatusInternalServerError, "failed to get event by id", err.Error())
		return
	}

	var commentList []domain.CommentList
	for _, comment := range result.Comments {
		commentList = append(commentList, domain.CommentList{
			Name:    comment.User.Name,
			Comment: comment.Comment,
		})
	}

	var participantList []domain.ParticipantList
	for _, participant := range result.Participants {
		participantList = append(participantList, domain.ParticipantList{
			Name:  participant.User.Name,
			Email: participant.User.Email,
		})
	}

	e.response.Success(w, http.StatusOK, "successfully retrived event id", domain.EventResponse{
		ID:           result.ID,
		Name:         result.Name,
		Description:  result.Description,
		UserID:       result.User.ID,
		UserName:     result.User.Name,
		EndDate:      result.EndDate,
		Participants: participantList,
		UpdatedAt:    result.UpdatedAt,
		CreatedAt:    result.CreatedAt,
		Comments:     commentList,
	})
}

func (e *EventHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req domain.EventRequest

	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "invalid user id", err.Error())
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to read request body", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err.Error())
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)
	req.User.ID = uint(id)

	errValidate := util.Validate(e.validate, req)
	if errValidate != nil {
		e.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := e.service.Update(req, *claims)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed update event", err.Error())
		return
	}

	var participantList []domain.ParticipantList
	for _, participant := range result.Participants {
		participantList = append(participantList, domain.ParticipantList{
			Name:  participant.User.Name,
			Email: participant.User.Email,
		})
	}

	e.response.Success(w, http.StatusOK, "successfully update event", domain.EventResponse{
		ID:           result.ID,
		Name:         result.Name,
		Description:  result.Description,
		UserID:       result.User.ID,
		UserName:     result.User.Name,
		EndDate:      result.EndDate,
		Participants: participantList,
		UpdatedAt:    result.UpdatedAt,
		CreatedAt:    result.CreatedAt,
	})
}

func (e *EventHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var req domain.EventRequest

	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		e.response.Error(w, http.StatusBadGateway, "invalid event id", err.Error())
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)

	req.User.ID = uint(id)

	err = e.service.Delete(req, *claims)
	if err != nil {
		e.response.Error(w, http.StatusBadGateway, "failed deleted event", err.Error())
		return
	}

	e.response.Success(w, http.StatusOK, "successfully deleted event", nil)
}

func (e *EventHandler) JoinEvent(w http.ResponseWriter, r *http.Request) {
	var req domain.ParticipantRequest

	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "invalid user id", err.Error())
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)
	req.UserID = claims.ID
	req.EventID = uint(id)

	result, err := e.service.JoinEvent(req)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to join event", err.Error())
		return
	}

	e.response.Success(w, http.StatusOK, "successfully join event", domain.ParticipantResponse{
		ID:        result.ID,
		EventID:   result.EventID,
		UserID:    result.UserID,
		EventName: result.Event.Name,
		UserName:  result.User.Name,
		CreateAt:  result.CreatedAt,
		UpdateAt:  result.UpdatedAt,
	})
}
