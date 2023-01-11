package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/dto/response"
	"github.com/renaldyhidayatt/blogGoEnt/services"
	"github.com/renaldyhidayatt/blogGoEnt/utils"
)

type postHandler struct {
	post services.PostService
}

func NewPostHandler(post services.PostService) *postHandler {
	return &postHandler{post: post}
}

func (h *postHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.post.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *postHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.post.FindById(int(id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	response.ResponseMessage(w, "Berhasil mendapatkan data", res, http.StatusOK)
}

func (h *postHandler) Create(w http.ResponseWriter, r *http.Request) {
	var postRequest request.PostRequest

	fileupload := utils.NewLocationUpload("posts", "")

	postRequest.Title = r.FormValue("title")
	postRequest.Slug = r.FormValue("slug")
	postRequest.Image = fileupload.FileUpload(w, r)
	postRequest.Body = r.FormValue("body")

	err := postRequest.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.post.Create(postRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}

}

func (h *postHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	var postRequest request.PostRequest
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}
	res, err := h.post.FindById(int(id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	fileupload := utils.NewLocationUpload("posts", res.Img)

	postRequest.Title = r.FormValue("title")
	postRequest.Slug = r.FormValue("slug")
	postRequest.Image = fileupload.FileUpload(w, r)
	postRequest.Body = r.FormValue("body")

	resupdate, err := h.post.UpdateById(int(id), postRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	} else {
		response.ResponseMessage(w, "Berhasil update data", resupdate, http.StatusCreated)
	}
}

func (h *postHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	}

	res, err := h.post.DeleteById(int(id))

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
	} else {
		response.ResponseMessage(w, "Berhasil delete data", res, http.StatusCreated)
	}
}
