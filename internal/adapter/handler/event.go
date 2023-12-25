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
	var req domain.Event

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to read request body", err)
		return
	}

	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err)
		return
	}

	claims := r.Context().Value("claims").(*domain.Claims)
	req.User = response.User{
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

	var participantList []response.ParticipantList
	for _, participant := range result.Participant {
		participantList = append(participantList, response.ParticipantList{
			Name:  participant.User.Name,
			Email: participant.User.Email,
		})
	}

	e.response.Success(w, http.StatusOK, "success create event", response.EventInfo{
		Event: response.Event{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Owner: response.User{
				ID:   result.User.ID,
				Name: result.User.Name,
			},
			Participant: participantList,
		},
		EventDetails: response.EventDetails{
			UpdatedAt: result.UpdatedAt,
			CreatedAt: result.CreatedAt,
			EndDate:   result.EndDate,
		},
	})
}

func (e *EventHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := e.service.GetAll()
	if err != nil {
		e.response.Error(w, http.StatusInternalServerError, "failed to retrieve event", err.Error())
		return
	}

	var eventList []response.EventInfo
	for _, event := range result {
		var commentList []response.CommentList
		for _, comment := range event.Comments {
			commentList = append(commentList, response.CommentList{
				Name:    comment.User.Name,
				Comment: comment.Comment,
			})
		}

		var participantList []response.ParticipantList
		for _, participant := range event.Participant {
			participantList = append(participantList, response.ParticipantList{
				Name:  participant.User.Name,
				Email: participant.User.Email,
			})
		}

		eventList = append(eventList, response.EventInfo{
			Event: response.Event{
				ID:          event.ID,
				Title:       event.Title,
				Description: event.Description,
				Owner: response.User{
					ID:   event.User.ID,
					Name: event.User.Name,
				},
				Comments:    commentList,
				Participant: participantList,
			},
			EventDetails: response.EventDetails{
				UpdatedAt: event.UpdatedAt,
				CreatedAt: event.CreatedAt,
				EndDate:   event.EndDate,
			},
		})
	}

	e.response.Success(w, http.StatusOK, "successfully retrived event", eventList)
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
	for _, participant := range result.Participant {
		participantList = append(participantList, response.ParticipantList{
			Name:  participant.User.Name,
			Email: participant.User.Email,
		})
	}

	e.response.Success(w, http.StatusOK, "successfully retrived event id", response.EventInfo{
		Event: response.Event{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Owner: response.User{
				ID:   result.User.ID,
				Name: result.User.Name,
			},
			Comments:    commentList,
			Participant: participantList,
		},
		EventDetails: response.EventDetails{
			UpdatedAt: result.UpdatedAt,
			CreatedAt: result.CreatedAt,
			EndDate:   result.EndDate,
		},
	})
}

func (e *EventHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req domain.Event

	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "invalid user id", err.Error())
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to read request body", err)
		return
	}

	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed to unmarshal request body", err)
		return
	}

	errValidate := util.Validate(e.validate, req)
	if errValidate != nil {
		e.response.Error(w, http.StatusBadRequest, "validation failed", errValidate)
		return
	}

	result, err := e.service.Update(uint(id), req)
	if err != nil {
		e.response.Error(w, http.StatusBadRequest, "failed update event", err.Error())
	}

	var commentList []response.CommentList
	for _, comment := range result.Comments {
		commentList = append(commentList, response.CommentList{
			Name:    comment.User.Name,
			Comment: comment.Comment,
		})
	}

	var participantList []response.ParticipantList
	for _, participant := range result.Participant {
		participantList = append(participantList, response.ParticipantList{
			Name:  participant.User.Name,
			Email: participant.User.Email,
		})
	}

	e.response.Success(w, http.StatusOK, "successfully retrived event id", response.EventInfo{
		Event: response.Event{
			ID:          result.ID,
			Title:       result.Title,
			Description: result.Description,
			Owner: response.User{
				ID:   result.User.ID,
				Name: result.User.Name,
			},
			Comments:    commentList,
			Participant: participantList,
		},
		EventDetails: response.EventDetails{
			UpdatedAt: result.UpdatedAt,
			CreatedAt: result.CreatedAt,
			EndDate:   result.EndDate,
		},
	})
}

func (e *EventHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	params := vars["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		e.response.Error(w, http.StatusBadGateway, "invalid event id", err.Error())
	}

	err = e.service.Delete(uint(id))
	if err != nil {
		e.response.Error(w, http.StatusBadGateway, "failed deleted event", err.Error())
	}

	e.response.Success(w, http.StatusOK, "successfully deleted event", nil)
}
