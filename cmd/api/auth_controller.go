package main

import (
	"encoding/json"
	"net/http"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/response"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/utils"
)

// @Summary      Login
// @Description  Perform login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request	body	  request.LoginRequest	true "Login request"
// @Success      200  		{object}  response.LoginResponse
// @Failure      400  		{object}  response.BaseResponse
// @Failure      404  		{object}  response.BaseResponse
// @Router       /auth/login	[post]
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var data request.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	err = app.validator.Struct(data)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, token, err := app.store.IAuth.Login(r.Context(), data)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Invalid email/password!")
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.LoginResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Success",
		},
		Data: response.LoginData{
			ID:    int(user.ID),
			Token: token,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

// @Summary      Register
// @Description  Perform register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request		body	  request.RegisterReq	true "Register request"
// @Success      200  			{object}  response.BaseResponse
// @Failure      400  			{object}  response.BaseResponse
// @Failure      404  			{object}  response.BaseResponse
// @Router       /auth/register	[post]
func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var data request.RegisterReq
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	err = app.validator.Struct(data)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = app.store.IAuth.Register(r.Context(), data)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.BaseResponse{
		Status: http.StatusOK,
		Message: "Success register account",
	})
}

func (app *application) AuthRouter() *http.ServeMux {
	authRouter := http.NewServeMux()

	authRouter.HandleFunc("POST /login", app.login)
	authRouter.HandleFunc("POST /register", app.register)

	return authRouter
}
