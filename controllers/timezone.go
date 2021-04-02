package controllers

import (
	"net/http"
	"regexp"
)

type timezoneController struct {
	userIDPattern *regexp.Regexp
}

// Method
func (uc timezoneController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the Zone Controller!"))
}

// Constructor methods
func newUserController() *timezoneController {
	return &timezoneController{
		userIDPattern: regexp.MustCompile(`^/timezone/(\d+)/?`),
	}
}
