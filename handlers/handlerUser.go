package handlers

import (
	"day2/connections"
	"day2/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	Checkcookies(w, r)
	Refreshcookies(w, r)
	fmt.Fprintf(w, "Welcome! by Jimmi")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//read request
	payloads, _ := ioutil.ReadAll(r.Body)

	var user structs.User
	//masukin data request ke user
	json.Unmarshal(payloads, &user)
	var risk structs.Risk_profile
	//bikin risk profile dengan condition
	risk.UserID = user.ID
	if user.Age < 20 {
		risk.Stock_percent = 34.5
		risk.Bond_percent = 45.5
		risk.MM_percent = 100 - (risk.Stock_percent + risk.Bond_percent)
	} else if 30 > user.Age && user.Age >= 20 {
		risk.Stock_percent = 54.5
		risk.Bond_percent = 25.5
		risk.MM_percent = 100 - (risk.Stock_percent + risk.Bond_percent)
	} else if user.Age >= 30 {
		risk.Stock_percent = 72.5
		risk.Bond_percent = 21.5
		risk.MM_percent = 100 - (risk.Stock_percent + risk.Bond_percent)
	}
	user.Risk_profile = risk
	//create user
	connections.DB.Create(&user)
	//return
	res := structs.Result{Code: 200, Data: user, Message: "Success create new User & User's Risk profile"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	Checkcookies(w, r)
	Refreshcookies(w, r)
	//set limit & offset
	vars := mux.Vars(r)
	page := vars["page"]
	take := vars["take"]

	users := []structs.User{}
	//ambil data users dari db
	connections.DB.
		Limit(page).
		Offset(take).
		Find(&users)
	//return
	res := structs.Result{Code: 200, Data: users, Message: "Success get User list"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	Checkcookies(w, r)
	Refreshcookies(w, r)
	//set id
	vars := mux.Vars(r)
	id := vars["id"]

	user := structs.User{}
	//ambil data user dari db
	connections.DB.
		First(&user, id)
	risk_profile := structs.Risk_profile{}
	//ambil data risk dari db
	connections.DB.
		First(&risk_profile, "user_id =?", user.ID)
	user.Risk_profile = risk_profile
	//return
	res := structs.Result{Code: 200, Data: user, Message: "Success get User Detail"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}
