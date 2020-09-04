package app

import (
	"log"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/rh-eu/golang-example-for-testing-che/pkg/htmlutils"
	"github.com/rh-eu/golang-example-for-testing-che/pkg/sitedata"
)

// App ...
type App struct {
	r *httprouter.Router
	tg *htmlutils.TemplateGroup
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func (k *App) getRoothandler() httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		k.tg.Render(w, "index.html")
	})
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

    router.GET("/", k.getRoothandler())
	router.GET("/hello/:name", Hello)

	router.Handler("GET", "/built/*filepath", http.FileServer(sitedata.Assets))
	
	return k
}