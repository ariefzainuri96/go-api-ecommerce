package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/response"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/utils"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// @Summary      Add Product
// @Description  Add new product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        request		body	  request.AddProductRequest	true "Add Product request"
// @security 	 ApiKeyAuth
// @Success      200  			{object}  response.ProductResponse
// @Failure      400  			{object}  response.BaseResponse
// @Failure      404  			{object}  response.BaseResponse
// @Router       /product/add	[post]
func (app *application) addProduct(w http.ResponseWriter, r *http.Request) {
	var data request.AddProductRequest
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

	product, err := app.store.IProduct.AddProduct(r.Context(), &data)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.ProductResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Success add product",
		},
		Product: product,
	})
}

// @Summary      Get Product
// @Description  Get All product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        request			query	  request.PaginationRequest	true "Get Product request"
// @security 	 ApiKeyAuth
// @Success      200  				{object}  response.ProductsResponse
// @Failure      400  				{object}  response.BaseResponse
// @Failure      404  				{object}  response.BaseResponse
// @Router       /product/getall	[get]
func (app *application) getProduct(w http.ResponseWriter, r *http.Request) {
	var data request.PaginationRequest

	err := decoder.Decode(&data, r.URL.Query())

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	product, err := app.store.IProduct.GetProduct(r.Context(), data)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.ProductsResponse{
		BaseResponse: response.BaseResponse{
			Message: "Success",
			Status:  http.StatusOK,
		},
		Products:   product.Data,
		Pagination: product.Pagination,
	})
}

// @Summary      Delete Product
// @Description  Delete product
// @Tags         product
// @Produce      json
// @Param        id   					path      int  true  "Product ID"
// @security 	 ApiKeyAuth
// @Success      200  					{object}  response.BaseResponse
// @Failure      400  					{object}  response.BaseResponse
// @Failure      404  					{object}  response.BaseResponse
// @Router       /product/remove/{id}	[delete]
func (app *application) deleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	err = app.store.IProduct.DeleteProduct(r.Context(), uint(id))

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.BaseResponse{
		Status:  http.StatusOK,
		Message: "Success delete product",
	})
}

// @Summary      Patch Product
// @Description  Patch product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param 		 id						path      int  true  "Product ID"
// @Param        request				body	  request.AddProductRequest	true "Add Product request"
// @security 	 ApiKeyAuth
// @Success      200  					{object}  response.ProductResponse
// @Failure      400  					{object}  response.BaseResponse
// @Failure      404  					{object}  response.BaseResponse
// @Router       /product/update/{id}	[patch]
func (app *application) patchProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	// Decode request body into a map
	var updateData map[string]any
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()

	// Ensure there's data to update
	if len(updateData) == 0 {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	product, err := app.store.IProduct.PatchProduct(r.Context(), uint(productID), updateData)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, response.ProductResponse{
		BaseResponse: response.BaseResponse{
			Status:  http.StatusOK,
			Message: "Success patch product",
		},
		Product: product,
	})
}

func (app *application) ProductRouter() *http.ServeMux {
	productRouter := http.NewServeMux()

	productRouter.HandleFunc("POST /add", middleware.AdminHandler(app.addProduct))
	productRouter.HandleFunc("GET /getall", app.getProduct)
	productRouter.HandleFunc("DELETE /remove/{id}", middleware.AdminHandler(app.deleteProduct))
	productRouter.HandleFunc("PATCH /update/{id}", middleware.AdminHandler(app.patchProduct))

	// Catch-all route for undefined paths
	productRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	return productRouter
}
