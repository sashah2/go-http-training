package routes

import (
	"github.com/gin-gonic/gin"
	"go-http-training/routes/user"
	"net/http"
)

type endpoint struct {
	Type    string
	Path    string
	Handler func(*gin.Context)
}

var Endpoints = []endpoint{
	{
		Type:    http.MethodPost,
		Path:    "/users/create",
		Handler: user.Create,
	},
	{
		Type:    http.MethodGet,
		Path:    "/users/get",
		Handler: user.Get,
	},
	{
		Type:    http.MethodGet,
		Path:    "/users/list",
		Handler: user.List,
	},
	{
		Type:    http.MethodPut,
		Path:    "/users/update",
		Handler: user.Update,
	},
}
