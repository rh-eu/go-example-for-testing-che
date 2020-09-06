package apiutils

import (
	"encoding/json"
	"encoding/base64"
	"net/http"
	"time"
	"log"
	"fmt"
	"strings"
	"database/sql"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

var epoch = time.Unix(0, 0).Format(time.RFC1123)
var noCacheHeaders = map[string]string{
	"Expires":       epoch,
	"Cache-Control": "no-cache, private, max-age=0",
	"Pragma":        "no-cache",
}

func NoCache(w http.ResponseWriter) {
	for k, v := range noCacheHeaders {
		w.Header().Set(k, v)
	}
}

func ServeJSON(w http.ResponseWriter, o interface{}) {
	w.Header().Set("Content-Type", "application/json")
	NoCache(w)
	json.NewEncoder(w).Encode(o)
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

// use viper package to read .env file
// return the value of the key
func ViperEnvVariable(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func IsBasicAuthAuthorized(endpoint func(http.ResponseWriter, *http.Request, httprouter.Params)) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		if r.Header["Authorization"] == nil {
			http.Error(w, "Not Authorized", 403)
			return
		}

		if r.Header["Token"] == nil {
			http.Error(w, "Not Authorized", 403)
			return
		}

		if r.ContentLength == 0 {
			http.Error(w, "Please send a request body", 400)
			return
		}

		// Validating via Basic Authentication
		var userPass = ViperEnvVariable("BASICAUTH_CREDENTIALS")
		valid := true
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Basic ") {
			log.Print("Invalid authorization:", auth)
			valid = false
		}
		up, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil {
			log.Print("authorization decode error:", err)
			valid = false
		}
		if string(up) != userPass {
			log.Print("invalid username:password:", string(up))
			valid = false
		}
		if valid {
			endpoint(w, r, p)
		} else {
			http.Error(w, "Not Authorized", 403)
			return
		}
	})
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request, httprouter.Params)) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		if r.Header["Token"] == nil {
			http.Error(w, "Not Authorized", 403)
			return
		}
		mySigningKey := []byte(ViperEnvVariable("APISigningKey"))

		// Validating Token
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return mySigningKey, nil
		})

		if err != nil {
			log.Printf("There was an Error ... %v", err.Error())
			JSONError(w, err.Error(), 400)
			return
		}

		if token.Valid {
			endpoint(w, r, p)
		} else {
			http.Error(w, "Not Authorized", 403)
			return
		}
	})
}

func GetDB() (*sql.DB, error) {
	// viper package read .env
	dbuser := ViperEnvVariable("MYSQL_USER")
	dbpass := ViperEnvVariable("MYSQL_PASS")
	mysqldb := ViperEnvVariable("MYSQL_DB")
	mysqlport := ViperEnvVariable("MYSQL_PORT")
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:"+mysqlport+")/"+mysqldb)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}