package main

import (
	"fmt"
	"github.com/wcl48/valval"
	"github.com/woremacx/go-crud-sample/models"
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
	"strconv"
)

var tpl *template.Template

type FormData struct {
	User models.User
	Mess string
}

func UserRoot(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserRoot")
	http.Redirect(w, r, "/user/", 301)
}

func UserIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	Users := []models.User{}
	db.Find(&Users)
	tpl = template.Must(template.ParseFiles("views/user/index.html"))
	tpl.Execute(w, Users)
}

func UserNew(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("views/user/new.html"))
	tpl.Execute(w, FormData{models.User{}, ""})
}

func UserCreate(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{Name: r.FormValue("Name")}
	if err := models.UserValidate(User); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("views/user/new.html"))
		tpl.Execute(w, FormData{User, Mess})
	} else {
		db.Create(&User)
		http.Redirect(w, r, "/user/", 301)
	}
}

func UserEdit(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Find(&User)
	tpl = template.Must(template.ParseFiles("views/user/edit.html"))
	tpl.Execute(w, FormData{User, ""})
}

func UserUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Find(&User)
	User.Name = r.FormValue("Name")
	if err := models.UserValidate(User); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("views/user/edit.html"))
		tpl.Execute(w, FormData{User, Mess})
	} else {
		db.Save(&User)
		http.Redirect(w, r, "/user/", 301)
	}
}

func UserDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Delete(&User)
	http.Redirect(w, r, "/user/", 301)
}
