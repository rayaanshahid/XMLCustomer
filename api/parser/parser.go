package parser

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
)

// our struct which contains the complete
// array of all Clients in the file
type Clients struct {
    XMLName xml.Name `xml:"Client_Export"`
    Clients   []Client   `xml:"Client"`
}

// our struct which contains the complete
// array of all Assignments in the file
type Assignments struct {
    XMLName xml.Name `xml:"Assignment_Export"`
    Assignments   []Assignment   `xml:"Assignment"`
}

// our struct which contains the complete
// array of all Timesheets in the file
type Timesheets struct {
    XMLName xml.Name `xml:"Timesheet_Export"`
    Timesheets   []Timesheet   `xml:"Timesheet"`
}

type Client struct {
    XMLName xml.Name `xml:"Client"`
    Client_Name    string   `xml:"Client_Name"`
    Invoice_Name  string   `xml:"Invoice_Name"`
    Invoice_Point_ID  string  `xml:"Invoice_Point_ID"`
    Invoice_Add1  string  `xml:"Invoice_Add1"`
    Invoice_Add2  string  `xml:"Invoice_Add2"`
    Invoice_Town  string  `xml:"Invoice_Town"`
    Invoice_Country string  `xml:"Invoice_Country"`
    Invoice_Postcode  string   `xml:"Invoice_Postcode"`
    Invoice_Tel_No  string   `xml:"Invoice_Tel_No"`
    Invoice_Email string   `xml:"Invoice_Email"`
    Owning_Region string   `xml:"Owning_Region"`
}

type Assignment struct {
    XMLName xml.Name `xml:"Assignment"`
    Assignment_ID    string   `xml:"Assignment_ID"`
    Job_Title  string   `xml:"Job_Title"`
    Client_ID  string  `xml:"Client_ID"`
    Product  string  `xml:"Product"`
    Pay_Type  string  `xml:"Pay_Type"`
    Charge_Amount  string  `xml:"Charge_Amount"`
    PO_Number string  `xml:"PO_Number"`
    Cand_FName  string   `xml:"Cand_FName"`
    Cand_LName  string   `xml:"Cand_LName"`
    Cand_ID string   `xml:"Cand_ID"`
    Invoice_ID string   `xml:"Invoice_ID"`
    Owning_Cons string   `xml:"Owning_Cons"`
    Assig_Con string   `xml:"Assig_Con"`
    Invoice_Fee string   `xml:"Invoice_Fee"`
    Job_ID string   `xml:"Job_ID"`
    Rate_1_Amount string   `xml:"Rate_1_Amount"`
}

type Timesheet struct {
    XMLName xml.Name `xml:"Timesheet"`
    Monthly_Timesheet_ID    string   `xml:"Monthly_Timesheet_ID"`
    Time_Line_ID  string   `xml:"Time_Line_ID"`
    Time_Line_Date  string  `xml:"Time_Line_Date"`
    Normal  string  `xml:"Normal"`
    Semester  string  `xml:"Semester"`
    Sjukfranvaro_karensdag  string  `xml:"Sjukfrånvaro_karensdag"`
    Sjukfranvaro_dag_2_14 string  `xml:"Sjukfrånvaro_dag_2-14"`
    Vard_av_barn  string   `xml:"Vård_av_barn"`
    Overtid_1  string   `xml:"Övertid_1"`
    Overtid_2 string   `xml:"Övertid_2"`
    Foraldraledighet string   `xml:"Föräldraledighet"`
    Ovrig_Franvaro string   `xml:"Övrig_Frånvaro"`
    Rate1_Amount string   `xml:"Rate1_Amount"`
    Cand_ID string   `xml:"Cand_ID"`
    Job_ID string   `xml:"Job_ID"`
    Invoice_ID string   `xml:"Invoice_ID"`
    Assignment_ID string   `xml:"Assignment_ID"`
}

func GetClients() Clients{
  // Open our xmlFile
  xmlFileClient, err := os.Open("files/Client.xml")
  // if we os.Open returns an error then handle it
  if err != nil {
      fmt.Println(err)
  }

  fmt.Println("Successfully Opened Client.xml")
  // defer the closing of our xmlFile so that we can parse it later on
  defer xmlFileClient.Close()

  // read our opened xmlFile as a byte array.
  byteValue, _ := ioutil.ReadAll(xmlFileClient)

  // we initialize our clients array
  var clients Clients
  // we unmarshal our byteArray which contains our
  // xmlFiles content into 'clients' which we defined above
  xml.Unmarshal(byteValue, &clients)
  return clients
}

func GetAssignments() Assignments{
  // Open our xmlFile
  xmlFileAssignment, err := os.Open("files/Assignment.xml")
  // if we os.Open returns an error then handle it
  if err != nil {
      fmt.Println(err)
  }

  fmt.Println("Successfully Opened Assignment.xml")
  // defer the closing of our xmlFile so that we can parse it later on
  defer xmlFileAssignment.Close()

  // read our opened xmlFile as a byte array.
  byteValue1, _ := ioutil.ReadAll(xmlFileAssignment)

  // we initialize our clients array
  var assignments Assignments
  // we unmarshal our byteArray which contains our
  // xmlFiles content into 'assignments' which we defined above
  xml.Unmarshal(byteValue1, &assignments)
  return assignments
}

func GetTimesheets() Timesheets{
  // Open our xmlFile
  xmlFileTimesheet, err := os.Open("files/Timesheet.xml")
  // if we os.Open returns an error then handle it
  if err != nil {
      fmt.Println(err)
  }

  fmt.Println("Successfully Opened Timesheet.xml")
  // defer the closing of our xmlFile so that we can parse it later on
  defer xmlFileTimesheet.Close()

  // read our opened xmlFile as a byte array.
  byteValue2, _ := ioutil.ReadAll(xmlFileTimesheet)

  // we initialize our clients array
  var timesheets Timesheets
  // we unmarshal our byteArray which contains our
  // xmlFiles content into 'timesheets' which we defined above
  xml.Unmarshal(byteValue2, &timesheets)
  return timesheets
}

/*func main() {

    for i := 0; i < len(clients.Clients); i++ {
        fmt.Println()
        fmt.Println("Entry Number ",i,": ")
        fmt.Println()
        fmt.Println("Client_Name: " + clients.Clients[i].Client_Name)
        fmt.Println("Invoice_Name: " + clients.Clients[i].Invoice_Name)
        fmt.Println("Invoice_Point_ID: " + clients.Clients[i].Invoice_Point_ID)
        fmt.Println("Invoice_Add1: " + clients.Clients[i].Invoice_Add1)
        fmt.Println("Invoice_Add2: " + clients.Clients[i].Invoice_Add2)
        fmt.Println("Invoice_Town: " + clients.Clients[i].Invoice_Town)
        fmt.Println("Invoice_Country: " + clients.Clients[i].Invoice_Country)
        fmt.Println("Invoice_Postcode: " + clients.Clients[i].Invoice_Postcode)
        fmt.Println("Invoice_Tel_No: " + clients.Clients[i].Invoice_Tel_No)
        fmt.Println("Invoice_Email: " + clients.Clients[i].Invoice_Email)
        fmt.Println("Owning_Region: " + clients.Clients[i].Owning_Region)
        fmt.Println()
        fmt.Println("******************************************************************")
        fmt.Println()
    }

    for i := 0; i < len(assignments.Assignments); i++ {
        fmt.Println("Entry Number ",i,": ")
        fmt.Println()
        fmt.Println("Assignment_ID: " + assignments.Assignments[i].Assignment_ID)
        fmt.Println("Job_Title: " + assignments.Assignments[i].Job_Title)
        fmt.Println("Client_ID: " + assignments.Assignments[i].Client_ID)
        fmt.Println("Product: " + assignments.Assignments[i].Product)
        fmt.Println("Pay_Type: " + assignments.Assignments[i].Pay_Type)
        fmt.Println("Charge_Amount: " + assignments.Assignments[i].Charge_Amount)
        fmt.Println("PO_Number: " + assignments.Assignments[i].PO_Number)
        fmt.Println("Cand_FName: " + assignments.Assignments[i].Cand_FName)
        fmt.Println("Cand_LName: " + assignments.Assignments[i].Cand_LName)
        fmt.Println("Cand_ID: " + assignments.Assignments[i].Cand_ID)
        fmt.Println("Invoice_ID: " + assignments.Assignments[i].Invoice_ID)
        fmt.Println("Owning_Cons: " + assignments.Assignments[i].Owning_Cons)
        fmt.Println("Assig_Con: " + assignments.Assignments[i].Assig_Con)
        fmt.Println("Invoice_Fee: " + assignments.Assignments[i].Invoice_Fee)
        fmt.Println("Job_ID: " + assignments.Assignments[i].Job_ID)
        fmt.Println("Rate_1_Amount: " + assignments.Assignments[i].Rate_1_Amount)
        fmt.Println()
        fmt.Println("******************************************************************")
        fmt.Println()
    }

    for i := 0; i < len(timesheets.Timesheets); i++ {
        fmt.Println("Entry Number ",i,": ")
        fmt.Println()
        fmt.Println("Monthly_Timesheet_ID: " + timesheets.Timesheets[i].Monthly_Timesheet_ID)
        fmt.Println("Time_Line_ID: " + timesheets.Timesheets[i].Time_Line_ID)
        fmt.Println("Time_Line_Date: " + timesheets.Timesheets[i].Time_Line_Date)
        fmt.Println("Normal: " + timesheets.Timesheets[i].Normal)
        fmt.Println("Semester: " + timesheets.Timesheets[i].Semester)
        fmt.Println("Sjukfranvaro_karensdag: " + timesheets.Timesheets[i].Sjukfranvaro_karensdag)
        fmt.Println("Sjukfranvaro_dag_2_14: " + timesheets.Timesheets[i].Sjukfranvaro_dag_2_14)
        fmt.Println("Vard_av_barn: " + timesheets.Timesheets[i].Vard_av_barn)
        fmt.Println("overtid_1: " + timesheets.Timesheets[i].overtid_1)
        fmt.Println("overtid_2: " + timesheets.Timesheets[i].overtid_2)
        fmt.Println("Foraldraledighet: " + timesheets.Timesheets[i].Foraldraledighet)
        fmt.Println("ovrig_Franvaro: " + timesheets.Timesheets[i].ovrig_Franvaro)
        fmt.Println("Rate1_Amount: " + timesheets.Timesheets[i].Rate1_Amount)
        fmt.Println("Cand_ID: " + timesheets.Timesheets[i].Cand_ID)
        fmt.Println("Job_ID: " + timesheets.Timesheets[i].Job_ID)
        fmt.Println("Invoice_ID: " + timesheets.Timesheets[i].Invoice_ID)
        fmt.Println("Assignment_ID: " + timesheets.Timesheets[i].Assignment_ID)
        fmt.Println()
        fmt.Println("******************************************************************")
        fmt.Println()
    }
}*/
