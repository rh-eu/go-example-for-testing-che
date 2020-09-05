package app

import (
	"log"
	"fmt"
	"os"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/rh-eu/golang-example-for-testing-che/pkg/htmlutils"
	"github.com/rh-eu/golang-example-for-testing-che/pkg/sitedata"
	"github.com/rh-eu/golang-example-for-testing-che/pkg/version"
)

type serverData struct {
	URLBase      string       `json:"urlBase"`
	Hostname     string       `json:"hostname"`
	Addrs        []string     `json:"addrs"`
	Version      string       `json:"version"`
	RequestDump  string       `json:"requestDump"`
	RequestProto string       `json:"requestProto"`
	RequestAddr  string       `json:"requestAddr"`
}

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

func (k *App) getServerData(r *http.Request, urlBase string) *serverData {
	c := &serverData{}
	c.URLBase = urlBase
	c.Hostname, _ = os.Hostname()

	addrs, _ := net.InterfaceAddrs()
	c.Addrs = []string{}
	for _, addr := range addrs {
		// check the address type and if it is not a loopback
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				c.Addrs = append(c.Addrs, ipnet.IP.String())
			}
		}
	}

	c.Version = version.VERSION
	reqDump, _ := httputil.DumpRequest(r, false)
	c.RequestDump = strings.TrimSpace(string(reqDump))
	c.RequestProto = r.Proto
	c.RequestAddr = r.RemoteAddr

	return c
}

func (k *App) getRoothandler() httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		urlBase := "/"
		k.tg.Render(w, "index.html", k.getServerData(r, urlBase))
	})
}

// Run ...
func (k *App) Run() {
	log.Printf("goservd is up and running")
	log.Fatal(http.ListenAndServe(":8080", k.r))
}

func faviconHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	file, err := sitedata.Assets.Open("static/images/favicon.ico")
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(string(content)))
	defer file.Close()
}

// NewApp ...
func NewApp() *App {
	k := &App{
		r: httprouter.New(),
	}

	router := k.r

	router.GET("/favicon.ico", faviconHandler)

    router.GET("/", k.getRoothandler())
	router.GET("/hello/:name", Hello)

	router.Handler("GET", "/static/*filepath", http.FileServer(sitedata.Assets))
	router.Handler("GET", "/built/*filepath", http.FileServer(sitedata.Assets))
	
	return k
}