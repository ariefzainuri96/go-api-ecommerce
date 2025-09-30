// @title           Your Ecommerce API
// @version         1.0
// @description     This is the documentation for the main e-commerce service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /v1

package main

import (
	"log"
	"os"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/docs"
	"github.com/ariefzainuri96/go-api-blogging/internal/db"
	"github.com/ariefzainuri96/go-api-blogging/internal/store"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	db, err := db.New(os.Getenv("DB_ADDR"), 30, 30, "10m")

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	defer db.Close()

	envHost := os.Getenv("SWAGGER_HOST")

	if envHost == "" {
		envHost = "localhost:8080"
	}

	docs.SwaggerInfo.Host = envHost

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			addr:         os.Getenv("DB_ADDR"),
			maxOpenCons:  30,
			maxIdleConns: 30,
			maxIdleTime:  "10m",
		},
	}

	store := store.NewStorage(db)

	validate := validator.New()

	app := &application{
		config:    cfg,
		store:     store,
		validator: validate,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
