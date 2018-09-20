package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//App central struct of the app
type App struct {
	router *mux.Router
}

//NewApp app struct constructor, return pointer to app
func NewApp() *App {
	return &App{}
}

//Initialize . initializing app: filling the router with func etc.
func (app *App) Initialize() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", app.Hello)

	app.router = router
}

//Run running the server (router defined in initilize / initRouting)
func (app *App) Run(addr string) {
	server := http.Server{
		Handler: app.router,
		Addr:    addr,
	}
	fmt.Println("> Running on :: ", addr)
	server.ListenAndServe()
}
