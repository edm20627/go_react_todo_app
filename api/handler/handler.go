package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) http.Handler {
	e := echo.New()

	h := &handler{
		DB: db,
	}

	e.GET("/todo", h.getTodoLists)
	e.POST("/todo", h.createTodo)
	e.DELETE("/todo/:id", h.deleteTodo)

	return e
}
