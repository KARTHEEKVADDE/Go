package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
	//"crypto/md5"

	_ "github.com/go-sql-driver/mysql"
)

//ID Name Years Of Experience Email ID Project Name Designation Skillset Completed Trainings
//Project Aquired Skills Achievements Employee Status View Edit Delete
type User struct {
	ID                   int
	Name                 string
	YearsOfExperience    string
	EmailID              string
	ProjectName          string
	Designation          string
	Skillset             []string
	CompletedTrainings   []string
	ProjectAquiredSkills []string
	Achievements         []string
	EmployeeStatus       string
	Password             string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "pwd"
	dbName := "Kartheek"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM User ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	res := []User{}
	for selDB.Next() {

		var ID, YearsOfExperience int
		var Name, EmailID, ProjectName, Designation, EmployeeStatus, Password string
		var Skillset, CompletedTrainings, ProjectAquiredSkills, Achievements []string

		err = selDB.Scan(&id, &name, &city, &skills, &password)
		if err != nil {
			panic(err.Error())
		}
		user.ID = ID
		user.Name = Name
		user.EmailID = EmailID
		user.ProjectName = ProjectName
		user.Designation = Designation
		user.EmployeeStatus = EmployeeStatus
		user.Skillset = Skillset
		user.CompletedTrainings = CompletedTrainings
		user.ProjectAquiredSkills = ProjectAquiredSkills
		user.Achievements = Achievements
		user.Password = Password
		res = append(res, user)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM User WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	for selDB.Next() {
		var ID, YearsOfExperience int
		var Name, EmailID, ProjectName, Designation, EmployeeStatus, Password string
		var Skillset, CompletedTrainings, ProjectAquiredSkills, Achievements []string
		err = selDB.Scan(&id, &name, &city, &amount, &password)
		if err != nil {
			panic(err.Error())
		}
		user.ID = ID
		user.Name = Name
		user.EmailID = EmailID
		user.ProjectName = ProjectName
		user.Designation = Designation
		user.EmployeeStatus = EmployeeStatus
		user.Skillset = Skillset
		user.CompletedTrainings = CompletedTrainings
		user.ProjectAquiredSkills = ProjectAquiredSkills
		user.Achievements = Achievements
		user.Password = Password
		res = append(res, user)
	}
	tmpl.ExecuteTemplate(w, "Show", user)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM User WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	for selDB.Next() {
		var ID, YearsOfExperience int
		var Name, EmailID, ProjectName, Designation, EmployeeStatus, Password string
		var Skillset, CompletedTrainings, ProjectAquiredSkills, Achievements []string

		err = selDB.Scan(&id, &name, &city, &amount, &password)
		if err != nil {
			panic(err.Error())
		}
		user.ID = ID
		user.Name = Name
		user.EmailID = EmailID
		user.ProjectName = ProjectName
		user.Designation = Designation
		user.EmployeeStatus = EmployeeStatus
		user.Skillset = Skillset
		user.CompletedTrainings = CompletedTrainings
		user.ProjectAquiredSkills = ProjectAquiredSkills
		user.Achievements = Achievements
		user.Password = Password

	}
	tmpl.ExecuteTemplate(w, "Edit", user)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		amount := r.FormValue("amount")
		password := r.FormValue("password")
		insForm, err := db.Prepare("INSERT INTO User(name, city, amount, password) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, amount, password)
		log.Println("INSERT: Name: " + name + " | City: " + city + " | Amount: " + amount + " | Password: " + password)
	}
	http.Redirect(w, r, "/", 301)
	defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		amount := r.FormValue("amount")
		password := r.FormValue("password")
		insForm, err := db.Prepare("UPDATE User SET name=?, city=?, amount=?, password=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, amount, password, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city + " | Amount: " + amount + " | Password: " + password)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	user := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM User WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(user)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
