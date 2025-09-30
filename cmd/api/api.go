package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/ariefzainuri96/go-api-blogging/cmd/api/docs"
	middleware "github.com/ariefzainuri96/go-api-blogging/cmd/api/middleware"
	"github.com/ariefzainuri96/go-api-blogging/internal/store"
	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger"
)

type application struct {
	config    config
	store     store.Storage
	validator *validator.Validate
}

type config struct {
	db   dbConfig
	addr string
}

type dbConfig struct {
	addr         string
	maxOpenCons  int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/v1/product/", middleware.Authentication(http.StripPrefix("/v1/product", app.ProductRouter())))

	mux.Handle("/v1/cart/", middleware.Authentication(http.StripPrefix("/v1/cart", app.CartRouter())))

	mux.Handle("/v1/order/", middleware.Authentication(http.StripPrefix("/v1/order", app.OrderRouter())))

	mux.Handle("/v1/auth/", http.StripPrefix("/v1/auth", app.AuthRouter()))

	mux.Handle("/v1/xendit-callback/", http.StripPrefix("/v1/xendit-callback", app.XenditCallbackRouter()))

	mux.Handle("/v1/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/v1/swagger/doc.json"),
	))

	return mux
}

func (app *application) run(mux *http.ServeMux) error {

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      stack(mux),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	log.Printf("Server has started on %s", app.config.addr)

	return srv.ListenAndServe()
}
