package htmlutils

import (
	"html/template"
	"log"
	"net/http"

	"github.com/rh-eu/golang-example-for-testing-che/pkg/sitedata"
	"github.com/shurcooL/httpfs/html/vfstemplate"
)

// TemplateGroup ...
type TemplateGroup struct {
}

// Render ...
func (g *TemplateGroup) Render(w http.ResponseWriter, name string) {

	//ts := template.New(name)

	ts, err := vfstemplate.ParseFiles(sitedata.Assets, template.New(name), name)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}