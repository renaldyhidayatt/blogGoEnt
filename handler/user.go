package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/dto/response"
	"github.com/renaldyhidayatt/blogGoEnt/services"
)

type userHandler struct {
	user services.UserService
}

func NewUserHandler(user services.UserService) *userHandler {
	return &userHandler{user: user}
}

func (h *userHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.user.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *userHandler) FindById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.user.FindById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = userRequest.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.user.Create(userRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *userHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}
	var userRequest request.UserRequest

	err = json.NewDecoder(r.Body).Decode(&userRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = userRequest.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	res, err := h.user.UpdateById(int(Id), userRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil update data", res, http.StatusCreated)
	}
}

func (h *userHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.user.DeleteById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)

		return
	} else {
		response.ResponseMessage(w, "Berhasil delete data", res, http.StatusCreated)
	}
}
