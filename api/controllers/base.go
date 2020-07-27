package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver

	"github.com/snow-dev/super_blog/models"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(driver, user, password, port, host, dbname string) {
	var err error

	if driver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
		server.DB, err = gorm.Open(driver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", driver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("We are connected to the %s database", driver)
		}
	}

	// database migration
	server.DB.Debug().AutoMigrate(
		&models.User{},
		&models.Post{},
	)

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Printf("Listening to port %s\n", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}
