package memory

import (
	"net/http"
	"log"
	"runtime"
	"runtime/debug"
	"strconv"

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
	r.GET(base+"/api", apiutils.IsAuthorized(m.APIGet))
	r.POST(base+"/api/alloc", apiutils.IsAuthorized(m.APIAlloc))
	r.POST(base+"/api/clear", apiutils.IsAuthorized(m.APIClear))
}

func (m *MemoryAPI) APIGet(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Println("access memory api")

	resp := &MemoryStatus{}

	runtime.ReadMemStats(&resp.MemStats)

	apiutils.ServeJSON(w, resp)

}

func (m *MemoryAPI) APIAlloc(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("memory api alloc")

	sSize := r.URL.Query().Get("size")
	if len(sSize) == 0 {
		http.Error(w, "size not specified", http.StatusBadRequest)
		return
	}

	i, err := strconv.ParseInt(sSize, 10, 64)
	if err != nil {
		http.Error(w, "bad size param", http.StatusBadRequest)
	}

	leak := make([]byte, i, i)
	for i := 0; i < len(leak); i++ {
		leak[i] = 'x'
	}

	m.leaks = append(m.leaks, leak)	
}

func (m *MemoryAPI) APIClear(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Println("memory api clear")

	m.leaks = nil
	runtime.GC()
	debug.FreeOSMemory()	
}