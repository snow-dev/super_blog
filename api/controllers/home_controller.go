package controllers

import (
	"github.com/snow-dev/super_blog/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this Super API")
}
