package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"net/http"
)

var db gorm.DB

func main() {
	router(goji.DefaultMux)
	goji.Serve()
}

func router(m *web.Mux) http.Handler {
	m.Get("/", UserRoot)
	user := web.New()
	goji.Handle("/user/*", user)
	user.Use(middleware.SubRouter)
	user.Get("/", UserIndex)
	user.Get("/new", UserNew)
	user.Post("/new", UserCreate)
	user.Get("/edit/:id", UserEdit)
	user.Post("/update/:id", UserUpdate)
	user.Get("/delete/:id", UserDelete)
	return m
}

func init() {
	db, _ = gorm.Open("mysql", "root@/testdb?charset=utf8&parseTime=True")
}
