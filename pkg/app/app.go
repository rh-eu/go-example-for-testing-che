package app

import (
	"log"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// App ...
type App struct {
	r *httprouter.Router
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// Run ...
func (k *App) Run() {
	log.Printf("goservd is up and running")
	log.Fatal(http.ListenAndServe(":8080", k.r))
}

// NewApp ...
func NewApp() *App {
	k := &App{
		r: httprouter.New(),
	}

	router := k.r

    router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	
	return k
}