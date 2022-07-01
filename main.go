package main

import (
	// "bytes"
	// "context"
	"go-fiber-playground/storage"
	"log"
	"os"
	// "reflect"
	// "time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"go-fiber-playground/API/App"
	"go-fiber-playground/API/Context"
	"go-fiber-playground/API/Constants"

	"go-fiber-playground/GUIDE/Routing"

)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host: 		os.Getenv("DB_HOST"),
		Port: 		os.Getenv("DB_PORT"),
		Password: 	os.Getenv("DB_PASS"),
		User: 		os.Getenv("DB_USER"),
		SSLMode: 	os.Getenv("DB_SSLMODE"),
		DBName: 	os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	context := Context.Repository{
		DB: db,
	}

	constants := Constants.Repository{
		DB: db,
	}

	appDoc := App.Repository{
		DB: db,
	}

	routing := Routing.Repository{
		DB: db,
	}

	// app := fiber.New()
	engine := html.New("./template", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routing.SetupRoutesRouting(app)
	appDoc.SetupRoutesApp(app)
	constants.SetupRoutesConstants(app)
	context.SetupRoutesContext(app)
	log.Fatal(app.Listen(":3000"))
}

