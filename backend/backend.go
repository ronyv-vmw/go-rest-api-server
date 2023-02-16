package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB     *sql.DB
	Port   string
	Router *mux.Router
}

func (a *App) Initialize() {
	fmt.Println("backend: Initializing...")
	db, err := sql.Open("sqlite3", "../practiceit.db")
	if err != nil {
		log.Fatal("backend: Occured fatal error: ", err.Error())
	} else {
		fmt.Println("backend: Successfully connected to sqlite3 database!")
	}

	// set receiver values
	a.DB = db
	a.Router = mux.NewRouter()
	a.InitializeRouter()
}

func (a *App) InitializeRouter() {
	a.Router.HandleFunc("/products", a.allProducts).Methods("GET")
	a.Router.HandleFunc("/product/{id}", a.fetchProduct).Methods("GET")
}

func (a *App) Run() {
	fmt.Printf("###### Start running service on localhost port %v##### \n", a.Port)
	// listen and server service with router.
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}

func (a *App) allProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is allProducts")
	products, err := getProducts(a.DB)

	if err != nil {
		fmt.Printf("getProducts - error %s \n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) fetchProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is fetchProduct")
	vars := mux.Vars(r)
	id := vars["id"]

	var p product
	p.ID, _ = strconv.Atoi(id)
	err := p.getProduct(a.DB)

	if err != nil {
		fmt.Printf("getProduct - error %s \n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondWithJSON(w, http.StatusOK, p)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
