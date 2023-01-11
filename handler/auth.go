package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/renaldyhidayatt/blogGoEnt/dto/request"
	"github.com/renaldyhidayatt/blogGoEnt/dto/response"
	"github.com/renaldyhidayatt/blogGoEnt/security"
	"github.com/renaldyhidayatt/blogGoEnt/services"
)

type authHandler struct {
	auth services.AuthService
}

func NewAuthHandler(auth services.AuthService) *authHandler {
	return &authHandler{auth: auth}
}

func (h *authHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userRequest request.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)

	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = userRequest.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.auth.RegisterUser(userRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	} else {
		response.ResponseMessage(w, "Berhasil membuat data", res, http.StatusCreated)
	}
}

func (h *authHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var authRequest request.AuthRequest

	err := json.NewDecoder(r.Body).Decode(&authRequest)

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	err = authRequest.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.auth.LoginUser(authRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if res.ID > 0 {
		hashPwd := res.Password
		pwd := authRequest.Password

		hash := security.VerifyPassword(hashPwd, pwd)

		if hash == nil {
			token, err := security.GenerateToken(res.Email)

			if err != nil {
				response.ResponseError(w, http.StatusInternalServerError, err)
				return
			}

			response.ResponseToken(w, "Login Berhasil", token, res, http.StatusOK)
		} else {
			response.ResponseError(w, http.StatusBadRequest, errors.New("password tidak sesuai"))
			return
		}
	}
}
