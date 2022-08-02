package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	//access log
	log.Println("Starting API...")

	//endpints
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)  
	//http.ListenAndServe(":8080", router)
	http.ListenAndServe(":"+port, router)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Aplication is running")
}