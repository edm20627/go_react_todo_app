package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/edm20627/go_react_todo_app/handler"
)

const defaultPort = "8080"

var Db *gorm.DB

func init() {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		"root",
		"root",
		"mysql:3306",
		"development",
	)
	var err error
	Db, err = gorm.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func port() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":" + defaultPort
}

func main() {
	h := handler.New(Db)

	server := &http.Server{
		Addr:    port(),
		Handler: h,
	}

	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

		<-sigch
		log.Println("Shutting down...")

		if err := server.Shutdown(context.Background()); err != nil {
			log.Println("Unable to shutdown:", err)
		}

		log.Println("Server stopped")
	}()

	log.Println("Listening on http://localhost" + port())
	log.Fatal(server.ListenAndServe())
}
