package main

import (
	"database"
	"XMLCustomer/api/parser"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func getClientByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func getAssignmentByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//setCors(w)
	var tableassignment database.Tableassignment
	database.DB.Where("Assignment_ID = ?", ps.ByName("assignmentId")).First(&tableassignment)
	res, err := json.Marshal(tableassignment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(res)
}

func getTimesheetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//setCors(w)
	var tabletimesheet database.Tabletimesheet
	database.DB.Where("Time_Line_ID = ?", ps.ByName("timelineId")).First(&tabletimesheet)
	res, err := json.Marshal(tabletimesheet)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(res)
}

func getAllClients(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//setCors(w)
	//log.Println("Entered getAllClients() in API !!")
	var tableclients []database.Tableclient
	DB := database.DB.Find(&tableclients)
	if DB.Error != nil {
		log.Println("Error while finding in database in getAllClients() in API !!")
		http.Error(w, "error while querying",500)
		return
	}
	res, err := json.Marshal(tableclients)
	log.Println("took result in getAllClients() in API !!")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(res)
}

func getAllAssignments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//setCors(w)
	//log.Println("Entered getAllAssignments() in API !!")
	var tableassignments []database.Tableassignment
	DB := database.DB.Find(&tableassignments)
	if DB.Error != nil {
		log.Println("Error while finding in database in getAllAssignments() in API !!")
		http.Error(w, "error while querying",500)
		return
	}
	res, err := json.Marshal(tableassignments)
	log.Println("took result in getAllAssignments() in API !!")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(res)
}

func getAllTimesheets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//setCors(w)
	log.Println("Entered getAllTimesheets() in API !!")
	var tabletimesheets []database.Tabletimesheet
	DB := database.DB.Find(&tabletimesheets)
	if DB.Error != nil {
		log.Println("Error while finding in database in getAllTimesheets() in API !!")
		http.Error(w, "error while querying",500)
		return
	}
	res, err := json.Marshal(tabletimesheets)
	log.Println("took result in getAllTimesheets() in API !!")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(res)
}

func getMergedDataByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//setCors(w)
	log.Println("Entered getMergedDataByID() in API !!")
	var results database.MergedData
	DB := database.DB.Table("tabletimesheets").Select("tableclients.Invoice_Point_ID, tableassignments.Assignment_ID, tabletimesheets.Time_Line_ID,tabletimesheets.Monthly_Timesheet_ID, tableclients.Client_Name, tableassignments.Job_Title").Joins("JOIN tableclients on tableclients.Invoice_Point_ID = tabletimesheets.Invoice_ID").Joins("JOIN tableassignments on tableassignments.Assignment_ID = tabletimesheets.Assignment_ID").Where("tabletimesheets.Invoice_ID = ? and tabletimesheets.Assignment_ID = ?",ps.ByName("invoiceID"),ps.ByName("assignmentID")).Find(&results)

	//var tabletimesheets []database.Tabletimesheet
	//DB := database.DB.Find(&tabletimesheets)
	if DB.Error != nil {
		log.Println("Error while finding in database in getMergedDataByID() in API !!", DB.Error)
		http.Error(w, "error while querying",500)
		return
	}
	res, err := json.Marshal(results)
	log.Println("took result in getMergedDataByID() in API !!")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(res)
}

func ReadDataFromXMLToDB(){
	//Read data from xml files
	clients := parser.GetClients()
	assignments := parser.GetAssignments()
	timesheets := parser.GetTimesheets()

	//Insert data into Table Models
	for i:=0;i<len(clients.Clients);i++{
			tableclient := database.Tableclient {
				Client_Name: clients.Clients[i].Client_Name,
				Invoice_Name:clients.Clients[i].Invoice_Name,
				Invoice_Point_ID:clients.Clients[i].Invoice_Point_ID,
				Invoice_Add1:clients.Clients[i].Invoice_Add1,
				Invoice_Add2:clients.Clients[i].Invoice_Add2,
				Invoice_Town:clients.Clients[i].Invoice_Town,
				Invoice_Country:clients.Clients[i].Invoice_Country,
				Invoice_Postcode:clients.Clients[i].Invoice_Postcode,
				Invoice_Tel_No:clients.Clients[i].Invoice_Tel_No,
				Invoice_Email:clients.Clients[i].Invoice_Email,
				Owning_Region:clients.Clients[i].Owning_Region,
			}
			DB := database.DB.Create(&tableclient)
			if DB.Error != nil {
				log.Print("Row did not enter in client table!!!",DB.Error)
			}else {
				log.Println("Row entered in client table!!!")
			}
	}

	for i:=0;i<len(assignments.Assignments);i++{
			tableassignment := database.Tableassignment	{
					Assignment_ID: assignments.Assignments[i].Assignment_ID,
					Job_Title: assignments.Assignments[i].Job_Title,
					Client_ID: assignments.Assignments[i].Client_ID,
					Product: assignments.Assignments[i].Product,
					Pay_Type: assignments.Assignments[i].Pay_Type,
					Charge_Amount: assignments.Assignments[i].Charge_Amount,
					PO_Number: assignments.Assignments[i].PO_Number,
					Cand_FName: assignments.Assignments[i].Cand_FName,
					Cand_LName: assignments.Assignments[i].Cand_LName,
					Cand_ID: assignments.Assignments[i].Cand_ID,
					Invoice_ID: assignments.Assignments[i].Invoice_ID,
					Owning_Cons: assignments.Assignments[i].Owning_Cons,
					Assig_Con: assignments.Assignments[i].Assig_Con,
					Invoice_Fee: assignments.Assignments[i].Invoice_Fee,
					Job_ID: assignments.Assignments[i].Job_ID,
					Rate_1_Amount: assignments.Assignments[i].Rate_1_Amount,
				}
			DB := database.DB.Create(&tableassignment)
			if DB.Error != nil {
				log.Print("Row did not enter in assignment table!!!",DB.Error)
			}else {
				log.Println("Row entered in assignment table!!!")
			}
	}

	for i:=0;i<len(timesheets.Timesheets);i++{
			tabletimesheet := database.Tabletimesheet		{
				    Monthly_Timesheet_ID: timesheets.Timesheets[i].Monthly_Timesheet_ID,
				    Time_Line_ID: timesheets.Timesheets[i].Time_Line_ID,
				    Time_Line_Date: timesheets.Timesheets[i].Time_Line_Date,
				    Normal: timesheets.Timesheets[i].Normal,
				    Semester: timesheets.Timesheets[i].Semester,
				    Sjukfranvaro_karensdag: timesheets.Timesheets[i].Sjukfranvaro_karensdag,
				    Sjukfranvaro_dag_2_14: timesheets.Timesheets[i].Sjukfranvaro_dag_2_14,
				    Vard_av_barn: timesheets.Timesheets[i].Vard_av_barn,
				    Overtid_1: timesheets.Timesheets[i].Overtid_1,
				    Overtid_2: timesheets.Timesheets[i].Overtid_2,
				    Foraldraledighet: timesheets.Timesheets[i].Foraldraledighet,
				    Ovrig_Franvaro: timesheets.Timesheets[i].Ovrig_Franvaro,
				    Rate1_Amount: timesheets.Timesheets[i].Rate1_Amount,
				    Cand_ID: timesheets.Timesheets[i].Cand_ID,
				    Job_ID: timesheets.Timesheets[i].Job_ID,
				    Invoice_ID: timesheets.Timesheets[i].Invoice_ID,
				    Assignment_ID: timesheets.Timesheets[i].Assignment_ID,
				}
			DB := database.DB.Create(&tabletimesheet)
			if DB.Error != nil {
				log.Print("Row did not enter in timesheet table!!!",DB.Error)
			}else {
				log.Println("Row entered in timesheet table!!!")
			}
	}
}

func main() {
	defer database.DB.Close()

	// add router and routes
	router := httprouter.New()

	router.GET("/clients/:invoicePointId", getClientByID)
	router.GET("/clients", getAllClients)
	router.GET("/assignments/:assignmentId", getAssignmentByID)
	router.GET("/assignments", getAllAssignments)
	router.GET("/timesheets/:timelineId", getTimesheetByID)
	router.GET("/timesheets", getAllTimesheets)
	router.GET("/merged/:invoiceID/:assignmentID",getMergedDataByID)

	// add database
	_, err := database.Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}

	log.Println("connected to DB")

	ReadData := false // change this to true to read data from XML files and posting to DB tables
	if ReadData == true {
		ReadDataFromXMLToDB()
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
