package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type book struct {
	boarding  string
	departure string
	name      string
	payment   int
}

var db *gorm.DB

func buy(abhi http.ResponseWriter, mahesh *http.Request) {
	mapp := mux.Vars(mahesh)
	key := mapp["banglore"]
	key1 := mapp["kurnool"]
	key2 := mapp["Manobhi"]
	key3 := mapp["800"]
	p, err := strconv.Atoi(key3)

	if err != nil {

	}
	ram := book{boarding: key, departure: key1, name: key2, payment: p}
	db.Save(&ram)
	json.NewEncoder(abhi).Encode(ram)

}
func Api(abhi http.ResponseWriter, yesh *http.Request) {
	mapp := mux.Vars(yesh)

	key3 := mapp["800"]
	p, err := strconv.Atoi(key3)

	if err != nil {

	}
	var customers []book

	db.Where("800 = ?", p).First(&customers)

	json.NewEncoder(abhi).Encode(customers)
}
func init() {
	var err error
	db, err = gorm.Open("mysql", "root:manobhi8686@tcp(127.0.0.1:3306)/sys?charset=utf8&parseTime=True")
	if err != nil {

	}
	db.AutoMigrate(&book{})
}

func main() {
	handler := mux.NewRouter().StrictSlash(true)
	handler.HandleFunc("/raju/{banglore}/{kurnool}/{Manobhi}/{800}/", buy).Methods("POST")
	handler.HandleFunc("/raju/{800}", Api).Methods("GET")
	log.Fatal(http.ListenAndServe(":1200", handler))
}
