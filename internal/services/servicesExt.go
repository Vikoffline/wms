package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	errors "wms/errors"
	db_services "wms/internal/db_services"

	"golang.org/x/crypto/bcrypt"
)

func chooseData(w http.ResponseWriter, r *http.Request, reqAttr string) url.Values {
	var data url.Values

	if r.Method == http.MethodPost {
		data = r.Form
	} else if r.URL.Query().Has(reqAttr) {
		data = r.URL.Query()
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return nil
	}

	return data
}

func userSignIn(login, password string) (string, error) {
	var err error

	Mn := db_services.NewManager()
	err = Mn.Find(login)
	if err != nil {
		return "Forbidden", fmt.Errorf("CError: %s: %w", errors.ErrUserNotFound, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(Mn.Password), []byte(password))

	if err != nil {
		return "Forbidden", fmt.Errorf("CError: %s: %w", errors.ErrWrongPassword, err)
	}

	Sn := db_services.NewSession()
	Sn.ManagerId = Mn.Id
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errors.ErrSesKeyGen, err)
	}
	Sn.Token = base64.URLEncoding.EncodeToString(key)

	err = Sn.Create()
	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errors.ErrSesCreation, err)
	}

	return Sn.Token, nil
}

func userSignUp(login, password, password2 string) (string, error) {
	var err error
	var IsValidLogin = strings.Trim(login, " ") != ""
	var IsValidPass = (password == password2) && strings.Trim(password, " ") != ""
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /
	// ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ / ЗАГЛУШКА / ДОПИСАТЬ /

	if !IsValidLogin {
		return "Forbidden", fmt.Errorf("CError: %s: %w", errors.ErrInvalidLogin, err)
	}
	if !IsValidPass {
		return "Forbidden", fmt.Errorf("CError: %s: %w", errors.ErrInvalidPasssword, err)
	}

	Mn := db_services.NewManager()
	Mn.Login = login
	Mn.RoleId = "Rl_2"
	hashedPswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errors.ErrPasEncryption, err)
	}

	Mn.Password = string(hashedPswd)

	err = Mn.Create()

	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errors.ErrUserCreation, err)
	}

	Sn := db_services.NewSession()
	Sn.ManagerId = Mn.Id
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errors.ErrSesKeyGen, err)
	}
	Sn.Token = base64.URLEncoding.EncodeToString(key)

	err = Sn.Create()
	if err != nil {
		return "InnerError", fmt.Errorf("CError: %s: %w", errors.ErrSesCreation, err)
	}

	return Sn.Token, nil
}

func CheckRights(r *http.Request, w http.ResponseWriter) (data url.Values) {
	cke, err := r.Cookie("wms_manager_token")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}

	if data = chooseData(w, r, "table"); data == nil {
		return
	}

	action := strings.Replace(data.Get("action"), " ", "", -1)
	tableName := strings.Replace(data.Get("table"), " ", "", -1)
	if action == "" || tableName == "" {
		return nil
	}

	res := db_services.SnCheckRights.QueryRow(cke.Value, action, tableName)

	err = res.Scan(&action)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}

	return data
}
