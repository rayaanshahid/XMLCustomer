package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"log"
)

var DB *gorm.DB
var err error

type Tableclient struct {
	//gorm.Model
	Client_Name    string
	Invoice_Name  string
	Invoice_Point_ID  string	`gorm:"primary_key"`
	Invoice_Add1  string
	Invoice_Add2  string
	Invoice_Town  string
	Invoice_Country string
	Invoice_Postcode  string
	Invoice_Tel_No  string
	Invoice_Email string
	Owning_Region string
}

type Tableassignment struct {
	//gorm.Model
	Assignment_ID   string	`gorm:"primary_key"`
	Job_Title  string
	Client_ID  string
	Product  string
	Pay_Type  string
	Charge_Amount  string
	PO_Number string
	Cand_FName  string
	Cand_LName  string
	Cand_ID string
	Invoice_ID string
	Owning_Cons string
	Assig_Con string
	Invoice_Fee string
	Job_ID string
	Rate_1_Amount string
}

type Tabletimesheet struct {
    //gorm.Model
    Monthly_Timesheet_ID    string
    Time_Line_ID  string	`gorm:"primary_key"`
    Time_Line_Date  string
    Normal  string
    Semester  string
    Sjukfranvaro_karensdag  string
    Sjukfranvaro_dag_2_14 string
    Vard_av_barn  string
    Overtid_1  string
    Overtid_2 string
    Foraldraledighet string
    Ovrig_Franvaro string
    Rate1_Amount string
    Cand_ID string
    Job_ID string
    Invoice_ID string
    Assignment_ID string
}

func Init() (*gorm.DB, error) {
	// set up DB connection and then attempt to connect 5 times over 25 seconds
	connectionParams := "user=docker password=docker sslmode=disable host=db"
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open("postgres", connectionParams) // gorm checks Ping on Open
		if err == nil {
			log.Println("Successfully connected to DB !!!")
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return DB, err
	}

	// create table if it does not exist
	if !DB.HasTable(&Tableclient{}) {
		DB.CreateTable(&Tableclient{})
		log.Println("Successfully Created client Table !!!")
	}else{
		//DB.DropTable(&Tableclient{})
		log.Println("Already exists client Table !!!")
	}

	// create table if it does not exist
	if !DB.HasTable(&Tableassignment{}) {
		DB.CreateTable(&Tableassignment{})
		log.Println("Successfully Created assignment Table !!!")
	}else{
		//DB.DropTable(&Tableassignment{})
		log.Println("Already exists assignment Table !!!")
	}

	// create table if it does not exist
	if !DB.HasTable(&Tabletimesheet{}) {
		DB.CreateTable(&Tabletimesheet{})
		log.Println("Successfully Created timesheet Table !!!")
	}else{
		//DB.DropTable(&Tabletimesheet{})
		log.Println("Already exists timesheet Table !!!")
	}

	return DB, err
}
