package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	_ "github.com/lib/pq"
	"encoding/json"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

/*
	Initialize - Initializes the database connection with PostgreSQL and initializes the endpoint routes
		user - Username of PostgreSQL
		password - Password of PostgreSQL
		dbname - Database name
		schema - Schema within the database
*/
func (a *App) Initialize(user, password, dbname, schema string) {
	dbProperties := fmt.Sprintf("user=%s password=%s dbname=%s search_path=%s sslmode=disable", user, password, dbname, schema)
	var err error
	a.DB, err = sql.Open("postgres", dbProperties)
	if err != nil {
		log.Fatal(err)
	}
	
	a.Router = mux.NewRouter()
	a.initializeRoutes()
 }

 /*
	 StartServer - Starts the http server to listen to http requests
 		port - Port number to access
 */
func (a *App) StartServer(port string) {
	// Enabled cors to be able to do post requests with axios in the frontend
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
	})  

	err := http.ListenAndServe(":8080", c.Handler(a.Router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
	initializeRoutes - initializes the endpoints to handle http requests
*/
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/d/{urlCode}", a.getOriginalUrl)
	a.Router.HandleFunc("/submit", a.createUrl).Methods("POST")
	a.Router.HandleFunc("/log/{urlCode}", a.getUrlLogs).Methods("GET")
}

/*
	getUrlLogs - Requests the logs of the url within the /log/{id} endpoint
*/
func (a *App) getUrlLogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	urlCode := vars["urlCode"]

	urlLogs, err := GetUrlLogs(a.DB, urlCode)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJsonResponse(w, http.StatusCreated, urlLogs)
}

/*
	createUrl - Creates the shortened url 
*/
func (a *App) createUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
    var u url
    err := decoder.Decode(&u)
    if err != nil {
        panic(err)
	}
	
	urlCode := RandomString(10)
	u.UrlCode = urlCode
	// Change to real domain
	u.ShortUrl = "http://localhost:8080/d/" + urlCode

	if err := u.CreateUrl(a.DB); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJsonResponse(w, http.StatusCreated, u.ShortUrl)
}

/*
	getOriginalUrl - Gets the original url and redirects user if valid
*/
func (a *App) getOriginalUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uCode := vars["urlCode"]
	// Get the ip address of the client request
	ip,_,_ := net.SplitHostPort(r.RemoteAddr)
	l := urlLog{ UrlCode: uCode, IpAddress: ip}
	if err := l.LogInstance(a.DB); err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	u := url{UrlCode: uCode}
	if err := u.GetOriginalUrl(a.DB); err != nil {
		// TODO: Reset to page displaying that it doesnt exist
		u.OriginalUrl = "https://google.ca"
	}
	http.Redirect(w, r, u.OriginalUrl, http.StatusMovedPermanently)
}

/*
	sendErrorResponse - Sends a json formatted resposne when error occurs
*/
func sendErrorResponse(w http.ResponseWriter, errorId int, message string) {
	sendJsonResponse(w, errorId, map[string]string{"Error": message})
}

/*
	sendJsonResponse - Prepares the paylod into a json response
*/
func sendJsonResponse(w http.ResponseWriter, errorId int, payload interface{}) {
	jsonResponse, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Json conversion failed")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorId)
	w.Write(jsonResponse)
}