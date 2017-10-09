package main

import (
	"fmt"
	"github.com/bradmang18/lotterypicker/contexts"
	"github.com/bradmang18/lotterypicker/models"
	"github.com/go-sql-driver/mysql"
	"github.com/gocraft/web"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
)

type APIContext struct {
	User models.User
}

var DB *gorm.DB

func main() {
	// Database connection
	config := mysql.Config{
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		User:   getEnv("DB_USER", "lottery_picker"),
		Passwd: getEnv("DB_PASSWORD", "lottery_picker"),
		DBName: getEnv("DB_DATABASE", "lottery_picker"),
	}
	var err error
	if DB, err = gorm.Open(getEnv("DB_DRIVER", "mysql"), config.FormatDSN()); err != nil {
		panic("Couldn't connect to database")
	}

	// Root router
	root := web.New(contexts.RootContext{})
	root.Post("/register", (*contexts.RootContext).Register)
	root.Post("/login", (*contexts.RootContext).Login)
	root.Post("/forgot", (*contexts.RootContext).Forgot)
	root.Post("/reset", (*contexts.RootContext).Reset)

	// API Router
	api := root.Subrouter(contexts.APIContext{}, "/api")
	api.Middleware((*contexts.APIContext).UserRequired)
	api.Get("/factories", (*contexts.APIContext).List)
	api.Post("/factories", (*contexts.APIContext).Create)
	api.Get("/factories/:id:\\d+", (*contexts.APIContext).Find)
	api.Put("/factories/:id\\d+", (*contexts.APIContext).Update)
	api.Delete("/factories/:id\\d+", (*contexts.APIContext).Delete)

	// Listen
	http.ListenAndServe(fmt.Sprintf("localhost:%s", getEnv("APP_PORT", "9000")), root)
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
