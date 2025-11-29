package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/response"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/utils"
)

// @Summary      Add Cart
// @Description  Add cart data
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        request		body	  request.AddToCartRequest	true "Add cart request"
// @security 	 ApiKeyAuth
// @Success      200  			{object}  response.CartsResponse
// @Failure      400  			{object}  response.BaseResponse
// @Failure      404  			{object}  response.BaseResponse
// @Router       /cart/add		[post]
func (app *application) addToCart(w http.ResponseWriter, r *http.Request) {
	var data request.AddToCartRequest

	user, ok := middleware.GetUserFromContext(r)

	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "Unauthorized, please re login!")
		return
	}	

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request!")
		return
	}
	defer r.Body.Close()

	err = app.validator.Struct(data)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	cart, err := app.store.ICart.AddToCart(r.Context(), data, user["user_id"].(int))

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Internal server error!")
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.CartResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Success",
		},
		Cart: cart,
	})
}

// @Summary      Get Cart
// @Description  Get cart data
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        request		query	  request.PaginationRequest	true "Pagination request"
// @security 	 ApiKeyAuth
// @Success      200  			{object}  response.CartsResponse
// @Failure      400  			{object}  response.BaseResponse
// @Failure      404  			{object}  response.BaseResponse
// @Router       /cart/getall	[get]
func (app *application) getCart(w http.ResponseWriter, r *http.Request) {
	var request request.PaginationRequest

	user, ok := middleware.GetUserFromContext(r)

	if !ok {
		http.Error(w, "Unauthorized, please re login!", http.StatusUnauthorized)
		return
	}

	err := decoder.Decode(&request, r.URL.Query())

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	result, err := app.store.ICart.GetCart(r.Context(), user["user_id"].(int), request)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.CartsResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Success",
		},
		Carts:      result.Data,
		Pagination: result.Pagination,
	})
}

// @Summary      Delete Cart
// @Description  Delete cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        id   				path      int  true  "Cart ID"
// @security 	 ApiKeyAuth
// @Success      200  				{object}  response.BaseResponse
// @Failure      400  				{object}  response.BaseResponse
// @Failure      404  				{object}  response.BaseResponse
// @Router       /cart/remove/{id}	[delete]
func (app *application) deleteCart(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	err = app.store.ICart.DeleteFromCart(r.Context(), id)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success delete cart!",
	})
}

// @Summary      Update Cart
// @Description  Update cart data
// @Tags         cart
// @Accept       json
// @Produce      json
// @Param        id   					path      int  true  "Cart ID"
// @Param        request				body	  request.AddToCartRequest	true "Update cart request"
// @security 	 ApiKeyAuth
// @Success      200  					{object}  response.CartResponse
// @Failure      400  					{object}  response.BaseResponse
// @Failure      404  					{object}  response.BaseResponse
// @Router       /cart/update/{id}		[patch]
func (app *application) updateCart(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var updateData request.UpdateCartRequest
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	err = app.validator.Struct(updateData)

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := updateData.ToMap()

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	cart, err := app.store.ICart.UpdateQuantityCart(r.Context(), id, data)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.CartResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Success updating cart!",
		},
		Cart: cart,
	})
}

func (app *application) CartRouter() *http.ServeMux {
	cartRouter := http.NewServeMux()

	cartRouter.HandleFunc("POST /add", middleware.UserHandler(app.addToCart))
	cartRouter.HandleFunc("DELETE /remove/{id}", middleware.UserHandler(app.deleteCart))
	cartRouter.HandleFunc("PATCH /update/{id}", middleware.UserHandler(app.updateCart))
	cartRouter.HandleFunc("GET /getall", middleware.UserHandler(app.getCart))

	// Catch-all route for undefined paths
	cartRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	return cartRouter
}
