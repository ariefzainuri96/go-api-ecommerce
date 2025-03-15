package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

func (app *application) addToCart(w http.ResponseWriter, r *http.Request) {
	var baseResp response.BaseResponse

	var data request.AddToCartRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "Invalid request"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = app.store.Cart.AddToCart(r.Context(), data)

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = err.Error()
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success add to cart!"

	resp, _ := baseResp.MarshalBaseResponse()
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *application) getCart(w http.ResponseWriter, r *http.Request) {
	var baseResp response.BaseResponse

	user, ok := middleware.GetUserFromContext(r)

	if !ok {
		http.Error(w, "Unauthorized, please re login!", http.StatusUnauthorized)
		return
	}

	userId := user["user_id"].(int64)

	log.Println(user)

	carts, err := app.store.Cart.GetCart(r.Context(), userId)

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = err.Error()
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success"

	cartResp, _ := response.CartsResponse{
		BaseResponse: baseResp,
		Carts:        carts,
	}.MarshalResponse()

	w.WriteHeader(http.StatusOK)
	w.Write(cartResp)
}

func (app *application) deleteCart(w http.ResponseWriter, r *http.Request) {
	var baseResp response.BaseResponse

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "invalid id"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}

	err = app.store.Cart.DeleteFromCart(r.Context(), int64(id))

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = err.Error()
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success delete cart"

	resp, _ := baseResp.MarshalBaseResponse()
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *application) updateCart(w http.ResponseWriter, r *http.Request) {
	var baseResp response.BaseResponse

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "invalid id"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}

	var updateData request.UpdateCartRequest
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "invalid request"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = app.store.Cart.UpdateQuantityCart(r.Context(), int64(id), updateData.Quantity)

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = err.Error()
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success update cart"

	resp, _ := baseResp.MarshalBaseResponse()
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *application) CartRouter() *http.ServeMux {
	cartRouter := http.NewServeMux()

	cartRouter.HandleFunc("POST /add", app.addToCart)
	cartRouter.HandleFunc("DELETE /remove/{id}", app.deleteCart)
	cartRouter.HandleFunc("PATCH /update/{id}", app.updateCart)
	cartRouter.HandleFunc("GET /getall", app.getCart)

	// Catch-all route for undefined paths
	cartRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	return cartRouter
}
