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
// @Tags         login
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
			Status: http.StatusOK,
			Message: "Success",
		},
		Data: response.LoginData{
			ID: int(user.ID),
			Token: token,
			Name: user.Name,
			Email: user.Email,
		},
	})	
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var baseResp response.BaseResponse

	var data request.RegisterReq
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "Invalid request"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = app.validator.Struct(data)

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = err.Error()
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}

	err = app.store.IAuth.Register(r.Context(), data)

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = err.Error()
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success register!"

	resp, _ := baseResp.MarshalBaseResponse()
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *application) AuthRouter() *http.ServeMux {
	authRouter := http.NewServeMux()

	authRouter.HandleFunc("POST /login", app.login)
	authRouter.HandleFunc("POST /register", app.register)

	return authRouter
}
