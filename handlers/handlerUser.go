package handlers

import (
	"day2/connections"
	"day2/structs"
	"day2/vendor/github.com/dgrijalva/jwt-go"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s! by jimmi", claims.Name)))
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
