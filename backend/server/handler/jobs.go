package handler

import (
	"fmt"
	"net/http"
)

type JobHandler struct{}

func NewJobHandler(usecase string) *JobHandler {
	return &JobHandler{}
}

func (handler *JobHandler) Test(http.ResponseWriter, *http.Request) {
	fmt.Println("a")
}
