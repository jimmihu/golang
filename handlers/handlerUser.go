package handlers

import (
	"day2/connections"
	"day2/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	//ambil data login
	var user structs.User
	json.Unmarshal(payloads, &user)

	var dbuser structs.User
	//ambil data user dari db
	connections.DB.
		Preload("Risk_profile").
		Where("name =?", user.Name).
		First(&dbuser)
	//cek password
	var res structs.Result
	if CheckPasswordHash(user.Password, dbuser.Password) {
		res = structs.Result{Code: 200, Data: dbuser, Message: "Logged in!"}
	} else {
		res = structs.Result{Code: 200, Data: dbuser, Message: "Wrong name or password!"}
	}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! by Jimmi")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//read request
	payloads, _ := ioutil.ReadAll(r.Body)

	var user structs.User
	//masukin data request ke user
	json.Unmarshal(payloads, &user)
	password, _ := HashPassword(user.Password)
	user.Password = password
	if user.Age < 20 {
		user.Risk_profile.Stock_percent = 34.5
		user.Risk_profile.Bond_percent = 45.5
		user.Risk_profile.MM_percent = 100 - (user.Risk_profile.Stock_percent + user.Risk_profile.Bond_percent)
	} else if 30 > user.Age && user.Age >= 20 {
		user.Risk_profile.Stock_percent = 54.5
		user.Risk_profile.Bond_percent = 25.5
		user.Risk_profile.MM_percent = 100 - (user.Risk_profile.Stock_percent + user.Risk_profile.Bond_percent)
	} else if user.Age >= 30 {
		user.Risk_profile.Stock_percent = 72.5
		user.Risk_profile.Bond_percent = 21.5
		user.Risk_profile.MM_percent = 100 - (user.Risk_profile.Stock_percent + user.Risk_profile.Bond_percent)
	}
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
	//set limit & offset
	intlim, err1 := strconv.Atoi(r.URL.Query().Get("limit"))
	intoffs, err2 := strconv.Atoi(r.URL.Query().Get("offset"))
	if err1 != nil {
		intlim = 10
	}
	if err2 != nil {
		intoffs = 0
	}
	users := []structs.User{}
	//ambil data users dari db
	connections.DB.
		Preload("Risk_profile").
		Limit(intlim).
		Offset(intoffs).
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
