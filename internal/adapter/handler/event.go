package handler

import (
	"encoding/json"
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

type EventHandler struct {
	service  port.EventService
	validate *validator.Validate
	response util.Response
	claims   domain.Claims
}

func NewEventHandler(service port.EventService) port.EventHandler {
	return &EventHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (e *EventHandler) Create(w http.ResponseWriter, r *http.Request) {
	var entity domain.Event

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to read request body", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &entity)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err.Error())
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)
	entity.User = response.User{
		ID:   claims.ID,
		Name: claims.Name,
	}

	errValidate := util.Validate(e.validate, entity)
	if errValidate != nil {
		e.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := e.service.Create(entity)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed create event", err.Error())
		return
	}

	var participantList []response.ParticipantList
	for _, participant := range result.Participants {
		participantList = append(participantList, response.ParticipantList{
			Name:  participant.User.Name,
			Email: participant.User.Email,
		})
	}

	e.response.Success(w, http.StatusOK, "success create event", response.Event{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Owner: response.User{
			ID:   result.User.ID,
			Name: result.User.Name,
		},
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

	var eventList []response.Event
	for _, event := range result {
		var commentList []response.CommentList
		for _, comment := range event.Comments {
			commentList = append(commentList, response.CommentList{
				Name:    comment.User.Name,
				Comment: comment.Comment,
			})
		}

		var participantList []response.ParticipantList
		for _, participant := range event.Participants {
			participantList = append(participantList, response.ParticipantList{
				Name:  participant.User.Name,
				Email: participant.User.Email,
			})
		}

		eventList = append(eventList, response.Event{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			Owner: response.User{
				ID:   event.User.ID,
				Name: event.User.Name,
			},
			EndDate:      event.EndDate,
			Participants: participantList,
			UpdatedAt:    event.UpdatedAt,
			CreatedAt:    event.CreatedAt,
		})
	}

	e.response.Success(w, http.StatusOK, "successfully retrived event", eventList)
}

func (e *EventHandler) GetAllByUser(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*domain.Claims)

	result, err := e.service.GetAllByUser(claims.ID)
	if err != nil {
		e.response.Error(w, http.StatusInternalServerError, "failed to retrieve event", err.Error())
		return
	}

	var eventList []response.Event
	for _, event := range result {
		var commentList []response.CommentList
		for _, comment := range event.Comments {
			commentList = append(commentList, response.CommentList{
				Name:    comment.User.Name,
				Comment: comment.Comment,
			})
		}

		var participantList []response.ParticipantList
		for _, participant := range event.Participants {
			participantList = append(participantList, response.ParticipantList{
				Name:  participant.User.Name,
				Email: participant.User.Email,
			})
		}

		eventList = append(eventList, response.Event{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			Owner: response.User{
				ID:   event.User.ID,
				Name: event.User.Name,
			},
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
	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "invalid event id", err.Error())
		return
	}

	result, err := e.service.GetByID(uint(id))
	if err != nil {
		e.response.Error(w, http.StatusInternalServerError, "failed to get event by id", err.Error())
		return
	}

	var commentList []response.CommentList
	for _, comment := range result.Comments {
		commentList = append(commentList, response.CommentList{
			Name:    comment.User.Name,
			Comment: comment.Comment,
		})
	}

	var participantList []response.ParticipantList
	for _, participant := range result.Participants {
		participantList = append(participantList, response.ParticipantList{
			Name:  participant.User.Name,
			Email: participant.User.Email,
		})
	}

	e.response.Success(w, http.StatusOK, "successfully retrived event id", response.Event{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Owner: response.User{
			ID:   result.User.ID,
			Name: result.User.Name,
		},
		EndDate:      result.EndDate,
		Participants: participantList,
		UpdatedAt:    result.UpdatedAt,
		CreatedAt:    result.CreatedAt,
		Comments:     commentList,
	})
}

func (e *EventHandler) Update(w http.ResponseWriter, r *http.Request) {
	var entity domain.Event

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

	err = json.Unmarshal(reqBody, &entity)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err.Error())
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)
	entity.User = response.User{
		ID:   claims.ID,
		Name: claims.Name,
	}

	errValidate := util.Validate(e.validate, entity)
	if errValidate != nil {
		e.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := e.service.Update(entity, uint(id))
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed update event", err.Error())
		return
	}

	var commentList []response.CommentList
	for _, comment := range result.Comments {
		commentList = append(commentList, response.CommentList{
			Name:    comment.User.Name,
			Comment: comment.Comment,
		})
	}

	var participantList []response.ParticipantList
	for _, participant := range result.Participants {
		participantList = append(participantList, response.ParticipantList{
			Name:  participant.User.Name,
			Email: participant.User.Email,
		})
	}

	e.response.Success(w, http.StatusOK, "successfully update event", response.Event{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Owner: response.User{
			ID:   result.User.ID,
			Name: result.User.Name,
		},
		EndDate:      result.EndDate,
		Participants: participantList,
		UpdatedAt:    result.UpdatedAt,
		CreatedAt:    result.CreatedAt,
		Comments:     commentList,
	})
}

func (e *EventHandler) JoinEvent(w http.ResponseWriter, r *http.Request) {
	var entity domain.Participant

	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "invalid user id", err.Error())
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)
	entity.UserID = claims.ID
	entity.EventID = uint(id)

	result, err := e.service.JoinEvent(entity)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to join event", err.Error())
		return
	}

	e.response.Success(w, http.StatusOK, "successfully join event", result)
}

func (e *EventHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var entity domain.Event

	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		e.response.Error(w, http.StatusBadGateway, "invalid event id", err.Error())
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)
	entity.User = response.User{
		ID:   claims.ID,
		Name: claims.Name,
	}

	entity.ID = uint(id)

	err = e.service.Delete(entity)
	if err != nil {
		e.response.Error(w, http.StatusBadGateway, "failed deleted event", err.Error())
		return
	}

	e.response.Success(w, http.StatusNoContent, "successfully deleted event", nil)
}
