package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jacky-htg/inventory/libraries/api"
	"github.com/jacky-htg/inventory/libraries/token"
	"github.com/jacky-htg/inventory/models"
	"github.com/jacky-htg/inventory/payloads/request"
	"github.com/jacky-htg/inventory/payloads/response"
	"golang.org/x/crypto/bcrypt"
)

// Auths struct
type Auths struct {
	Db  *sql.DB
	Log *log.Logger
}

// Login http handler
func (u *Auths) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest request.LoginRequest
	err := api.Decode(r, &loginRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		api.ResponseError(w, err)
		return
	}

	uLogin := models.User{Username: loginRequest.Username}
	err = uLogin.GetByUsername(r.Context(), u.Db)
	if err != nil {
		err = fmt.Errorf("call login: %v", err)
		u.Log.Printf("ERROR : %+v", err)
		api.ResponseError(w, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(uLogin.Password), []byte(loginRequest.Password))
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		api.ResponseError(w, api.ErrBadRequest(fmt.Errorf("compare password: %v", err), ""))
		return
	}

	token, err := token.ClaimToken(uLogin.Username)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		api.ResponseError(w, fmt.Errorf("claim token: %v", err))
		return
	}

	var response response.TokenResponse
	response.Token = token

	api.ResponseOK(w, response, http.StatusOK)
}

// CheckUsername : http handler for checking username availability
func (u *Auths) CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		err := errors.New("Username parameter is required")
		u.Log.Printf("error : %s", err)
		api.ResponseError(w, api.ErrBadRequest(err, ""))
		return
	}

	var user models.User
	exists, err := user.CheckUsernameExists(r.Context(), u.Db, username)
	if err != nil {
		u.Log.Printf("error checking username: %s", err)
		api.ResponseError(w, err)
		return
	}

	response := struct {
		Available bool `json:"available"`
	}{
		Available: !exists,
	}

	api.ResponseOK(w, response, http.StatusOK)
}

// CheckEmail : http handler for checking email availability
func (u *Auths) CheckEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		err := errors.New("Email parameter is required")
		u.Log.Printf("error : %s", err)
		api.ResponseError(w, api.ErrBadRequest(err, ""))
		return
	}

	var user models.User
	exists, err := user.CheckEmailExists(r.Context(), u.Db, email)
	if err != nil {
		u.Log.Printf("error checking email: %s", err)
		api.ResponseError(w, err)
		return
	}

	response := struct {
		Available bool `json:"available"`
	}{
		Available: !exists,
	}

	api.ResponseOK(w, response, http.StatusOK)
}
