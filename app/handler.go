package app

import (
	"net/http"

	"bitbucket.org/zkrhm-fdn/fire-starter/messages"

	"github.com/zkrhm/httphelper"
)

//Hello example template handler. please remove once you have serious handler.
func (app *App) Hello(w http.ResponseWriter, r *http.Request) {

	err := httphelper.WriteAsJSON(w, messages.GenericResponse{
		Code:    200,
		Message: "Hello",
	})

	if err != nil {
		httphelper.WriteAsJSON(w, messages.GenericResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

}
