package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	var token string
	var err error

	data := chooseData(w, r, "action")
	if data == nil {
		return
	}

	action := data.Get("action")
	if action == "SignOut" {
		cke := http.Cookie{
			Name:    "wms_manager_token",
			Value:   "",
			Expires: time.Unix(0, 0),
		}
		http.SetCookie(w, &cke)
		w.WriteHeader(http.StatusOK)
		return
	}

	if _, err := r.Cookie("wms_manager_token"); err == nil {
		w.WriteHeader(200)
		return
	}

	login := data.Get("login")
	password := data.Get("password")
	password2 := data.Get("password2")

	switch action {
	case "SignIn":
		token, err = userSignIn(login, password)
	case "SignUp":
		token, err = userSignUp(login, password, password2)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	if err != nil {
		switch token {
		case "Forbidden":
			w.WriteHeader(http.StatusForbidden)
		case "InnerError":
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	} else {
		cke := http.Cookie{
			Name:  "wms_manager_token",
			Value: token,
		}
		http.SetCookie(w, &cke)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	data := chooseData(w, r, "table")
	if data == nil {
		return
	}

	anymap, err := tableGet(data.Get("table"), data.Get("start"), data.Get("limit"))

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := json.MarshalIndent(anymap, "", "   ")
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(resp)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	data := chooseData(w, r, "table")
	if data == nil {
		return
	}

	cols, err := tableGetColumns(data.Get("table"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	values := []string{}
	for _, col := range cols {
		values = append(values, data.Get(col))
	}

	switch data.Get("table") {
	case "Instances":
		Tbl := NewInstance()
		fillStruct(Tbl, values)
		err = Tbl.Update()
	case "instancesInfo":
		Tbl := NewInstanceInfo()
		fillStruct(Tbl, values)
		err = Tbl.Update()
	case "instanceParts":
		Tbl := NewInstancePart()
		fillStruct(Tbl, values)
		err = Tbl.Update()
	case "Items":
		Tbl := NewItem()
		fillStruct(Tbl, values)
		err = Tbl.Update()
	case "Permissions":
		Tbl := NewPermission()
		fillStruct(Tbl, values)
		err = Tbl.Update()
	case "Roles":
		Tbl := NewRole()
		fillStruct(Tbl, values)
		err = Tbl.Update()
	case "Managers":
		Tbl := NewManager()
		fillStruct(Tbl, values)
		err = Tbl.Update()
	default:
		err = errors.New("CError: No such table or the table is not available")
	}

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	data := chooseData(w, r, "table")
	if data == nil {
		return
	}

	cols, err := tableGetColumns(data.Get("table"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	values := []string{}
	for _, col := range cols {
		values = append(values, data.Get(col))
	}

	switch data.Get("table") {
	case "Instances":
		Tbl := NewInstance()
		fillStruct(Tbl, values)
		err = Tbl.Create()
	case "instanceParts":
		Tbl := NewInstancePart()
		fillStruct(Tbl, values)
		err = Tbl.Create()
	case "Items":
		Tbl := NewItem()
		fillStruct(Tbl, values)
		err = Tbl.Create()
	case "Permissions":
		Tbl := NewPermission()
		fillStruct(Tbl, values)
		err = Tbl.Create()
	case "Roles":
		Tbl := NewRole()
		fillStruct(Tbl, values)
		err = Tbl.Create()
	case "Managers":
		Tbl := NewManager()
		fillStruct(Tbl, values)
		err = Tbl.Create()
	default:
		err = errors.New("CError: No such table or the table is not available")
	}

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/Auth", Auth)
	mux.HandleFunc("/GetAll", GetAll)
	mux.HandleFunc("/Update", Update)
	mux.HandleFunc("/Create", Create)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

// func authRequired(handleFunc http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token, err := r.Cookie("wms_manager_token")
// 		tokenValue := strings.Trim(token.Value, " ")
// 		if err != nil && len(tokenValue) < 45 {
// 			w.Write([]byte("Authorization required"))
// 			w.WriteHeader(http.StatusForbidden)
// 			return
// 		}

// 		rightsErr := checkRights(token)

// 		switch rightsErr{
// 		case nil:
// 			handleFunc.ServeHTTP(w, r)
// 		case :
// 		}
// 	})
// }
