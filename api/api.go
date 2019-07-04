package main

import (
	"database"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func showPostHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//setCors(w)
	var tableclient database.Tableclient
	database.DB.Where("Invoice_Point_ID = ?", ps.ByName("invoicePointId")).First(&tableclient)
	res, err := json.Marshal(tableclient)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(res)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//setCors(w)
	var tableclients []database.Tableclient
	database.DB.Find(&tableclients)
	res, err := json.Marshal(tableclients)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(res)
}

func main() {
	defer database.DB.Close()

	// add router and routes
	router := httprouter.New()

	router.GET("/clients/:invoicePointId", showPostHandler)
	router.GET("/clients", indexPostHandler)

	// add database
	_, err := database.Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}

	log.Println("connected to DB")

	clients := getClients()
	testPost := Post{Author: "Dorper", Message: "GoDoRP is Dope"}
	DB.Create(&testPost)

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}
	log.Println("Listening on server !!")
	http.ListenAndServe(":8080", router)
}
