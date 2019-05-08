package main

import (
    "database/sql"
    "log"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type Employee struct {
    Id    int
    Name  string
    Email string
    YearsOfExperience    int
	ProjectName          string
	Designation          string
	Skillset             string
	CompletedTrainings   string
	ProjectAquiredSkills string
	Achievements         string
	EmployeeStatus       string
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
    selDB, err := db.Query("SELECT * FROM employee ORDER BY id")
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    res := []Employee{}
    for selDB.Next() {
        var id, yearsOfExperience int
        var name, email, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus string
        err = selDB.Scan(&id, &name, &email, &yearsOfExperience, &projectName, &designation, &skillset, &completedTrainings, &projectAquiredSkills, &achievements, &employeeStatus)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Email = email
        emp.YearsOfExperience = yearsOfExperience
        emp.ProjectName = projectName
        emp.Designation = designation
        emp.Skillset = skillset
        emp.CompletedTrainings = completedTrainings
        emp.ProjectAquiredSkills = projectAquiredSkills
        emp.Achievements = achievements
        emp.EmployeeStatus = employeeStatus
        res = append(res, emp)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id, yearsOfExperience int
        var name, email, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus string
        err = selDB.Scan(&id, &name, &email, &yearsOfExperience, &projectName, &designation, &skillset, &completedTrainings, &projectAquiredSkills, &achievements, &employeeStatus)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Email = email
        emp.YearsOfExperience = yearsOfExperience
        emp.ProjectName = projectName
        emp.Designation = designation
        emp.Skillset = skillset
        emp.CompletedTrainings = completedTrainings
        emp.ProjectAquiredSkills = projectAquiredSkills
        emp.Achievements = achievements
        emp.EmployeeStatus = employeeStatus
    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id, yearsOfExperience int
        var name, email, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus string
        err = selDB.Scan(&id, &name, &email, &yearsOfExperience, &projectName, &designation, &skillset, &completedTrainings, &projectAquiredSkills, &achievements, &employeeStatus)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.Email = email
        emp.YearsOfExperience = yearsOfExperience
        emp.ProjectName = projectName
        emp.Designation = designation
        emp.Skillset = skillset
        emp.CompletedTrainings = completedTrainings
        emp.ProjectAquiredSkills = projectAquiredSkills
        emp.Achievements = achievements
        emp.EmployeeStatus = employeeStatus
    }
    tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        email := r.FormValue("email")
        yearsOfExperience := r.FormValue("yearsOfExperience")
        projectName := r.FormValue("projectName")
        designation := r.FormValue("designation")
        skillset := r.FormValue("skillset")
        completedTrainings := r.FormValue("completedTrainings")
        projectAquiredSkills := r.FormValue("projectAquiredSkills")
        achievements := r.FormValue("achievements")
        employeeStatus := r.FormValue("employeeStatus")
        insForm, err := db.Prepare("INSERT INTO employee(name, email, yearsOfExperience, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus) VALUES(?,?,?,?,?,?,?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, email, yearsOfExperience, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus)
        log.Println("INSERT: Name: " + name)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        email := r.FormValue("email")
        yearsOfExperience := r.FormValue("yearsOfExperience")
        projectName := r.FormValue("projectName")
        designation := r.FormValue("designation")
        skillset := r.FormValue("skillset")
        completedTrainings := r.FormValue("completedTrainings")
        projectAquiredSkills := r.FormValue("projectAquiredSkills")
        achievements := r.FormValue("achievements")
        employeeStatus := r.FormValue("employeeStatus")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE employee SET name=?, email=?, yearsOfExperience=?, projectName=?, designation=?, skillset=?, completedTrainings=?, projectAquiredSkills=?, achievements=?, employeeStatus=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, email, yearsOfExperience, projectName, designation, skillset, completedTrainings, projectAquiredSkills, achievements, employeeStatus, id)
        log.Println("UPDATE: Name: " + name)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE")
    http.Redirect(w, r, "/", 301)
    defer db.Close()
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