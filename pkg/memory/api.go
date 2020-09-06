package memory

import (
	"net/http"
	"log"
	"runtime"

	"github.com/julienschmidt/httprouter"
	"github.com/rh-eu/golang-example-for-testing-che/pkg/apiutils"

)	

// MemoryAPI ...
type MemoryAPI struct {
	leaks [][]byte
}

// MemoryStatus is returned from a GET to this API endpoint
type MemoryStatus struct {
	MemStats runtime.MemStats `json:"memStats"`
}

func New() *MemoryAPI {
	return &MemoryAPI{}
}

func (m *MemoryAPI) AddRoutes(r *httprouter.Router, base string) {
	r.GET(base+"/api", m.APIGet)
	r.POST(base+"/api/alloc", m.APIAlloc)
	r.POST(base+"/api/clear", m.APIClear)
}

func (m *MemoryAPI) APIGet(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Println("access memory api")

	resp := &MemoryStatus{}

	runtime.ReadMemStats(&resp.MemStats)

	apiutils.ServeJSON(w, resp)

}

func (m *MemoryAPI) APIAlloc(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("memory api alloc")
}

func (m *MemoryAPI) APIClear(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Println("memory api clear")
}