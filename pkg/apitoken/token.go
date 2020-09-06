package token

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
    "github.com/rh-eu/golang-example-for-testing-che/pkg/apiutils"
)

type Token struct {
}

func New() *Token {
	return &Token{}
}

type Message struct {
	Token string `json:"token"`
}

type APIuser struct {
	Name string `json:"name"`
}

func (t *Token) AddRoutes(r *httprouter.Router, base string) {
	r.POST(base, apiutils.IsBasicAuthAuthorized(t.APIGet))
}

func (t *Token) APIGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	mySigningKey := []byte(apiutils.ViperEnvVariable("APISigningKey"))

	log.Printf("Request Headers: %v", r.Header)
	log.Printf("Request Body: %s", r.Body)

	var u APIuser
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if u.Name == "" {
		http.Error(w, "Not Authorized", 403)
		return
	}

	log.Println("The name is ... ", u.Name)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "API Token"
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	ss, err := token.SignedString(mySigningKey)
	log.Printf("%v %v", ss, err)

	data := Message{ss}
	apiutils.ServeJSON(w, data)
}