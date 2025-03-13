package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

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

	err = app.store.Product.AddProduct(r.Context(), &data)

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

func (app *application) getProduct(w http.ResponseWriter, r *http.Request) {
	products, err := app.store.Product.GetAllProduct(r.Context())

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

// func (app *application) getBlogById(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.PathValue("id"))

// 	baseResp := response.BaseResponse{}

// 	if err != nil {
// 		baseResp.Status = http.StatusBadRequest
// 		baseResp.Message = "invalid id"
// 		resp, _ := baseResp.MarshalBaseResponse()
// 		http.Error(w, string(resp), http.StatusBadRequest)
// 		return
// 	}

// 	blog, err := app.store.Blogs.GetById(r.Context(), int64(id))

// 	if err != nil {
// 		log.Println(err.Error())
// 		baseResp.Status = http.StatusInternalServerError
// 		baseResp.Message = "internal server error"
// 		resp, _ := baseResp.MarshalBaseResponse()
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 		return
// 	}

// 	baseResp.Status = http.StatusOK
// 	baseResp.Message = "Success"

// 	blogResp, _ := response.BlogResponse{
// 		BaseResponse: baseResp,
// 		Blog:         blog,
// 	}.MarshalBlogResponse()
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(blogResp)
// }

// func (app *application) deleteBlog(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.PathValue("id"))

// 	baseResp := response.BaseResponse{}

// 	if err != nil {
// 		baseResp.Status = http.StatusBadRequest
// 		baseResp.Message = "invalid id"
// 		resp, _ := baseResp.MarshalBaseResponse()
// 		http.Error(w, string(resp), http.StatusBadRequest)
// 		return
// 	}

// 	err = app.store.Blogs.DeleteById(r.Context(), int64(id))

// 	if err != nil {
// 		log.Println(err.Error())
// 		baseResp.Status = http.StatusInternalServerError
// 		baseResp.Message = "internal server error"
// 		resp, _ := baseResp.MarshalBaseResponse()
// 		http.Error(w, string(resp), http.StatusInternalServerError)
// 		return
// 	}

// 	baseResp.Status = http.StatusOK
// 	baseResp.Message = "Success delete blog"

// 	baseRespJson, _ := baseResp.MarshalBaseResponse()
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(baseRespJson)
// }

// func (app *application) putBlog(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("put blog"))
// }

// func (app *application) patchBlog(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("patch blog"))
// }

// func (app *application) blogComments(w http.ResponseWriter, r *http.Request) {
// 	id := r.PathValue("id")

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintf("blog comments %s", id)))
// }

func (app *application) ProductRouter() *http.ServeMux {
	blogRouter := http.NewServeMux()

	blogRouter.HandleFunc("POST /add", middleware.AdminHandler(app.addProduct))
	blogRouter.HandleFunc("GET /getall", app.getProduct)
	// blogRouter.HandleFunc("GET /{id}/comments", app.blogComments)
	// blogRouter.HandleFunc("GET /{id}", app.getBlogById)
	// blogRouter.HandleFunc("DELETE /{id}", app.deleteBlog)
	// blogRouter.HandleFunc("PUT /{id}", app.putBlog)
	// blogRouter.HandleFunc("PATCH /{id}", app.patchBlog)

	// Catch-all route for undefined paths
	blogRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	return blogRouter
}
