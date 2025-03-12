package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

func (app *application) addSomething(w http.ResponseWriter, r *http.Request) {
	// do something
}

func (app *application) postBlog(w http.ResponseWriter, r *http.Request) {
	baseResp := response.BaseResponse{}

	var data response.Blog
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "Invalid request"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = app.store.Blogs.CreateWithDB(r.Context(), &data)

	if err != nil {
		log.Println(err.Error())
		baseResp.Status = http.StatusInternalServerError
		baseResp.Message = "Internal server error"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success"
	resp, _ := response.BlogResponse{
		BaseResponse: baseResp,
		Blog:         data,
	}.MarshalBlogResponse()

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (app *application) getBlog(w http.ResponseWriter, r *http.Request) {
	blogs, err := app.store.Blogs.GetAll(r.Context())

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
	blogResp, _ := response.BlogsResponse{
		BaseResponse: baseResp,
		Blogs:        blogs,
	}.MarshalBlogsResponse()

	w.WriteHeader(http.StatusOK)
	w.Write(blogResp)
}

func (app *application) getBlogById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	baseResp := response.BaseResponse{}

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "invalid id"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}

	blog, err := app.store.Blogs.GetById(r.Context(), int64(id))

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

	blogResp, _ := response.BlogResponse{
		BaseResponse: baseResp,
		Blog:         blog,
	}.MarshalBlogResponse()
	w.WriteHeader(http.StatusOK)
	w.Write(blogResp)
}

func (app *application) deleteBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	baseResp := response.BaseResponse{}

	if err != nil {
		baseResp.Status = http.StatusBadRequest
		baseResp.Message = "invalid id"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusBadRequest)
		return
	}

	err = app.store.Blogs.DeleteById(r.Context(), int64(id))

	if err != nil {
		log.Println(err.Error())
		baseResp.Status = http.StatusInternalServerError
		baseResp.Message = "internal server error"
		resp, _ := baseResp.MarshalBaseResponse()
		http.Error(w, string(resp), http.StatusInternalServerError)
		return
	}

	baseResp.Status = http.StatusOK
	baseResp.Message = "Success delete blog"

	baseRespJson, _ := baseResp.MarshalBaseResponse()
	w.WriteHeader(http.StatusOK)
	w.Write(baseRespJson)
}

func (app *application) putBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("put blog"))
}

func (app *application) patchBlog(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("patch blog"))
}

func (app *application) blogComments(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("blog comments %s", id)))
}

func (app *application) BlogRouter() *http.ServeMux {
	blogRouter := http.NewServeMux()

	blogRouter.HandleFunc("POST /", app.postBlog)
	blogRouter.HandleFunc("GET /", app.getBlog)
	blogRouter.HandleFunc("GET /{id}/comments", app.blogComments)
	blogRouter.HandleFunc("GET /{id}", app.getBlogById)
	blogRouter.HandleFunc("DELETE /{id}", app.deleteBlog)
	blogRouter.HandleFunc("PUT /{id}", app.putBlog)
	blogRouter.HandleFunc("PATCH /{id}", app.patchBlog)

	// Catch-all route for undefined paths
	blogRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	return blogRouter
}
