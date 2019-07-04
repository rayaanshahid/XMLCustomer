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
	gorm.Model
	Client_Name    string
	Invoice_Name  string
	Invoice_Point_ID  string
	Invoice_Add1  string
	Invoice_Add2  string
	Invoice_Town  string
	Invoice_Country string
	Invoice_Postcode  string
	Invoice_Tel_No  string
	Invoice_Email string
	Owning_Region string
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
		log.Println("Successfully Created a Table !!!")
	}else{
		log.Println("Already exists Table !!!")
	}

	//testPost := Post{Author: "Dorper", Message: "GoDoRP is Dope"}
	//DB.Create(&testPost)

	return DB, err
}
