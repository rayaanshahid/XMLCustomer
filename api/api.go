package main

import (
	"database"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func getClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func getClients(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	router.GET("/clients/:invoicePointId", getClient)
	router.GET("/clients", getClients)

	// add database
	_, err := database.Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}

	log.Println("connected to DB")

	clients := getClients()
	//You can insert multiple records too
	for i:=0;i<len(clients.Clients);i++{
			tableclient := database.Tableclient {Client_Name: clients.Clients[i].Client_Name, Invoice_Name:clients.Clients[i].Invoice_Name,
				Invoice_Point_ID:clients.Clients[i].Invoice_Point_ID, Invoice_Add1:clients.Clients[i].Invoice_Add1,
				Invoice_Add2:clients.Clients[i].Invoice_Add2, Invoice_Town:clients.Clients[i].Invoice_Town, Invoice_Country:clients.Clients[i].Invoice_Country,
				Invoice_Postcode:clients.Clients[i].Invoice_Postcode, Invoice_Tel_No:clients.Clients[i].Invoice_Tel_No,
				Invoice_Email:clients.Clients[i].Invoice_Email, Owning_Region:clients.Clients[i].Owning_Region}
			DB := db.Create(&tableclient)
			if DB.Error != nil {
				log.Print("Row did not enter !!!")
			}else {
				log.Println("One row entered !!!")
			}
	}
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
