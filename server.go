package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.URL.Query().Has("action") {
		var data url.Values
		var token string
		var err error

		switch r.URL.Path {
		case "/FormAuth":
			data = r.Form
		case "/QueryAuth":
			data = r.URL.Query()
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

		if token, err := r.Cookie("wms_manager_token"); err == nil {
			fmt.Println(token, err)

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
			fmt.Println(err)
			switch token {
			case "Forbidden":
				w.WriteHeader(http.StatusForbidden)
			case "InnerError":
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		} else {
			fmt.Println(token)
			cke := http.Cookie{
				Name:  "wms_manager_token",
				Value: token,
			}
			http.SetCookie(w, &cke)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Query().Get("tableName")
	anymap, err := TableGetAll(tableName)

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

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/FormAuth", Auth)
	mux.HandleFunc("/QueryAuth", Auth)
	mux.HandleFunc("/GetAll", GetAll)

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
