package server

import (
	"strings"

	"github.com/google/uuid"
)

type Routes struct{}

const (
	//API
	api  = "/api"
	jobs = api + "/jobs"

	register = "/register"
	login    = "/login"
)

func NewRoutes() *Routes {
	return &Routes{}
}

func joinRoutes(a string, b string) string {
	a = strings.TrimRight(a, "/") + "/"
	b = strings.Trim(b, "/")
	return a + b
}

// API
func (routes *Routes) APIJobs() string {
	return jobs
}

func (routes *Routes) APIJobId(id uuid.UUID) string {
	return joinRoutes(jobs, id.String())
}

// Public
func (routes *Routes) Register() string {
	return register
}

func (routes *Routes) Login() string {
	return login
}
