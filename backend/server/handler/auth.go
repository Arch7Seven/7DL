package handler

import (
	"fmt"
	"net/http"
)

type AuthHandler struct{}

func NewAuthHandler() *JobHandler {
	return &JobHandler{}
}

func (handler *JobHandler) Register(http.ResponseWriter, *http.Request) {
	fmt.Println("a")
}

func (handler *JobHandler) Login(http.ResponseWriter, *http.Request) {
	fmt.Println("a")
}
