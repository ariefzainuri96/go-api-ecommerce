package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/entity"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/request"
	"github.com/ariefzainuri96/go-api-ecommerce/cmd/api/response"
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
// @Success      200  			{object}  response.BaseResponse
// @Failure      400  			{object}  response.BaseResponse
// @Failure      404  			{object}  response.BaseResponse
// @Router       /product/add	[post]
func (app *application) addProduct(w http.ResponseWriter, r *http.Request) {
	baseResp := response.BaseResponse{}

	var data request.AddProductRequest
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

	err = app.store.IProduct.AddProduct(r.Context(), &data)

	if err != nil {
		log.Println(err.Error())
		baseResp.Status = http.StatusInternalServerError
		baseResp.Message = "Internal server error"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success add product"
	resp, _ := baseResp.MarshalBaseResponse()

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
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
	var products []entity.Product

	err := decoder.Decode(&data, r.URL.Query())

	if err != nil {
		log.Printf("Error getProduct: %v", err.Error())
		baseResp := response.BaseResponse{}
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "invalid request"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}

	if data.SearchAll == "" {
		products, err = app.store.IProduct.GetAllProduct(r.Context())
	} else {
		products, err = app.store.IProduct.SearchProduct(r.Context(), data.SearchAll)
	}

	baseResp := response.BaseResponse{}

	if err != nil {
		log.Println(err.Error())
		baseResp.Status = http.StatusInternalServerError
		baseResp.Message = "internal server error"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success"
	productResp, _ := response.ProductsResponse{
		BaseResponse: baseResp,
		Products:     products,
	}.MarshalProductsResponse()

	w.WriteHeader(http.StatusOK)
	w.Write(productResp)
}

func (app *application) deleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	baseResp := response.BaseResponse{}

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "invalid id"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}

	err = app.store.IProduct.DeleteProduct(r.Context(), int64(id))

	if err != nil {
		log.Println(err.Error())
		baseResp.Status = http.StatusInternalServerError
		baseResp.Message = err.Error()
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success delete product"

	baseRespJson, _ := baseResp.MarshalBaseResponse()
	w.WriteHeader(http.StatusOK)
	w.Write(baseRespJson)
}

func (app *application) patchProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.PathValue("id"))

	baseResp := response.BaseResponse{}

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "invalid id"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}

	// Decode request body into a map
	var updateData map[string]any
	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "invalid request"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Ensure there's data to update
	if len(updateData) == 0 {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	err = app.store.IProduct.PatchProduct(r.Context(), int64(productID), updateData)

	if err != nil {
		log.Println(err.Error())
		baseResp.Status = http.StatusInternalServerError
		baseResp.Message = "internal server error"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success patch product"

	resp, _ := baseResp.MarshalBaseResponse()
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
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
