package htmlutils

import (
	"html/template"
	"log"
	"net/http"
	"encoding/json"

	"github.com/rh-eu/golang-example-for-testing-che/pkg/sitedata"
	"github.com/shurcooL/httpfs/html/vfstemplate"
)

// TemplateGroup ...
type TemplateGroup struct {
}

// FuncMap ...
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"jsonstring": JSONString,
	}
}

// JSONString ...
func JSONString(v interface{}) (template.JS, error) {
	a, err := json.Marshal(v)
	if err != nil {
		return template.JS(""), err
	}
	return template.JS(a), nil
}

// Render ...
func (g *TemplateGroup) Render(w http.ResponseWriter, name string, context interface{}) {

	//ts := template.New(name)

	ts, err := vfstemplate.ParseFiles(sitedata.Assets, template.New(name).Funcs(FuncMap()), name)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, context)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}