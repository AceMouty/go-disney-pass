package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/acemouty/disney-pass/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	// Env setup
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	dbUrl := os.Getenv("PGCON")
	log.Print("db url::", dbUrl, "\n")

	//  DB Setup
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Printf("Error connecting to DB:\n")
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	// Setup api config
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	appRouter := chi.NewRouter()
	appRouter.Use(cors.AllowAll().Handler)
	// Sanity Checks...
	appRouter.Get("/readyness", handlerReadiness)
	appRouter.Get("/err", handlerErr)

	// App Route registration
	// TODO: refactor into something cleaner...
	appRouter.Post("/api/user/register", apiCfg.handleCreateUser)
	appRouter.Post("/api/user/login", apiCfg.handleLoginUser)
	appRouter.Post("/api/parent-posts", apiCfg.handleCreateParentPost)
	appRouter.Get("/api/tester", apiCfg.handleGetParkInformation)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: appRouter,
	}

	log.Printf("Servering on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
